package configmaps

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/sirupsen/logrus"

	"k8s.io/apimachinery/pkg/runtime"

	deployerv1 "github.com/Azure/aks-deployer/pkg/api/v1"
)

var _ = Describe("Test Parse", func() {
	var (
		logger *logrus.Entry
		scheme *runtime.Scheme
	)

	BeforeEach(func() {
		logger = logrus.NewEntry(logrus.New())
		scheme = runtime.NewScheme()
		_ = deployerv1.AddToScheme(scheme)
	})

	It("Parse empty configuration", func() {
		configData := ""

		objs, err := ParseConfig(logger, scheme, configData)
		Expect(objs).To(BeNil())
		Expect(err).To(BeNil())
	})

	It("Parse one configuration", func() {
		configData := `apiVersion: deployer.aks/v1
kind: AksApp
metadata:
  name: ccp-pool-controller
  namespace: deployer
spec:
  type: ccp-pool-controller
  version: ${CCP_POOL_CONTROLLER_VERSION}
  variables:
    # Default value set as "0"
    CCP_POOL_CAPACITY: "0"
    REPLICAS: "3"
    TOGGLE_URL: ${TOGGLE_URL}
  credentials:
    AZURE: ${AZURE_CRED_SECRET}
    DOCKERCFG: ${DOCKER_CFG_SECRET}`

		objs, err := ParseConfig(logger, scheme, configData)
		Expect(objs).NotTo(BeNil())
		Expect(err).To(BeNil())

		for _, obj := range objs {
			app := obj.(*deployerv1.AksApp)
			Expect(app).NotTo(BeNil())
			Expect(app.Name).To(Equal("ccp-pool-controller"))
		}
	})

	It("Parse a wrong configuration", func() {
		configData := `apiVersion: deployer.aks/v1
metadata:
  name: ccp-pool-controller
  namespace: deployer`

		objs, err := ParseConfig(logger, scheme, configData)
		Expect(objs).To(BeNil())
		Expect(err).NotTo(BeNil())
	})

	It("Parse two configurations", func() {
		configData := `apiVersion: deployer.aks/v1
kind: AksApp
metadata:
  name: ccp-pool-controller
  namespace: deployer
spec:
  type: ccp-pool-controller
  version: ${CCP_POOL_CONTROLLER_VERSION}
  variables:
    # Default value set as "0"
    CCP_POOL_CAPACITY: "0"
    REPLICAS: "3"
    TOGGLE_URL: ${TOGGLE_URL}
  credentials:
    AZURE: ${AZURE_CRED_SECRET}
    DOCKERCFG: ${DOCKER_CFG_SECRET}
  ---
apiVersion: deployer.aks/v1
kind: AksApp
metadata:
  name: overlay-manager
  namespace: overlay-manager
spec:
  type: overlay-manager
  version: ${CCP_POOL_CONTROLLER_VERSION}
  variables:
      REPLICAS: "3"
      TOGGLE_URL: ${TOGGLE_URL}
  credentials:
    AZURE: ${AZURE_CRED_SECRET}
    DOCKERCFG: ${DOCKER_CFG_SECRET}`

		objs, err := ParseConfig(logger, scheme, configData)
		Expect(objs).NotTo(BeNil())
		Expect(err).To(BeNil())

		app1 := objs[0].(*deployerv1.AksApp)
		Expect(app1).NotTo(BeNil())
		Expect(app1.Name).To(Equal("ccp-pool-controller"))

		app2 := objs[1].(*deployerv1.AksApp)
		Expect(app2).NotTo(BeNil())
		Expect(app2.Name).To(Equal("overlay-manager"))
	})
})
