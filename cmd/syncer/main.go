package main

import (
	"context"
	"flag"
	"os"
	"time"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/Azure/aks-deployer/pkg/clients/blobclient"
	"github.com/Azure/aks-deployer/pkg/version"
	"github.com/Azure/aks-deployer/pkg/configmaps"
	"github.com/Azure/aks-deployer/pkg/leader"
	"github.com/Azure/aks-deployer/pkg/log"
	"github.com/Azure/aks-deployer/pkg/signals"
	"github.com/Azure/aks-deployer/pkg/syncer"

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/sirupsen/logrus"
)

var (
	toggleURLFlag string
	namespace     string
	logger        *logrus.Entry

	useIdentityResourceID = false
	storageEndpointSuffix = "core.windows.net"
)

const (
	syncTimeInterval         = 5 * time.Minute
	identityResourceIDStr    = "IDENTITY_RESOURCE_ID"
	storageEndpointSuffixStr = "STORAGE_ENDPOINT_SUFFIX"
)

func init() {
	flag.StringVar(&toggleURLFlag, "toggle-url", "", "toggle URL")
	flag.StringVar(&namespace, "namespace", configmaps.DefaultDeployerNamespace, "namespace")

	// logger setup
	// note: source field is used in fluentd to match rules and add tags
	logger = log.New("syncer", version.String()).WithField("source", "Deployer")

	if os.Getenv("AZURE_STORAGE_ACCOUNT") == "" {
		logger.Fatal("Azure storage account environment variables not set, exiting...")
	}

	if os.Getenv("CLUSTER_NAME") == "" {
		logger.Fatal("CLUSTER_NAME not set, exiting...")
	}

	if os.Getenv(identityResourceIDStr) != "" {
		// Use identity resource ID if existed, otherwise fall back to use access key
		useIdentityResourceID = true
		if os.Getenv(storageEndpointSuffixStr) != "" {
			storageEndpointSuffix = os.Getenv(storageEndpointSuffixStr)
			logger.Infof("Set storage endpoint suffix with %s", storageEndpointSuffix)
		}
	} else {
		if os.Getenv("AZURE_STORAGE_ACCESS_KEY") == "" {
			logger.Fatal("Azure storage credential environment variables not set, exiting...")
		}
	}
}

func main() {
	flag.Parse()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	// get the inClusterConfig
	cfg, err := clientcmd.BuildConfigFromFlags("", "")
	if err != nil {
		logger.Fatalf("Error building in-cluster kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logger.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	clusterName := os.Getenv("CLUSTER_NAME")

	var s *syncer.Syncer
	if useIdentityResourceID {
		storageAccountName := os.Getenv("AZURE_STORAGE_ACCOUNT")
		logger.Infof("Use identity resource ID to access storage account %s", storageAccountName)

		blobClient, err := blobclient.NewMsiBlobProviderWithMiResourceID(os.Getenv(identityResourceIDStr))
		if err != nil {
			logger.Fatalf("Error creating blob client, %s", err.Error())
		}
		s = syncer.NewSyncerWithBlobClient(
			kubeClient, blobClient, storageAccountName, storageEndpointSuffix, clusterName, logger)
	} else {
		storageAccountName := os.Getenv("AZURE_STORAGE_ACCOUNT")
		storageAccessKey := os.Getenv("AZURE_STORAGE_ACCESS_KEY")

		storageClient, err := storage.NewBasicClient(storageAccountName, storageAccessKey)
		if err != nil {
			logger.Fatalf("Error creating storage client: %s", err.Error())
		}

		s = syncer.NewSyncer(kubeClient, storageClient, clusterName, logger)
	}

	opt := syncer.Options{
		ToggleURL: toggleURLFlag,
		Namespace: namespace,
	}
	s.SetOptions(opt)

	run := func(ctx context.Context) {
		s.Run()

		for {
			select {
			case <-time.After(syncTimeInterval):
				s.Run()
			case <-stopCh:
				logger.Info("Shutting down...")
				os.Exit(0)
			}
		}
	}

	leader.RunWithLeaderElection(run, logger, namespace, "syncer")
}
