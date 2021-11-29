// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package controllers

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/klog"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	deployerv1 "github.com/Azure/aks-deployer/pkg/api/v1"
	"github.com/Azure/aks-deployer/pkg/configmaps"
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
func (r *AksAppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
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

	// TODO: port remaining code

	return ctrl.Result{}, nil
}

func (r *AksAppReconciler) updateFailedReconciliation(app deployerv1.AksApp,
	reason, operationID string, logger *logrus.Entry) {
	// TODO: add actual code.
}

// SetupWithManager sets up the controller with the Manager.
func (r *AksAppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&deployerv1.AksApp{}).
		Complete(r)
}
