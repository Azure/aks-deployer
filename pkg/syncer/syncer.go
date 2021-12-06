package syncer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/sirupsen/logrus"

	"github.com/Azure/aks-deployer/pkg/clients/blobclient"
	deployerv1 "github.com/Azure/aks-deployer/pkg/api/v1"
	"github.com/Azure/aks-deployer/pkg/configmaps"
)

const (
	clusterContainerName = "cluster"
	aksAppContainerName  = "aksapp"

	clusterBlobNameFormat = "%s.yaml"
	aksAppBlobNameFormat  = "%s/%s.yaml"
)

// Options syncer options
type Options struct {
	ToggleURL string
	Namespace string
}

// Syncer syncs configuration from remote to local ConfigMaps
type Syncer struct {
	kubeClient            kubernetes.Interface
	storageClient         storage.Client
	blobClient            blobclient.Interface
	storageAccountName    string
	storageEndpointSuffix string
	clusterName           string
	options               Options
	logger                *logrus.Entry
}

// NewSyncer returns a Syncer
func NewSyncer(
	kubeClient kubernetes.Interface,
	storageClient storage.Client,
	clusterName string,
	logger *logrus.Entry) *Syncer {

	syncer := &Syncer{
		kubeClient:    kubeClient,
		storageClient: storageClient,
		clusterName:   clusterName,
		logger:        logger,
	}

	logger.Info("Create a new Syncer using stoage client")
	return syncer
}

// NewSyncerWithBlobClient returns a Syncer using BlobClient
func NewSyncerWithBlobClient(
	kubeClient kubernetes.Interface,
	blobClient blobclient.Interface,
	storageAccountName string,
	storageEndpointSuffix string,
	clusterName string,
	logger *logrus.Entry) *Syncer {

	syncer := &Syncer{
		kubeClient:            kubeClient,
		blobClient:            blobClient,
		storageAccountName:    storageAccountName,
		storageEndpointSuffix: storageEndpointSuffix,
		clusterName:           clusterName,
		logger:                logger,
	}

	logger.Info("Create a new Syncer using blob client")
	return syncer
}

// SetOptions sets the options for syncer
func (s *Syncer) SetOptions(opt Options) {
	s.options = opt
}

// Run runs the syncer routine
func (s *Syncer) Run() {
	s.logger.Infof("Syncing %s cluster configuration...", s.clusterName)

	// Fetch cluster configuration
	clusterBlobName := fmt.Sprintf(clusterBlobNameFormat, s.clusterName)
	// Check the folder of cluster name
	body, err := s.getAzureStorageBlobsWithPrefix(clusterContainerName, s.clusterName+"/")
	if err != nil {
		s.logger.Warnf("Failed to get configuration files in cluster folder, %s", err.Error())
		return
	}
	if body == "" {
		s.logger.Warnf("No configuration files found in cluster folder")
		return
	}

	scheme := runtime.NewScheme()
	_ = deployerv1.AddToScheme(scheme)
	objs, err := configmaps.ParseConfig(s.logger, scheme, body)
	if err != nil {
		s.logger.Errorf("Failed to parse cluster configuration %s, %s",
			configmaps.ClusterConfigMapName, err.Error())
		return
	}

	// Save new ConfigMaps to avoid being cleaned up
	configMapTypeMap := make(map[string]string)

	// Create new ConfigMaps
	for _, obj := range objs {
		app := obj.(*deployerv1.AksApp)
		s.logger.Infof("Syncing %s:%s configuration...", app.Spec.Type, app.Spec.Version)
		aksAppBlobName := fmt.Sprintf(aksAppBlobNameFormat, app.Spec.Type, app.Spec.Version)
		body, err := s.getAzureStorageBlob(aksAppContainerName, aksAppBlobName)
		if err != nil {
			s.logger.Warningf("Failed to fetch aksapp configuration, %s", err.Error())
			continue
		}

		configMapName := configmaps.GetAksAppConfigMapName(*app)
		configMapTypeMap[app.Spec.Type] = configMapName

		_, err = s.createOrUpdateConfigMap(s.newConfigMap(configMapName, body, aksAppBlobName))
		if err != nil {
			s.logger.Errorf("Failed to create/update aksapp configuration, %s", err.Error())
			continue
		}
	}

	// Create/Update the cluster ConfigMap after creating the component
	// configurations to avoid configurationMissingErr during reconciliation.
	_, err = s.createOrUpdateConfigMap(s.newConfigMap(configmaps.ClusterConfigMapName, body, clusterBlobName))
	if err != nil {
		s.logger.Errorf("Failed to create/update cluster configuration, %s", err.Error())
	}

	// Clean up the obsolete ConfigMaps at last to avoid
	// configurationMissingErr when reconciling the old AksApps.
	err = s.cleanUpObsoleteConfigMaps(configMapTypeMap)
	if err != nil {
		s.logger.Errorf("Failed to cleanup obsolete configmaps: %s", err.Error())
	}
}

func (s *Syncer) getAzureStorageBlobViaBlobClient(containerName, blobName string) (string, error) {
	ctx := context.TODO()
	url := fmt.Sprintf("https://%s.blob.%s/%s/%s",
		s.storageAccountName, s.storageEndpointSuffix, containerName, blobName)

	data, err := s.blobClient.GetBlobData(ctx, url)
	if err != nil {
		s.logger.Errorf("Failed to get container %s blob %s via blob client, %s",
			containerName, blobName, err.Error())
		return "", err
	}
	// GetBlobData doesn't return error when 404 occurs and return nil instead
	if data == nil {
		s.logger.Errorf("Container %s blob %s does not exist via blob client",
			containerName, blobName)
		return "", errors.New("Get a non-existent blob")
	}

	s.logger.Infof("Successfully get data from container %s blob %s",
		containerName, blobName)

	return string(data), nil
}

func (s *Syncer) getAzureStorageBlobViaBlobStorageClient(containerName, blobName string) (string, error) {
	blobStorageClient := s.storageClient.GetBlobService()

	blob := blobStorageClient.GetContainerReference(containerName).GetBlobReference(blobName)
	s.logger.Infof("Get blob form URL %s", blob.GetURL())

	resp, err := blob.Get(&storage.GetBlobOptions{})
	if err != nil {
		s.logger.Warnf("Failed to get container %s blob %s from Azure Storage: %s",
			containerName, blobName, err.Error())
		return "", err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp)
	if err != nil {
		s.logger.Warnf("Failed to read data from response, %s", err.Error())
		return "", err
	}

	return buf.String(), nil
}

func (s *Syncer) getAzureStorageBlob(containerName, blobName string) (string, error) {
	// Use blob client if not nil
	if s.blobClient != nil {
		return s.getAzureStorageBlobViaBlobClient(containerName, blobName)
	}

	return s.getAzureStorageBlobViaBlobStorageClient(containerName, blobName)
}

func (s *Syncer) getAzureStorageBlobsWithPrefixViaBlobClient(containerName, prefix string) (string, error) {
	ctx := context.TODO()
	url := fmt.Sprintf("https://%s.blob.%s/%s",
		s.storageAccountName, s.storageEndpointSuffix, containerName)

	blobList, err := s.blobClient.ListBlobsWithPrefix(ctx, url, prefix)
	if err != nil {
		s.logger.Errorf("Failed to list container %s blobs with prefix %s via blob client, %s",
			containerName, prefix, err.Error())
		return "", err
	}

	s.logger.Infof("Successfully list blobs in container %s with prefix %s",
		containerName, prefix)

	var body string
	for _, blob := range blobList.Blobs {
		blobBody, err := s.getAzureStorageBlob(containerName, blob.Name)
		if err != nil {
			s.logger.Errorf("Failed to get one blob of configurations %s, %s",
				blob.Name, err.Error())
			return "", err
		}
		body = body + blobBody
		body = body + "---\n"
	}

	return body, nil
}

func (s *Syncer) getAzureStorageBlobsWithPrefixViaBlobStorageClient(containerName, prefix string) (string, error) {
	blobStorageClient := s.storageClient.GetBlobService()

	blobList, err := blobStorageClient.GetContainerReference(containerName).ListBlobs(storage.ListBlobsParameters{
		Prefix: prefix,
	})
	if err != nil {
		s.logger.Warnf("Failed to list container %s blobs with prefix %s from Azure Storage, %s",
			containerName, prefix, err.Error())
		return "", err
	}

	s.logger.Infof("Get blob list with prefix %s", prefix)
	var body string
	for _, blob := range blobList.Blobs {
		s.logger.Infof("Get blob %s", blob.Name)
		blobBody, err := s.getAzureStorageBlob(containerName, blob.Name)
		if err != nil {
			s.logger.Warnf("Failed to get a piece of configuration %s, %s", blob.Name, err.Error())
			return "", err
		}
		body = body + blobBody
		body = body + "---\n"
	}

	return body, nil
}

func (s *Syncer) getAzureStorageBlobsWithPrefix(containerName, prefix string) (string, error) {
	// Use blob client if not nil
	if s.blobClient != nil {
		return s.getAzureStorageBlobsWithPrefixViaBlobClient(containerName, prefix)
	}

	return s.getAzureStorageBlobsWithPrefixViaBlobStorageClient(containerName, prefix)
}

func (s *Syncer) newConfigMap(name, data, blobURL string) *corev1.ConfigMap {
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: s.options.Namespace,
			Annotations: map[string]string{
				configmaps.ConfigAnnotationKey: blobURL,
			},
		},
		Data: map[string]string{
			configmaps.ConfigDataKey: data,
		},
	}
}

// TODO(shuche): Add mechanism to clean up ConfigMaps
func (s *Syncer) createOrUpdateConfigMap(configmap *corev1.ConfigMap) (*corev1.ConfigMap, error) {
	// Get specific namespace set to the syncer
	namespace := s.options.Namespace
	ctx := context.TODO()

	var cm *corev1.ConfigMap
	_, err := s.kubeClient.CoreV1().ConfigMaps(namespace).Get(ctx, configmap.Name, metav1.GetOptions{})
	if apierrors.IsNotFound(err) {
		s.logger.Infof("Creating ConfigMap:%s...", configmap.Name)
		cm, err = s.kubeClient.CoreV1().ConfigMaps(namespace).Create(ctx, configmap, metav1.CreateOptions{})
	} else {
		s.logger.Infof("Updating ConfigMap:%s...", configmap.Name)
		cm, err = s.kubeClient.CoreV1().ConfigMaps(namespace).Update(ctx, configmap, metav1.UpdateOptions{})
	}

	if err != nil {
		return nil, err
	}

	return cm, nil
}

// Cleanup obsolete configmaps after configmap and aksapp are created
func (s *Syncer) cleanUpObsoleteConfigMaps(configMapTypeMap map[string]string) error {
	namespace := s.options.Namespace
	ctx := context.TODO()
	configMapList, err := s.kubeClient.CoreV1().ConfigMaps(namespace).List(ctx, metav1.ListOptions{})

	if err != nil {
		s.logger.Errorf("Error when list configmaps: %s", err.Error())
		return err
	}

	for _, c := range configMapList.Items {
		name := c.GetName()
		if s.isConfigMapObsolete(name, configMapTypeMap) {
			s.logger.Infof("Cleanup obsolete configmap %s", name)
			err = s.kubeClient.CoreV1().ConfigMaps(namespace).Delete(ctx, name, metav1.DeleteOptions{})
			if err != nil {
				s.logger.Errorf("Error when delete configmap %s: %s", name, err.Error())
				return err
			}
			s.logger.Infof("Successfully delete configmap %s", name)
		}
	}

	return nil
}

func (s *Syncer) isConfigMapObsolete(configMapName string, configMapTypeMap map[string]string) bool {
	// 1. Found aksapp type for give configMapName by finding longest prefix match in aksapp type list
	// To avoid conflict when finding aksapp since there will be privatecluster and privatecluster-ccp-proxy both exist in the aksapp type list
	matchAksAppType := ""
	for aksappType := range configMapTypeMap {
		if strings.HasPrefix(configMapName, aksappType) {
			if len(aksappType) > len(matchAksAppType) {
				matchAksAppType = aksappType
			}
		}
	}
	// 2. Check if the configMapName is the same with current valid configMap name
	if matchAksAppType == "" {
		s.logger.Warningf("Configmap %s has not found matching aksapp", configMapName)
		return false
	}
	s.logger.Infof("Found configmaps of aksapp %s: %s - %s", matchAksAppType, configMapName, configMapTypeMap[matchAksAppType])
	return !strings.EqualFold(configMapTypeMap[matchAksAppType], configMapName)
}
