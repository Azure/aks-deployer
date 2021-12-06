package syncer

import (
	"context"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-sdk-for-go/storage"
	"github.com/sirupsen/logrus"
	"github.com/Azure/aks-deployer/pkg/clients/blobclient/mock_blobclient"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/Azure/aks-deployer/pkg/log"
)

var _ = Describe("Test Syncer with Blob Client", func() {
	var (
		logger     *logrus.Entry
		kubeClient *fake.Clientset
		blobClient *mock_blobclient.MockInterface
		syncer     *Syncer
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		logger = log.New("test-service", "test-version")

		kubeClient = fake.NewSimpleClientset()
		blobClient = mock_blobclient.NewMockInterface(mockCtrl)
		syncer = NewSyncerWithBlobClient(kubeClient, blobClient,
			"test-storage-account-name", "test-storage-endpoint-suffix", "test-cluster-name", logger)
	})

	It("Test getAzureStorageBlob with BlobClient with empty bytes", func() {
		blobClient.EXPECT().GetBlobData(gomock.Any(), gomock.Any()).Times(1).Return([]byte(""), nil)
		_, err := syncer.getAzureStorageBlob("test-container-name", "test-blob-name")
		Expect(err).To(BeNil())
	})

	It("Test getAzureStorageBlob with BlobClient with nil (404)", func() {
		blobClient.EXPECT().GetBlobData(gomock.Any(), gomock.Any()).Times(1).Return(nil, nil)
		_, err := syncer.getAzureStorageBlob("test-container-name", "test-blob-name")
		Expect(err).NotTo(BeNil())
	})

	It("Test getAzureStorageBlobsWithPrefix with BlobClient", func() {
		blobClient.EXPECT().ListBlobsWithPrefix(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(storage.BlobListResponse{}, nil)
		_, err := syncer.getAzureStorageBlobsWithPrefix("test-container-name", "test-blob-name")
		Expect(err).To(BeNil())
	})

	It("Test getAzureStorageBlobsWithPrefix with BlobClient with one blob", func() {
		blobList := storage.BlobListResponse{
			Blobs: []storage.Blob{
				{
					Name: "blob-1",
				},
			},
		}
		content := []byte("blob-1-data\n")

		blobClient.EXPECT().GetBlobData(gomock.Any(), gomock.Any()).Times(1).Return(content, nil)
		blobClient.EXPECT().ListBlobsWithPrefix(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(blobList, nil)

		data, err := syncer.getAzureStorageBlobsWithPrefix("test-container-name", "test-blob-name")
		Expect(err).To(BeNil())
		Expect(data).To(Equal(string(content) + "---\n"))
	})

	It("Test getAzureStorageBlobsWithPrefix with BlobClient with multiple blobs", func() {
		blobList := storage.BlobListResponse{
			Blobs: []storage.Blob{
				{
					Name: "blob-1",
				},
				{
					Name: "blob-2",
				},
			},
		}
		content1 := []byte("blob-1-data\n")
		content2 := []byte("blob-2-data\n")

		blobClient.EXPECT().GetBlobData(gomock.Any(), gomock.Any()).Times(1).Return(content1, nil)
		blobClient.EXPECT().GetBlobData(gomock.Any(), gomock.Any()).Times(1).Return(content2, nil)
		blobClient.EXPECT().ListBlobsWithPrefix(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(blobList, nil)

		data, err := syncer.getAzureStorageBlobsWithPrefix("test-container-name", "test-blob-name")
		Expect(err).To(BeNil())
		Expect(data).To(Equal(string(content1) + "---\n" + string(content2) + "---\n"))
	})

	It("Test getAzureStorageBlobsWithPrefix with BlobClient with empty data", func() {
		blobList := storage.BlobListResponse{
			Blobs: []storage.Blob{
				{
					Name: "blob-1",
				},
				{
					Name: "blob-2",
				},
			},
		}
		var content1 []byte
		var content2 []byte

		blobClient.EXPECT().GetBlobData(gomock.Any(), gomock.Any()).Times(1).Return(content1, nil)
		blobClient.EXPECT().GetBlobData(gomock.Any(), gomock.Any()).Times(1).Return(content2, nil)
		blobClient.EXPECT().ListBlobsWithPrefix(gomock.Any(), gomock.Any(), gomock.Any()).Times(1).Return(blobList, nil)

		_, err := syncer.getAzureStorageBlobsWithPrefix("test-container-name", "test-blob-name")
		Expect(err).NotTo(BeNil())
	})
})

var _ = Describe("Test Syncer clean up obsolete configmaps", func() {
	var (
		logger           *logrus.Entry
		kubeClient       *fake.Clientset
		blobClient       *mock_blobclient.MockInterface
		syncer           *Syncer
		ctx              context.Context
		configMapTypeMap map[string]string
	)

	BeforeEach(func() {
		mockCtrl := gomock.NewController(GinkgoT())
		logger = log.New("test-service", "test-version")
		kubeClient = fake.NewSimpleClientset()
		blobClient = mock_blobclient.NewMockInterface(mockCtrl)
		syncer = NewSyncerWithBlobClient(kubeClient, blobClient,
			"test-storage-account-name", "test-storage-endpoint-suffix", "test-cluster-name", logger)
		opt := Options{
			Namespace: "test-namespace",
		}
		syncer.SetOptions(opt)
		kubeClient.CoreV1().ConfigMaps("test-namespace").Create(ctx, syncer.newConfigMap("aksapp-test", "test", "test-blob-url"), metav1.CreateOptions{})
		kubeClient.CoreV1().ConfigMaps("test-namespace").Create(ctx, syncer.newConfigMap("aksapp-123-test", "test", "test-blob-url"), metav1.CreateOptions{})
		configMapTypeMap = make(map[string]string)
		configMapTypeMap["aksapp"] = "aksapp-test"
		configMapTypeMap["aksapp-123"] = "aksapp-123-test"
	})

	It("Test cleanUpObsoleteConfigMaps with configmaps to delete obsolete configmaps", func() {
		kubeClient.CoreV1().ConfigMaps("test-namespace").Create(ctx, syncer.newConfigMap("aksapp-test123", "test", "test-blob-url"), metav1.CreateOptions{})
		err := syncer.cleanUpObsoleteConfigMaps(configMapTypeMap)
		Expect(err).To(BeNil())
		listAfterCleanup, _ := kubeClient.CoreV1().ConfigMaps("test-namespace").List(ctx, metav1.ListOptions{})
		Expect(len(listAfterCleanup.Items)).To(Equal(2))
		for _, c := range listAfterCleanup.Items {
			Expect(c.GetName).NotTo(Equal("aksapp-test123"))
		}
		kubeClient.CoreV1().ConfigMaps("test-namespace").Create(ctx, syncer.newConfigMap("aksapp-123-test123", "test", "test-blob-url"), metav1.CreateOptions{})
		err = syncer.cleanUpObsoleteConfigMaps(configMapTypeMap)
		Expect(err).To(BeNil())
		listAfterCleanup, _ = kubeClient.CoreV1().ConfigMaps("test-namespace").List(ctx, metav1.ListOptions{})
		Expect(len(listAfterCleanup.Items)).To(Equal(2))
		for _, c := range listAfterCleanup.Items {
			Expect(c.GetName).NotTo(Equal("aksapp-123-test123"))
		}
	})

	It("Test cleanUpObsoleteConfigMaps and no need to cleanup", func() {
		err := syncer.cleanUpObsoleteConfigMaps(configMapTypeMap)
		Expect(err).To(BeNil())
		listAfterCleanup, _ := kubeClient.CoreV1().ConfigMaps("test-namespace").List(ctx, metav1.ListOptions{})
		Expect(len(listAfterCleanup.Items)).To(Equal(2))
	})

	It("Test cleanUpObsoleteConfigMaps and with no matching aksapp", func() {
		kubeClient.CoreV1().ConfigMaps("test-namespace").Create(ctx, syncer.newConfigMap("aks321-test123", "test", "test-blob-url"), metav1.CreateOptions{})
		err := syncer.cleanUpObsoleteConfigMaps(configMapTypeMap)
		Expect(err).To(BeNil())
		listAfterCleanup, _ := kubeClient.CoreV1().ConfigMaps("test-namespace").List(ctx, metav1.ListOptions{})
		Expect(len(listAfterCleanup.Items)).To(Equal(3))
	})
})
