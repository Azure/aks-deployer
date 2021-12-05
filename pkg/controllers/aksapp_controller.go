// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package controllers

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/2016-10-01/keyvault"
	"golang.org/x/time/rate"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	k8scontroller "sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	deployerv1 "github.com/Azure/aks-deployer/pkg/api/v1"
	"github.com/Azure/aks-deployer/pkg/auth/tokenprovider"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/auth"
	"github.com/Azure/aks-deployer/pkg/configmaps"
	"github.com/Azure/aks-deployer/pkg/log"
	"github.com/Azure/aks-deployer/pkg/secret"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

const (
	azureClientIDStr     = "AZURE_CLIENT_ID"
	azureClientSecretStr = "AZURE_CLIENT_SECRET" // #nosec only filed name with secret text
	azureTenantIDStr     = "AZURE_TENANT_ID"

	identityResourceIDStr = "IDENTITY_RESOURCE_ID"
	keyVaultResourceStr   = "KEY_VAULT_RESOURCE"

	defaultReconcileBaseDelay     = 10 * time.Second
	reconcileMaxDelay             = 300 * time.Second
	defaultRolloutRecheckInterval = 30 * time.Second
	monitorInterval               = 120 * time.Second
	metricResetInterval           = 30 * time.Minute
	reconcileSyncPeriod           = time.Hour

	// release result
	resultSucceeded = "Succeeded"
	resultFailed    = "Failed"

	// reconcile failure/retry reasons
	apiServerErr                 = "APIServerErr"
	applyComponentErr            = "ApplyComponentErr"
	regexpCompileErr             = "RegexpCompileErr"
	configurationMissingErr      = "ConfigurationMissingErr"
	getKeyVaultSecretErr         = "GetKeyVaultSecretErr" // #nosec only filed name with secret text
	getSecretProviderErr         = "GetSecretProviderErr"
	parseComponentConfigErr      = "ParseComponentConfigErr"
	parseSecretURLErr            = "ParseSecretURLErr" // #nosec only filed name with secret text
	placeholderNotAllReplacedErr = "PlaceholderNotAllReplacedErr"
	updateAnnotationsErr         = "UpdateAnnotationsErr"
	processUnmanagedSecretsErr   = "processUnmanagedSecretsErr"

	noRestartOnSecretUpdateStr = "noRestartOnSecretUpdate" // #nosec only filed name with secret text

	annotationPrefix                = "deployer.aks.io"
	ownerAnnotation                 = annotationPrefix + "/owner"
	secretAnnotationPrefix          = annotationPrefix + "/secret-"
	unmanagedSecretAnnotationPrefix = annotationPrefix + "/unmanaged-secret-"
)

var (
	useIdentityResourceID = false
)

func init() {
	// Check environment variables
	if os.Getenv(identityResourceIDStr) != "" {
		// Use identity resource ID if existed, otherwise fall back to SP
		useIdentityResourceID = true
		if os.Getenv(keyVaultResourceStr) == "" {
			klog.Fatalf("missing environment variable %s", keyVaultResourceStr)
		}
	} else {
		if os.Getenv(azureTenantIDStr) == "" ||
			os.Getenv(azureClientIDStr) == "" ||
			os.Getenv(azureClientSecretStr) == "" {
			klog.Fatal("missing environment variables for key vault client")
		}
	}
}

// AksAppReconciler reconciles an AksApp object
type AksAppReconciler struct {
	client.Client
	Logger    *logrus.Entry
	Scheme    *runtime.Scheme
	Namespace string

	UseOwnerReference bool

	rolloutRecheckInterval time.Duration
	reconcileBaseDelay     time.Duration
}

func NewAksAppReconciler(client client.Client, logger *logrus.Entry,
	scheme *runtime.Scheme, namespace string, useOwnerReference bool) *AksAppReconciler {
	return &AksAppReconciler{
		Client:                 client,
		Logger:                 logger,
		Scheme:                 scheme,
		Namespace:              namespace,
		UseOwnerReference:      useOwnerReference,
		rolloutRecheckInterval: defaultRolloutRecheckInterval,
		reconcileBaseDelay:     defaultReconcileBaseDelay,
	}
}

//+kubebuilder:rbac:groups=deployer.aks,resources=aksapps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=deployer.aks,resources=aksapps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=deployer.aks,resources=aksapps/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state according
// to the AksApp object referred to by the Request.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.10.0/pkg/reconcile
func (r *AksAppReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	operationID := uuid.NewV4().String()

	fields := map[string]interface{}{
		"aksapp":      fmt.Sprintf("%s/%s", req.Namespace, req.Name),
		"operationID": operationID,
	}

	logger := r.Logger.WithFields(fields)

	logger.Infof("Start reconciling AksApps %s", req.NamespacedName)
	var err error

	// 1. Get the AksApp component
	var app deployerv1.AksApp
	if err = r.Get(ctx, req.NamespacedName, &app); err != nil {
		if apierrors.IsNotFound(err) {
			logger.Infof("aksapp no longer exists, %s", err.Error())
			return ctrl.Result{}, nil // No requeue
		}
		logger.Errorf("unable to get aksapp, %s", err.Error())
		r.updateFailedReconciliation(app, apiServerErr, operationID, logger)
		return ctrl.Result{}, err
	}

	fields = map[string]interface{}{
		"appType":    app.Spec.Type,
		"appVersion": app.Spec.Version,
	}

	logger = logger.WithFields(fields)

	// 2. Retrieve the AksApp configuration
	var cm corev1.ConfigMap
	nn := types.NamespacedName{
		Namespace: r.Namespace,
		Name:      configmaps.GetAksAppConfigMapName(app),
	}
	if err = r.Get(ctx, nn, &cm); err != nil {
		if apierrors.IsNotFound(err) {
			logger.Errorf("aksapp configuration %s/%s does not exist", nn.Namespace, nn.Name)
			r.updateFailedReconciliation(app, configurationMissingErr, operationID, logger)
			return ctrl.Result{}, err
		}
		logger.Errorf("unable to get aksapp configuration: %s/%s", nn.Namespace, nn.Name)
		r.updateFailedReconciliation(app, apiServerErr, operationID, logger)
		return ctrl.Result{}, err
	}
	data := cm.Data["config"]

	// 3. Replace variables
	// Note: variable placeholders e.g. (V_XXX) will be replaced with real data
	//       per cluster configuration.
	for k, v := range app.Spec.Variables {
		keyPlaceHolder := "(V_" + k + ")"
		data = strings.ReplaceAll(data, keyPlaceHolder, v)
	}

	// 4. Replace credentials
	// Note: credential placeholders e.g. (V_XXX) will be replaced with key vault
	//       URLs that are retrieved automatically.
	var keyVaultSecretProvider secret.KeyvaultSecretProvider
	var keyVaultClient keyvault.BaseClient

	// TODO: Remove 'else' after fully migrating to identity resource ID
	if useIdentityResourceID {
		keyVaultSecretProvider, err = r.getKeyVaultSecretProvider()
		if err != nil {
			logger.Errorf("unable to get key vault secret provider, %s", err.Error())
			r.updateFailedReconciliation(app, getSecretProviderErr, operationID, logger)
			return ctrl.Result{}, err
		}
		logger.Infof("use managed service identity to get secret from key vault")
	} else {
		keyVaultClient, err = r.getKeyVaultClient()
		if err != nil {
			logger.Errorf("unable to get key vault client, %s", err.Error())
			r.updateFailedReconciliation(app, getSecretProviderErr, operationID, logger)
			return ctrl.Result{}, err
		}
		logger.Infof("use service principal to get secret from key vault")
	}

	// versioned: https://<vault-name>.vault.azure.net/secrets/<secret-name>/<secret-version>
	// latest: https://<vault-name>.vault.azure.net/secrets/<secret-name>
	validSecretURL, err := regexp.Compile("(https://[^/]+)/secrets/([\\w-]+)/?(\\w*)")
	if err != nil {
		logger.Errorf("unable to compile secret URL regexp, %s", err.Error())
		r.updateFailedReconciliation(app, regexpCompileErr, operationID, logger)
		return ctrl.Result{}, err
	}

	if app.Spec.Secrets == nil {
		app.Spec.Secrets = make(map[string]string)
	}

	secretAnnotations := map[string]string{}
	for key, sec := range app.Spec.Secrets {
		var secretBundle keyvault.SecretBundle
		var secretData string

		vaultData := validSecretURL.FindStringSubmatch(sec)
		if len(vaultData) == 0 {
			logger.Errorf("unable to parse secret URL: %s", sec)
			r.updateFailedReconciliation(app, parseSecretURLErr, operationID, logger)
			return ctrl.Result{}, err
		}

		// TODO: Remove 'else' after fully migrating to identity resource ID
		if useIdentityResourceID {
			// TODO: replace deployer log with rp logger
			logger := log.NewServiceLogger("Deployer", nil)
			secretBundle, err = keyVaultSecretProvider.Get(logger, sec)
		} else {
			// vaultData[1] - vaultBaseURL
			// vaultData[2] - secretName
			// vaultData[3] - secretVersion
			secretBundle, err = keyVaultClient.GetSecret(ctx, vaultData[1], vaultData[2], vaultData[3])
		}

		if err != nil {
			logger.Errorf("unable to get secret bundle %q, %s", sec, err.Error())
			r.updateFailedReconciliation(app, getKeyVaultSecretErr, operationID, logger)
			return ctrl.Result{}, err
		}

		if secretBundle.Value == nil {
			logger.Errorf("get nil secret in secret bundle %q", sec)
			r.updateFailedReconciliation(app, getKeyVaultSecretErr, operationID, logger)
			return ctrl.Result{}, err
		}

		logger.Infof("get secret bundle %s", *secretBundle.ID)
		secretData = *secretBundle.Value

		// Encode the secret with base64 and replace all the place holders
		keyPlaceHolder := "(V_" + key + ")"
		secretData = base64.StdEncoding.EncodeToString([]byte(secretData))
		data = strings.ReplaceAll(data, keyPlaceHolder, secretData)

		// Insert the entry deployer.aks.io/secret-<secret_key>: <secret_id>
		secretAnnotations[secretAnnotationPrefix+strings.ToLower(key)] = *secretBundle.ID
	}

	// 5.5 Process unmanaged secrets
	if err = r.processUnmanagedSecrets(ctx, secretAnnotations, &app, logger); err != nil {
		logger.Errorf("unable to process unmanaged secrets, %s", err.Error())
		r.updateFailedReconciliation(app, processUnmanagedSecretsErr, operationID, logger)
		return ctrl.Result{}, err
	}

	// 6. Check if there are sill (V_**) left
	// Regex match pattern: (V_*_-*)
	err = checkAfterReplace(data, logger)
	if err != nil {
		logger.Errorf("unable to pass the check after secret replacement, %s", err.Error())
		r.updateFailedReconciliation(app, placeholderNotAllReplacedErr, operationID, logger)
		return ctrl.Result{}, err
	}

	// 7. Parse the configuration data
	objs, err := configmaps.ParseConfigToUnstructured(logger, data)
	if err != nil {
		logger.Errorf("unable to parse component configuration, %s", err.Error())
		r.updateFailedReconciliation(app, parseComponentConfigErr, operationID, logger)
		return ctrl.Result{}, err
	}

	// 8.1 Update secret annotations for secrets
	if err := r.updateSecretAnnotations(secretAnnotations, objs, logger); err != nil {
		logger.Errorf("unable to update secret annotations, %s", err.Error())
		r.updateFailedReconciliation(app, parseComponentConfigErr, operationID, logger)
		return ctrl.Result{}, err
	}

	// 8.2 Deploy Secrets before processing other resources
	podAnnotations := map[string]string{}
	for _, obj := range objs {
		// Check if the resource is Secret
		if isV1Secret(obj) {
			if err := r.processUnstructuredObject(ctx, &app, obj, logger); err != nil {
				logger.Errorf("unable to process secrets for aksapp %s/%s, %s",
					app.Namespace, app.Name, err.Error())
				r.updateFailedReconciliation(app, applyComponentErr, operationID, logger)
				return ctrl.Result{}, err
			}

			// Insert the entry deployer.aks.io/secret-<secret_name>: <resource_version>
			resourceVersion := obj.GetResourceVersion()
			if resourceVersion == "" {
				logger.Errorf("get empty resource version from secret %s/%s",
					obj.GetNamespace(), obj.GetName())
			}
			logger.Infof("insert pod annotation secret %s with resource version %s",
				obj.GetName(), resourceVersion)
			podAnnotations[secretAnnotationPrefix+obj.GetName()] = resourceVersion
		}
	}

	// Do not update pods annotations if opted out
	// annotation: deployer.aks.io/noRestartOnSecretUpdate: true
	// TODO: Add a unit test to verify
	if value, found := app.ObjectMeta.GetAnnotations()[annotationPrefix+"/"+
		noRestartOnSecretUpdateStr]; found && strings.EqualFold(value, "true") {
		logger.Infof("skip updating pods annotations as %s annotation is set as 'true'", noRestartOnSecretUpdateStr)
	} else {
		// 8.3 Update secret annotations for pods
		if err := updatePodAnnotations(podAnnotations, objs, logger); err != nil {
			logger.Errorf("unable to update pod annotations, %s", err.Error())
			r.updateFailedReconciliation(app, updateAnnotationsErr, operationID, logger)
			return ctrl.Result{}, err
		}
	}

	// 9. Deploy
	currObjs := make([]*unstructured.Unstructured, len(objs))
	copy(currObjs, objs)
	var errObjs []*unstructured.Unstructured
	for len(currObjs) > 0 {
		for _, obj := range currObjs {
			// Skip Secret resource
			if isV1Secret(obj) {
				continue
			}

			if err := r.processUnstructuredObject(ctx, &app, obj, logger); err != nil {
				errObjs = append(errObjs, obj)
			}
		}

		if len(errObjs) == len(currObjs) {
			logger.Errorf("No more components can be processed for aksapp %s/%s, %d components with errors",
				app.Namespace, app.Name, len(errObjs))
			err = errors.New("unable to process more components in aksapp")
			r.updateFailedReconciliation(app, applyComponentErr, operationID, logger)
			return ctrl.Result{}, err
		}

		currObjs = errObjs
		errObjs = nil
	}

	logger.Infof("successfully processed aksapp component %s/%s", app.Namespace, app.Name)
	if err = r.updateSucceededReconciliation(&app, objs, operationID, logger); err != nil {
		logger.Errorf("unable to update status, %s", err.Error())
		return ctrl.Result{}, err
	}

	if app.Status.Rollout != deployerv1.RolloutCompleted {
		logger.Infof("rollout is in progress...")
		return ctrl.Result{
			RequeueAfter: r.rolloutRecheckInterval,
		}, nil
	}

	return ctrl.Result{}, nil
}

func isV1Secret(obj *unstructured.Unstructured) bool {
	secretGVK := schema.GroupVersionKind{Kind: "Secret", Version: "v1"}
	return obj.GroupVersionKind() == secretGVK
}

func isAppsV1Deployment(obj *unstructured.Unstructured) bool {
	deploymentGVK := schema.GroupVersionKind{Kind: "Deployment", Version: "v1", Group: "apps"}
	return obj.GroupVersionKind() == deploymentGVK
}

func (r *AksAppReconciler) processUnmanagedSecrets(ctx context.Context,
	secretAnnotations map[string]string,
	app *deployerv1.AksApp, logger *logrus.Entry) error {
	var err error
	for _, secretName := range app.Spec.UnmanagedSecrets {
		var secret corev1.Secret
		namespacedName := types.NamespacedName{
			Name:      secretName,
			Namespace: app.Namespace,
		}

		if err := r.Get(ctx, namespacedName, &secret); err != nil {
			if !apierrors.IsNotFound(err) {
				logger.Errorf("unable to get unmanaged secret %s, %s",
					namespacedName, err.Error())
				return err
			}
			logger.Warnf("unmanaged secret %s not found, %s",
				namespacedName, err.Error())
			continue
		}

		var fingerprint string
		var jsonData []byte
		if jsonData, err = json.Marshal(secret.Data); err != nil {
			logger.Errorf("unable to marshal unmanaged secret %s, %s",
				namespacedName, err.Error())
			return err
		}

		hash := sha256.New()
		if _, err = hash.Write([]byte(jsonData)); err != nil {
			logger.Errorf("unable to hash data of unmanaged secret %s, %s",
				namespacedName, err.Error())
			return err
		}
		result := hash.Sum(nil)
		fingerprint = hex.EncodeToString(result)

		secretAnnotations[unmanagedSecretAnnotationPrefix+secretName] = fingerprint
	}

	return nil
}

func (r *AksAppReconciler) updateSecretAnnotations(secretAnnotations map[string]string,
	objs []*unstructured.Unstructured, logger *logrus.Entry) error {
	ctx := context.TODO()
	for _, obj := range objs {
		// Check if the resource is Secret
		if isV1Secret(obj) {
			// Get the original annotations if the original secret exists
			secretKey := types.NamespacedName{
				Name:      obj.GetName(),
				Namespace: obj.GetNamespace(),
			}
			var secret corev1.Secret
			if err := r.Get(ctx, secretKey, &secret); err != nil {
				if !apierrors.IsNotFound(err) {
					logger.Errorf("unable to get secret %s/%s, %s",
						obj.GetNamespace(), obj.GetName(), err.Error())
					return err
				}
			}
			origAnnotations := secret.GetAnnotations()
			// Create a string map if the annotation is nil
			if origAnnotations == nil {
				origAnnotations = make(map[string]string)
			}

			currAnnotations := obj.GetAnnotations()
			if currAnnotations == nil {
				currAnnotations = make(map[string]string)
			}

			// Merge the secret annotations
			for k, v := range secretAnnotations {
				// Check if the value of the secret gets updated or if it is a new pair
				if orig, ok := origAnnotations[k]; ok {
					if v != orig {
						logger.Infof("update secret annotation %s value from %s to %s",
							k, orig, v)
					}
				} else {
					logger.Infof("add new secret annotation %s value %s",
						k, v)
				}
				currAnnotations[k] = v
			}

			// Set the annotations
			obj.SetAnnotations(currAnnotations)

			logger.Infof("updated annotations of the secret %s/%s",
				obj.GetNamespace(), obj.GetName())
		}
	}

	return nil
}

func updatePodAnnotations(podAnnotations map[string]string,
	objs []*unstructured.Unstructured, logger *logrus.Entry) error {
	for _, obj := range objs {
		// Check if the resource is to create containers
		_, found, err := unstructured.NestedFieldNoCopy(obj.Object,
			"spec", "template", "spec", "containers")
		if err != nil {
			logger.Errorf("unable to traverse %s object %s/%s, %s",
				obj.GetKind(), obj.GetNamespace(), obj.GetName(), err.Error())
			return err
		}

		// Skip unrelated resources
		if !found {
			continue
		}

		// Get the original annotations
		annotations, found, err := unstructured.NestedStringMap(obj.Object,
			"spec", "template", "metadata", "annotations")
		if err != nil {
			logger.Errorf("unable to traverse %s object %s/%s, %s",
				obj.GetKind(), obj.GetNamespace(), obj.GetName(), err.Error())
			return err
		}

		// Create a string map if the original annotation is not found
		if !found {
			annotations = make(map[string]string)
		}

		// Merge the secret annotations
		for k, v := range podAnnotations {
			annotations[k] = v
		}

		// Set the annotations
		if err := unstructured.SetNestedStringMap(obj.Object, annotations,
			"spec", "template", "metadata", "annotations"); err != nil {
			logger.Errorf("unable to set annotations of %s object %s/%s, %s",
				obj.GetKind(), obj.GetNamespace(), obj.GetName(), err.Error())
			return err
		}

		logger.Infof("updated annotations of %s object %s/%s",
			obj.GetKind(), obj.GetNamespace(), obj.GetName())
	}

	return nil
}

func setDeployerOwnerAnnotation(obj *unstructured.Unstructured, app *deployerv1.AksApp,
	logger *logrus.Entry) error {
	annotations := obj.GetAnnotations()
	if annotations == nil {
		annotations = make(map[string]string)
	}

	nn := types.NamespacedName{
		Namespace: app.Namespace,
		Name:      app.Name,
	}

	if val, ok := annotations[ownerAnnotation]; ok {
		if val != nn.String() && val != nn.Name {
			logger.Errorf("new owner %s does not match the old owner %s",
				nn.String(), val)
			return errors.New("conflict owners on the same resource")
		} else if val == nn.Name {
			// TODO: remove the logics after all the deprecated values are replaced
			logger.Infof("replace the deprecated name %s with namespaced name %s",
				val, nn.String())
		}
	}

	// Set annotation deployer.aks.io/owenr: <aksapp_namespace>/<aksapp_name>
	annotations[ownerAnnotation] = nn.String()
	obj.SetAnnotations(annotations)
	return nil
}

func isOwnedByAksApp(obj *unstructured.Unstructured) bool {
	ownerReferences := obj.GetOwnerReferences()
	for _, r := range ownerReferences {
		if r.Kind == deployerv1.AksAppKindStr {
			return true
		}
	}
	return false
}

func (r *AksAppReconciler) removeAksAppOwnerReferences(ctx context.Context,
	obj *unstructured.Unstructured, logger *logrus.Entry) error {
	// Remove AksApp owner references
	var ownerReferences []metav1.OwnerReference
	for _, r := range obj.GetOwnerReferences() {
		if r.Kind != deployerv1.AksAppKindStr {
			ownerReferences = append(ownerReferences, r)
		}
	}
	obj.SetOwnerReferences(ownerReferences)

	if err := r.Update(ctx, obj); err != nil {
		logger.Errorf("unable to remove owner reference from %s object %s/%s, %s",
			obj.GetKind(), obj.GetNamespace(), obj.GetName(), err.Error())
		return err
	}
	logger.Infof("successfully removed owner reference from %s object %s/%s",
		obj.GetKind(), obj.GetNamespace(), obj.GetName())
	return nil
}

func (r *AksAppReconciler) processUnstructuredObject(ctx context.Context, app *deployerv1.AksApp,
	obj *unstructured.Unstructured, logger *logrus.Entry) error {
	var err error
	if r.UseOwnerReference {
		// Set owner reference to the object
		obj.SetOwnerReferences([]metav1.OwnerReference{*metav1.NewControllerRef(
			app, deployerv1.GroupVersion.WithKind(deployerv1.AksAppKindStr)),
		})
	} else {
		if err = setDeployerOwnerAnnotation(obj, app, logger); err != nil {
			logger.Errorf("unable to set deployer owner annotation for %s object %s/%s for aksapp component %s/%s, %s",
				obj.GetKind(), obj.GetNamespace(), obj.GetName(), app.Namespace, app.Name, err.Error())
			return err
		}
	}

	tmp := obj.DeepCopyObject()
	objKey := types.NamespacedName{
		Name:      obj.GetName(),
		Namespace: obj.GetNamespace(),
	}
	if err = r.Get(ctx, objKey, tmp); err != nil {
		if apierrors.IsNotFound(err) {
			logger.Infof("create %s object %s/%s for aksapp component %s/%s",
				obj.GetKind(), obj.GetNamespace(), obj.GetName(), app.Namespace, app.Name)
			err = r.Create(ctx, obj)
		}
	} else {
		utmp := &unstructured.Unstructured{}
		if err = r.Scheme.Convert(tmp, utmp, nil); err != nil {
			logger.Errorf("unable to convert %s object %s/%s for aksapp component %s/%s to unstructured, %s",
				obj.GetKind(), obj.GetNamespace(), obj.GetName(), app.Namespace, app.Name, err.Error())
			return err
		}

		// Remove owner reference from the object if exists
		if isOwnedByAksApp(utmp) {
			logger.Info("remove existing AksApp owner references")
			if err = r.removeAksAppOwnerReferences(ctx, utmp, logger); err != nil {
				logger.Errorf("unable to remove AksApp owner reference from %s object %s/%s for aksapp component %s/%s, %s",
					obj.GetKind(), obj.GetNamespace(), obj.GetName(), app.Namespace, app.Name, err.Error())
				return err
			}
		}

		obj.SetResourceVersion(utmp.GetResourceVersion())
		if isV1Secret(obj) && reflect.DeepEqual(obj.GetAnnotations(), utmp.GetAnnotations()) {
			logger.Info("secret annotations remain the same; do not update secret resource version")
		} else {
			logger.Infof("patch %s object %s/%s for aksapp component %s/%s",
				obj.GetKind(), obj.GetNamespace(), obj.GetName(), app.Namespace, app.Name)
			err = r.Patch(ctx, obj, client.Merge)
		}
	}

	if err != nil {
		logger.Errorf("unable to process %s object %s/%s for aksapp component %s/%s, will retry later...",
			obj.GetKind(), obj.GetNamespace(), obj.GetName(), app.Namespace, app.Name)
		return err
	}

	logger.Infof("successfully processed %s object %s/%s for aksapp component %s/%s",
		obj.GetKind(), obj.GetNamespace(), obj.GetName(), app.Namespace, app.Name)

	return nil
}

func (r *AksAppReconciler) generationChangedPeriodicPredicate(e event.UpdateEvent) bool {
	if e.MetaOld == nil {
		return false
	}
	if e.ObjectOld == nil {
		return false
	}
	if e.ObjectNew == nil {
		return false
	}
	if e.MetaNew == nil {
		return false
	}

	if e.MetaNew.GetGeneration() != e.MetaOld.GetGeneration() {
		r.Logger.Infof("Object generation gets updated")
		return true
	}

	aksApp := &deployerv1.AksApp{}
	if utmp, err := runtime.DefaultUnstructuredConverter.ToUnstructured(e.ObjectOld); err != nil {
		r.Logger.Infof("unable to convert object %s/%s to unstructured, %s",
			e.MetaOld.GetNamespace(), e.MetaOld.GetName(), err.Error())
	} else {
		if err = runtime.DefaultUnstructuredConverter.FromUnstructured(utmp, aksApp); err != nil {
			r.Logger.Infof("unable to convert unstructured %s/%s to AksApp, %s",
				e.MetaOld.GetNamespace(), e.MetaOld.GetName(), err.Error())
		}
	}

	ret := metav1.Now().After(aksApp.Status.Reconciliation.LastReconcileTime.Add(reconcileSyncPeriod))
	if ret {
		r.Logger.Infof("The time till last reconciliation time exceeds the syncing interval")
	} else {
		r.Logger.Infof("Skip the same generation reconciliation within the syncing interval")
	}
	return ret
}

// SetupWithManager sets up AksApp manager
func (r *AksAppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	rateLimiter := workqueue.NewMaxOfRateLimiter(
		workqueue.NewItemExponentialFailureRateLimiter(r.reconcileBaseDelay, reconcileMaxDelay),
		// 10 qps, 100 bucket size.  This is only for retry speed and its only the overall factor (not per item)
		&workqueue.BucketRateLimiter{Limiter: rate.NewLimiter(rate.Limit(10), 100)},
	)

	// This allows a controller to reconcile an unchanged event no more frequent
	// than once per hour. For CustomResource objects the Generation is only
	// incremented when the status subresource is enabled.
	pred := predicate.Funcs{
		UpdateFunc: r.generationChangedPeriodicPredicate,
	}

	return ctrl.NewControllerManagedBy(mgr).
		For(&deployerv1.AksApp{}).
		WithOptions(k8scontroller.Options{RateLimiter: rateLimiter}).
		WithEventFilter(pred).
		Complete(r)
}

func (r *AksAppReconciler) isManagedByAksApp(d appsv1.Deployment) bool {
	if r.UseOwnerReference {
		for _, or := range d.OwnerReferences {
			if or.Kind == deployerv1.AksAppKindStr {
				return true
			}
		}
	} else {
		if _, ok := d.Annotations[ownerAnnotation]; ok {
			return true
		}
	}
	return false
}

func (r *AksAppReconciler) monitorUnavailableReplicas() {
	ctx := context.Background()
	fields := map[string]interface{}{
		"aksapp":      "monitorUnavailableReplicas",
		"operationID": uuid.NewV4().String(),
	}
	logger := r.Logger.WithFields(fields)

	var curDeploymentList appsv1.DeploymentList
	if err := r.List(ctx, &curDeploymentList); err != nil {
		logger.Errorf("unable to list deployments in all namespaces, %s", err.Error())
		return
	}

	unavailableReplicasVec.Reset()
	allReplicasVec.Reset()
	for _, d := range curDeploymentList.Items {
		if r.isManagedByAksApp(d) {
			unavailablePercentage := float64(0)
			if d.Status.Replicas > 0 {
				unavailablePercentage = float64(d.Status.UnavailableReplicas * 100 / d.Status.Replicas)
			}
			unavailableReplicasVec.WithLabelValues(d.Name, d.Namespace).Set(unavailablePercentage)
			allReplicasVec.WithLabelValues(d.Name, d.Namespace).Set(float64(d.Status.Replicas))
			if d.Status.UnavailableReplicas > 0 {
				logger.Infof("%d/%d unavailable replicas in deployment %s/%s",
					d.Status.UnavailableReplicas, d.Status.Replicas, d.Namespace, d.Name)
			}
		}
	}
}

func (r *AksAppReconciler) monitorServiceVersion() {
	ctx := context.Background()
	fields := map[string]interface{}{
		"aksapp":      "monitorServiceVersion",
		"operationID": uuid.NewV4().String(),
	}
	logger := r.Logger.WithFields(fields)

	var curAppList deployerv1.AksAppList
	if err := r.List(ctx, &curAppList); err != nil {
		logger.Errorf("unable to list aksapps in all namespaces, %s", err.Error())
		return
	}

	serviceVersionVec.Reset()
	for _, app := range curAppList.Items {
		serviceVersionVec.WithLabelValues(app.Name, app.Spec.Type, app.Spec.Version).Inc()
	}
}

func (r *AksAppReconciler) monitorSecretVersion() {
	ctx := context.Background()
	fields := map[string]interface{}{
		"aksapp":      "monitorSecretVersion",
		"operationID": uuid.NewV4().String(),
	}
	logger := r.Logger.WithFields(fields)

	var curSecretList corev1.SecretList
	if err := r.List(ctx, &curSecretList); err != nil {
		logger.Errorf("unable to list secrets in all namespaces, %s", err.Error())
		return
	}

	secretVersionVec.Reset()
	for _, secret := range curSecretList.Items {
		var appName string
		if r.UseOwnerReference {
			// Get name in OwnerReference
			for _, r := range secret.OwnerReferences {
				if r.Kind == deployerv1.AksAppKindStr {
					appName = r.Name
				}
			}
		} else {
			// Get name in NamespacedName
			annotations := secret.GetAnnotations()
			if _, ok := annotations[ownerAnnotation]; ok {
				appName = annotations[ownerAnnotation]
			}
		}

		if annotations := secret.GetAnnotations(); annotations != nil {
			for k, v := range annotations {
				if strings.HasPrefix(k, secretAnnotationPrefix) {
					secretVersionVec.WithLabelValues(appName, k, v).Set(1)
				}
			}
		}

	}
}

// MonitorRoutine monitors aksapp
func (r *AksAppReconciler) MonitorRoutine() {
	for {
		r.monitorServiceVersion()
		r.monitorSecretVersion()
		r.monitorUnavailableReplicas()
		time.Sleep(monitorInterval)
	}
}

// MetricResetRoutine resets metric
func (r *AksAppReconciler) MetricResetRoutine() {
	for {
		r.resetReleaseResult()
		time.Sleep(metricResetInterval)
	}
}

func (r *AksAppReconciler) getKeyVaultClient() (keyvault.BaseClient, error) {
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return keyvault.New(), err
	}
	keyVaultClient := keyvault.New()
	keyVaultClient.Authorizer = authorizer
	// keyVaultClient.Client.Sender
	return keyVaultClient, nil
}

func (r *AksAppReconciler) getKeyVaultSecretProvider() (secret.KeyvaultSecretProvider, error) {
	tokenProvider, err := tokenprovider.NewMsiTokenProvider(os.Getenv(identityResourceIDStr))
	if err != nil {
		return nil, err
	}
	return secret.NewKeyvaultSecretProviderFromTokenProvider(os.Getenv(keyVaultResourceStr), tokenProvider)
}

func (r *AksAppReconciler) updateRolloutStatus(app *deployerv1.AksApp,
	objs []*unstructured.Unstructured, logger *logrus.Entry) error {
	ctx := context.Background()
	rollouts := []deployerv1.Rollout{}
	rolloutStatus := deployerv1.RolloutCompleted
	var replicas int32
	var unavailableReplicas int32

	// Collect rollout status from all Deployments
	// TODO: collect rollout status from DaemonSets
	for _, obj := range objs {
		// Check if the resource is Deployment
		if isAppsV1Deployment(obj) {
			var deployment appsv1.Deployment
			namespacedName := types.NamespacedName{
				Name:      obj.GetName(),
				Namespace: obj.GetNamespace(),
			}

			// Get the Deployment
			if err := r.Get(ctx, namespacedName, &deployment); err != nil {
				logger.Errorf("unable to get deployment %s, %s",
					namespacedName, err.Error())
				return err
			}

			// Check the rollout status and collect replica numbers
			singleRolloutstatus := deployerv1.RolloutCompleted
			if !DeploymentComplete(&deployment) {
				singleRolloutstatus = deployerv1.RolloutInProgress
				rolloutStatus = deployerv1.RolloutInProgress
				logger.Warnf("deployment %s is not complete", namespacedName)
			}
			replicas += deployment.Status.Replicas
			unavailableReplicas += deployment.Status.UnavailableReplicas

			// Append the information
			rollouts = append(rollouts, deployerv1.Rollout{
				Name:                deployment.Name,
				Replicas:            deployment.Status.Replicas,
				Rollout:             singleRolloutstatus,
				UnavailableReplicas: deployment.Status.UnavailableReplicas,
			})
		}
	}

	// Log if the new status is different from the current status
	if rolloutStatus != app.Status.Rollout {
		logger.Infof("update aksapp %s/%s rollout status from %s to %s",
			app.Namespace, app.Name, app.Status.Rollout, rolloutStatus)
	}

	app.Status.Replicas = replicas
	app.Status.UnavailableReplicas = unavailableReplicas
	app.Status.RolloutVersion = app.Spec.Version
	app.Status.Rollout = rolloutStatus
	app.Status.Rollouts = rollouts

	return nil
}

func (r *AksAppReconciler) updateFailedReconciliation(app deployerv1.AksApp,
	reason, operationID string, logger *logrus.Entry) {
	// log release result
	r.logReleaseResult(app, resultFailed, reason, logger)

	// update askapp reconciliation status
	app.Status = deployerv1.AksAppStatus{
		Reconciliation: deployerv1.Reconciliation{
			LastReconcileTime: metav1.Now(),
			Message:           reason,
			OperationID:       operationID,
			Result:            deployerv1.ReconciliationFailed,
		},
	}

	ctx := context.TODO()
	if err := r.Status().Update(ctx, &app); err != nil {
		logger.Errorf("unable to update failed reconciliation status, %s", err.Error())
		// do not need to return error here since reconciliation will always be retried
	} else {
		logger.Info("updated failed reconciliation status")
	}
}

func (r *AksAppReconciler) updateSucceededReconciliation(app *deployerv1.AksApp,
	objs []*unstructured.Unstructured,
	operationID string, logger *logrus.Entry) error {
	// log release result
	r.logReleaseResult(*app, resultSucceeded, "", logger)

	// update aksapp reconciliation status
	app.Status.Reconciliation = deployerv1.Reconciliation{
		LastReconcileTime: metav1.Now(),
		Message:           "",
		OperationID:       operationID,
		Result:            deployerv1.ReconciliationSucceeded,
	}

	// update aksapp rollout status
	if err := r.updateRolloutStatus(app, objs, logger); err != nil {
		logger.Errorf("unable to update aksapp rollout status, %s", err.Error())
		return err
	}

	ctx := context.TODO()
	if err := r.Status().Update(ctx, app); err != nil {
		logger.Errorf("unable to update aksapp status, %s", err.Error())
		return err
	}

	logger.Info("updated succeeded reconciliation status")

	return nil
}

func (r *AksAppReconciler) logReleaseResult(app deployerv1.AksApp,
	result, reason string, logger *logrus.Entry) {
	fields := map[string]interface{}{
		"releaseResult": result,
		"reason":        reason,
	}
	logger = logger.WithFields(fields)
	logger.Infof("[release result] %s/%s: %q, reason: %q", app.Namespace, app.Name, result, reason)

	releaseResultVec.WithLabelValues(app.Namespace, app.Name, app.Spec.Type, result).Set(1)

	resultToReset := resultFailed
	if strings.EqualFold(result, resultFailed) {
		resultToReset = resultSucceeded
	}

	// set the other result to 0
	releaseResultVec.WithLabelValues(app.Namespace, app.Name, app.Spec.Type, resultToReset).Set(0)
}

func (r *AksAppReconciler) resetReleaseResult() {
	fields := map[string]interface{}{
		"aksapp":      "resetReleaseResult",
		"operationID": uuid.NewV4().String(),
	}
	logger := r.Logger.WithFields(fields)
	logger.Info("reset release result metric")
	releaseResultVec.Reset()
}

func checkAfterReplace(data string, logger *logrus.Entry) error {
	checkAfterReplace, err := regexp.Compile("\\(V_[a-zA-Z\\d_\\-]+\\)")
	if err != nil {
		logger.Errorf("unable to compile regexp for check, %s", err.Error())
		return err
	}

	allNotReplaced := checkAfterReplace.FindAllString(data, -1)
	if allNotReplaced != nil {
		logger.Errorf("There are still (V_*) left not replaced: %s", strings.Join(allNotReplaced, ","))
		return errors.New("deployer place holders left not replaced")
	}
	return nil
}
