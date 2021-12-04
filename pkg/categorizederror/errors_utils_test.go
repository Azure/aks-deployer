package categorizederror

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/stretchr/testify/assert"

	"github.com/Azure/aks-deployer/pkg/apierror"
	"github.com/Azure/aks-deployer/pkg/log"
)

func TestGetDependencyFromError(t *testing.T) {
	t.Run("ADAL dependency", func(t *testing.T) {
		err := errors.New("azure.BearerAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/33e6d93e-bf9c-402d-8c3e-1d40eed2f9a8/resourceGroups/hcp-underlay-southafricanorth/providers/Microsoft.Network/dnsZones/southafricanorth.azmk8s.io/A/e2edns-e2emtrcs-21021204-pqr-cd069758.hcp?api-version=2017-10-01: StatusCode=0 -- Original Error: adal: Failed to execute the refresh request. Error = 'Post \"https://login.microsoftonline.com/33e01921-4d64-4f8c-a055-5bdaffd5e33d/oauth2/token?api-version=1.0\": EOF'")
		dep := getDependencyFromError(err)

		assert.Equal(t, ADAL, dep)
	})

	t.Run("ADAL dependency for azure.multiTenantSPTAuthorizer", func(t *testing.T) {
		err := errors.New("azure.multiTenantSPTAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/33e6d93e-bf9c-402d-8c3e-1d40eed2f9a8/resourceGroups/hcp-underlay-southafricanorth/providers/Microsoft.Network/dnsZones/southafricanorth.azmk8s.io/A/e2edns-e2emtrcs-21021204-pqr-cd069758.hcp?api-version=2017-10-01: StatusCode=0 -- Original Error: adal: Failed to execute the refresh request. Error = 'Post \"https://login.microsoftonline.com/33e01921-4d64-4f8c-a055-5bdaffd5e33d/oauth2/token?api-version=1.0\": EOF'")
		dep := getDependencyFromError(err)

		assert.Equal(t, ADAL, dep)
	})

	t.Run("ARM dependency", func(t *testing.T) {
		err := errors.New("Code=\"PublicIPCountLimitReached\" Message=\"Cannot create more than 10 public IP addresses for this subscription in this region.\" Details=[]")
		dep := getDependencyFromError(err)

		assert.Equal(t, ARM, dep)
	})
}

func TestGetCategoryAndSubCodeFromError(t *testing.T) {
	t.Run("EOF error", func(t *testing.T) {

		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("this will cause EOF error at client")
		}))

		client := &http.Client{}
		request, _ := http.NewRequest("Get", server.URL, nil)
		_, err := client.Do(request)
		cat, subCode := getCategoryAndSubCodeFromError(nil, err)

		assert.Equal(t, EOF, subCode)
		assert.Equal(t, apierror.InternalError, cat)
	})

	t.Run("context cancel error", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))

		client := &http.Client{}
		request, _ := http.NewRequest("Get", server.URL, nil)
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		_, err := client.Do(request.WithContext(ctx))

		cat, subCode := getCategoryAndSubCodeFromError(nil, err)

		assert.Equal(t, ContextCanceled, subCode)
		assert.Equal(t, apierror.InternalError, cat)
	})

	t.Run("connectivity error", func(t *testing.T) {
		err := errors.New("Get \"https://management.azure.com/subscriptions/54aea2c5-5855-40ac-a5e7-5f3252640fae/resourceGroups/MC_aksrnr-f5bd2271_agentpoolCluster_westus/providers/Microsoft.Compute/virtualMachineScaleSets/aks-pool1-49601223-vmss?api-version=2020-06-01\": i/o timeout")
		cat, subCode := getCategoryAndSubCodeFromError(nil, err)

		assert.Equal(t, IOTimedout, subCode)
		assert.Equal(t, apierror.InternalError, cat)
	})

	t.Run("500 error", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 500,
		}

		err := errors.New("autorest/azure: Service returned an error. Status=500 Code=\"InternalServerError\" Message=\"Encountered internal server error. Diagnostic information: timestamp '20210213T041159Z', subscription id '54aea2c5-5855-40ac-a5e7-5f3252640fae', tracking id '27345547-4861-4376-806e-5f0612a2a003', request correlation id '27345547-4861-4376-806e-5f0612a2a003'.\"")
		cat, subCode := getCategoryAndSubCodeFromError(resp, err)

		assert.Equal(t, ErrorSubCode("InternalServerError"), subCode)
		assert.Equal(t, apierror.InternalError, cat)
	})

	t.Run("vmss extension error", func(t *testing.T) {
		err := errors.New("VMSSAgentPoolReconciler retry failed: deployment operations failed with error messages: {\"code\": \"VMExtensionProvisioningError\",\"message\": \"VM has reported a failure when processing extension 'vmssCSE'. Error message: \"Enable failed: failed to execute command: command terminated with exit status=50\n[stdout]\nThu Oct 31 05:53:51 UTC 2019,aks-default-92520957-vmss000000\n\n[stderr]\n\".\",}")
		cat, subCode := getCategoryAndSubCodeFromError(nil, err)

		assert.Equal(t, ErrorSubCode("OutboundConnFailVMExtensionError"), subCode)
		assert.Equal(t, apierror.ClientError, cat)

		svErr := &azure.ServiceError{
			Code:    "VMExtensionProvisioningError",
			Message: "VM has reported a failure when processing extension 'vmssCSE'. Error message: \"Enable failed: failed to execute command: command terminated with exit status=52\n[stdout]\n3-04 23:54:11 UTC; 20s ago\\n Main PID: 4093 (kubelet)\\n Tasks: 37 (limit: 4915)\\n CGroup: /system.slice/kubelet.service\\n ├─3941 /opt/cni/bin/azure-vnet-telemetry -d /opt/cni/bin\\n └─4093 /usr/local/bin/kubelet --enable-server --node-labels=kubernetes.azure.com/role=agent,agentpool=r360aksdev,storageprofile=managed,storagetier=Premium_LRS,kubernetes.azure.com/cluster=MC_R360_DEV_R360-dev-aks_southeastasia,kubernetes.azure.com/mode=system,kubernetes.azure.com/node-image-version=AKSUbuntu-1804-2020.12.15 --v=2 --volume-plugin-dir=/etc/kubernetes/volumeplugins --address=0.0.0.0 --anonymous-auth=false --authentication-token-webhook=true --authorization-mode=Webhook --azure-container-registry-config=/etc/kubernetes/azure.json --cgroups-per-qos=true --client-ca-file=/etc/kubernetes/certs/ca.crt --cloud-config=/etc/kubernetes/azure.json --cloud-provider=azure --cluster-dns=10.0.0.10 --cluster-domain=cluster.local --dynamic-config-dir=/var/lib/kubelet --enforce-node-allocatable=pods --event-qps=0 --eviction-hard=memory.available\u003c750Mi,nodefs.available\u003c10%,nodefs.inodesFree\u003c5% --feature-gates=RotateKubeletServerCertificate=true --image-gc-high-threshold=85 --image-gc-low-threshold=80 --image-pull-progress-deadline=30m --keep-terminated-pod-volumes=false --kube-reserved=cpu=180m,memory=3645Mi --kubeconfig=/var/lib/kubelet/kubeconfig --max-pods=30 --network-plugin=cni --node-status-update-frequency=10s --non-masquerade-cidr=0.0.0.0/0 --pod-infra-container-image=mcr.microsoft.com/oss/kubernetes/pause:1.3.1 --pod-manifest-path=/etc/kubernetes/manifests --pod-max-pids=-1 --protect-kernel-defaults=true --read-only-port=0 --resolv-conf=/run/systemd/resolve/resolv.conf --rotate-certificates=false --streaming-connection-idle-timeout=4h --tls-cert-file=/etc/kubernetes/certs/kubeletserver.crt --tls-cipher-suites=TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_256_GCM_SHA384,TLS_RSA_WITH_AES_128_GCM_SHA256 --tls-private-key-file=/etc/kubernetes/certs/kubeletserver.key\\n\\nMar 04 23:54:31 aks-r360aksdev-39858457-vmss00001G kubelet[4093]: E0304 23:54:31.131207 4093 kubelet.go:2270] node \\\"aks-r360aksdev-39858457-vmss00001g\\\" not found\\nMar 04 23:54:31 aks-r360aksdev-39858457-vmss00001G kubelet[4093]: E0304 23:54:31.231346 4093 kubelet.go:2270] node \\\"aks-r360aksdev-39858457-vmss00001g\\\" not found\\nMar 04 23:54:31 aks-r360aksdev-39858457-vmss00001G kubelet[4093]: E0304 23:54:31.331477 4093 kubelet.go:2270] node \\\"aks-r360aksdev-39858457-vmss00001g\\\" not found\\nMar 04 23:54:31 aks-r360aksdev-39858457-vmss00001G kubelet[4093]: I0304 23:54:31.396396 4093 csi_plugin.go:945] Failed to contact API server when waiting for CSINode publishing: Get https://r360-dev-aks-dns-94cf63ca.hcp.southeastasia.azmk8s.io:443/apis/storage.k8s.io/v1/csinodes/aks-r360aksdev-39858457-vmss00001g: dial tcp: lookup r360-dev-aks-dns-94cf63ca.hcp.southeastasia.azmk8s.io: no such host\\nMar 04 23:54:31 aks-r360aksdev-39858457-vmss00001G kubelet[4093]: E0304 23:54:31.431594 4093 kubelet.go:2270] node \\\"aks-r360aksdev-39858457-vmss00001g\\\" not found\\nMar 04 23:54:31 aks-r360aksdev-39858457-vmss00001G kubelet[4093]: E0304 23:54:31.531715 4093 kubelet.go:2270] node \\\"aks-r360aksdev-39858457-vmss00001g\\\" not found\\nMar 04 23:54:31 aks-r360aksdev-39858457-vmss00001G kubelet[4093]: E0304 23:54:31.631846 4093 kubelet.go:2270] node \\\"aks-r360aksdev-39858457-vmss00001g\\\" not found\\nMar 04 23:54:31 aks-r360aksdev-39858457-vmss00001G kubelet[4093]: E0304 23:54:31.731975 4093 kubelet.go:2270] node \\\"aks-r360aksdev-39858457-vmss00001g\\\" not found\\nMar 04 23:54:31 aks-r360aksdev-39858457-vmss00001G kubelet[4093]: E0304 23:54:31.832156 4093 kubelet.go:2270] node \\\"aks-r360aksdev-39858457-vmss00001g\\\" not found\\nMar 04 23:54:31 aks-r360aksdev-39858457-vmss00001G kubelet[4093]: E0304 23:54:31.932337 4093 kubelet.go:2270] node \\\"aks-r360aksdev-39858457-vmss00001g\\\" not found\", \"Error\": \"\" }\n\n[stderr]\n\"\r\n\r\nMore information on troubleshooting is available at https://aka.ms/VMExtensionCSELinuxTroubleshoot ",
		}
		cat, subCode = getCategoryAndSubCodeFromError(nil, svErr)

		assert.Equal(t, ErrorSubCode("K8SAPIServerDNSLookupFailVMExtensionError"), subCode)
		assert.Equal(t, apierror.ClientError, cat)

		err = errors.New("VMSSAgentPoolReconciler retry failed: deployment operations failed with error messages: {\"code\": \"VMExtensionProvisioningError\",\"message\": \"VM has reported a failure when processing extension 'MMAExtension'. Error message: \"Enable failed with exit code 53 Installation failed due to incorrect workspace key. Please check if the workspace key is correct. For details, check logs in /var/log/azure/Microsoft.EnterpriseCloud.Monitoring.OmsAgentForLinux/extension.log\"\r\n\r\nMore information on troubleshooting is available at https://aka.ms/VMExtensionOMSAgentLinuxTroubleshoot ")
		cat, subCode = getCategoryAndSubCodeFromError(nil, err)
		assert.Equal(t, ErrorSubCode("VMExtensionProvisioningError"), subCode)
		assert.Equal(t, apierror.ClientError, cat)

		err = errors.New("VMSSAgentPoolReconciler retry failed: deployment operations failed with error messages: {\"code\": \"VMExtensionProvisioningError\",\"message\": \"VM has reported a failure when processing extension 'vmssCSE'. Error message: \"Command execution finished, but failed because it returned a non-zero exit code of: '1'\"\r\n\r\nMore information on troubleshooting is available at https://aka.ms/VMExtensionCSEWindowsTroubleshoot ")
		cat, subCode = getCategoryAndSubCodeFromError(nil, err)
		assert.Equal(t, ErrorSubCode("VMExtensionProvisioningError_Windows"), subCode)
		assert.Equal(t, apierror.InternalError, cat)
	})

	t.Run("vmas extension error", func(t *testing.T) {
		err := errors.New("Code=\"VMExtensionProvisioningError\" Message=\"VM has reported a failure when processing extension 'cse-agent-13'. Error message: \"Enable failed: failed to execute command: command terminated with exit status=50\n[stdout]\n\n[stderr]\nnc: connect to mcr.microsoft.com port 443 (tcp) failed: Connection timed out\nCommand exited with non-zero status 1\n0.00user 0.00system 2:09.97elapsed 0%CPU (0avgtext+0avgdata 2260maxresident)k\n0inputs+8outputs (0major+112minor)pagefaults 0swaps\n\"\r\n\r\nMore information on troubleshooting is available at https://aka.ms/VMExtensionCSELinuxTroubleshoot \"")
		cat, subCode := getCategoryAndSubCodeFromError(nil, err)

		assert.Equal(t, ErrorSubCode("OutboundConnFailVMExtensionError"), subCode)
		assert.Equal(t, apierror.ClientError, cat)

		err = errors.New("deployment operations failed with error messages: {\"code\": \"VMExtensionProvisioningError\",\"message\": \"VM has reported a failure when processing extension 'MMAExtension'. Error message: \"Enable failed with exit code 53 Installation failed due to incorrect workspace key. Please check if the workspace key is correct. For details, check logs in /var/log/azure/Microsoft.EnterpriseCloud.Monitoring.OmsAgentForLinux/extension.log\"\r\n\r\nMore information on troubleshooting is available at https://aka.ms/VMExtensionOMSAgentLinuxTroubleshoot ")
		cat, subCode = getCategoryAndSubCodeFromError(nil, err)
		assert.Equal(t, ErrorSubCode("VMExtensionProvisioningError"), subCode)
		assert.Equal(t, apierror.ClientError, cat)
	})

	t.Run("vmss provision error", func(t *testing.T) {
		err := errors.New("VMSSAgentPoolReconciler retry failed: VMSS 'vmsstest' reported failure details={instance 1 has no instance views available, vmssInstanceErrorCode=NoVMSSInstanceView;}")
		cat, subCode := getCategoryAndSubCodeFromError(nil, err)

		assert.Equal(t, ErrorSubCode("NoVMSSInstanceView"), subCode)
		assert.Equal(t, apierror.InternalError, cat)
	})

	t.Run("vmss extension handler non transient error", func(t *testing.T) {
		err := errors.New("Code=\"VMExtensionHandlerNonTransientError\" Message=\"The handler for VM extension type 'Microsoft.EnterpriseCloud.Monitoring.OmsAgentForLinux' has reported terminal failure for VM extension 'MMAExtension' with error message: '[ExtensionOperationError] Non-zero exit code: 11, /var/lib/waagent/Microsoft.EnterpriseCloud.Monitoring.OmsAgentForLinux-1.13.40/omsagent_shim.sh -install\n[stdout]\n2021/09/30 11:53:54 [Microsoft.EnterpriseCloud.Monitoring.OmsAgentForLinux-1.13.40] Install,failed,11,Install failed due to an invalid parameter: Workspace ID is invalid")
		cat, subCode := getCategoryAndSubCodeFromError(nil, err)

		assert.Equal(t, ErrorSubCode("VMExtensionHandlerNonTransientError"), subCode)
		assert.Equal(t, apierror.ClientError, cat)

		err = errors.New("Code=\"VMExtensionHandlerNonTransientError\" Message=\"The handler for VM extension type 'Microsoft.Azure.Monitor.AzureMonitorLinuxAgent' has reported terminal failure for VM extension 'Microsoft.Azure.Monitor.AzureMonitorLinuxAgent' with error message: '[ExtensionOperationError] Non-zero exit code: 53, /var/lib/waagent/Microsoft.Azure.Monitor.AzureMonitorLinuxAgent-1.12.2/./shim.sh -install\n[stdout]\n2021/09/29 00dpkg  and  azuremonitoragent_1.12.2-build.master.260_x86_64.deb\nnstall,failed,53,Failed to add MCS Environment Variables in /etc/default/azuremonitoragent\n\n\n[stderr]\nRunning scope as unit: Microsoft.Azure.Monitor.AzureMonitorLinuxAgent_1.12.2_0cc1c241-8e8e-4159-a56e-c5c7f327f79c.scope\ndpkg: error: dpkg frontend is locked by another process")
		cat, subCode = getCategoryAndSubCodeFromError(nil, err)

		assert.Equal(t, ErrorSubCode("VMExtensionHandlerNonTransientError"), subCode)
		assert.Equal(t, apierror.InternalError, cat)

		err = errors.New("Code=\"VMExtensionHandlerNonTransientError\" Message=\"The handler for VM extension type 'Microsoft.Azure.KeyVault.KeyVaultForWindows' has reported terminal failure for VM extension 'Microsoft.Azure.KeyVault.KeyVaultForWindows' with error message: '[ExtensionOperationError] Non-zero exit code: 53, /var/lib/waagent/Microsoft.Azure.KeyVault.KeyVaultForWindows-1.12.2/./shim.sh -install\n[stdout]\n2021/09/29 00dpkg  failed,53,Failed to add MCS Environment Variables in /etc/default/.scope\ndpkg: error: dpkg frontend is locked by another process")
		cat, subCode = getCategoryAndSubCodeFromError(nil, err)

		assert.Equal(t, ErrorSubCode("VMExtensionHandlerNonTransientError_Windows"), subCode)
		assert.Equal(t, apierror.InternalError, cat)
	})

	t.Run("get code from serviceError", func(t *testing.T) {
		svErr := &azure.ServiceError{
			Code:    "VMStartTimedOut",
			Message: "VM 'aks-agentpool0-10470690-vmss_2' did not start in the allotted time. The VM may still start successfully. Please check the power state later.",
		}

		cat, subCode := getCategoryAndSubCodeFromError(nil, svErr)

		assert.Equal(t, ErrorSubCode("VMStartTimedOut"), subCode)
		assert.Equal(t, apierror.InternalError, cat)

		svErr = &azure.ServiceError{
			Code:    "StorageFailure/SocketException",
			Message: "Error while preparing to create storage object https://md-btdvg5nnh4zg.z6.blob.storage.azure.net/c3prfqcrhvcd/abcd  Target: '/subscriptions/762ae2ab-1cc6-405d-a518-24410dbaef6c/resourceGroups/MC_e2erg-e2emtrcs-21022316-FgFT4_e2eaks-iBQ_eastus2euap/providers/Microsoft.Compute/disks/aks-agentpool0-18133aks-agentpool0-181335OS__1_7ebfda5e65dc40cdbe380748f7bc0f71'.",
		}

		cat, subCode = getCategoryAndSubCodeFromError(nil, svErr)

		assert.Equal(t, ErrorSubCode("StorageFailure/SocketException"), subCode)
		assert.Equal(t, apierror.InternalError, cat)

	})

	t.Run("get code from serviceError inside request error", func(t *testing.T) {
		svErr := &azure.ServiceError{
			Code:    "OperationNotAllowed",
			Message: "\"The server rejected the request because too many requests have been received for this subscription.\"",
		}

		requestErr := &azure.RequestError{
			ServiceError: svErr,
		}

		resp := &http.Response{
			StatusCode: 409,
		}

		cat, subCode := getCategoryAndSubCodeFromError(resp, requestErr)

		assert.Equal(t, ErrorSubCode("OperationNotAllowed"), subCode)
		assert.Equal(t, apierror.ClientError, cat)
	})

	t.Run("get code from serviceError inside request error", func(t *testing.T) {
		svErr := &azure.ServiceError{
			Code:    "InvalidTemplateDeployment",
			Message: "The template deployment 'agent-21-06-09T23.42.44-12829964' is not valid according to the validation procedure. The tracking id is '6034f949-16f5-468c-9b3b-8c903a291384'. See inner errors for details.",
			Details: []map[string]interface{}{
				map[string]interface{}{
					"code": "QuotaExceeded",
				},
			},
		}

		requestErr := &azure.RequestError{
			ServiceError: svErr,
		}

		resp := &http.Response{
			StatusCode: 400,
		}

		cat, subCode := getCategoryAndSubCodeFromError(resp, requestErr)

		assert.Equal(t, ErrorSubCode("InvalidTemplateDeployment_QuotaExceeded"), subCode)
		assert.Equal(t, apierror.ClientError, cat)
	})

	t.Run("get code from serviceError inside request error, ignore the same code", func(t *testing.T) {
		svErr := &azure.ServiceError{
			Code:    "InvalidTemplateDeployment",
			Message: "The template deployment 'agent-21-06-09T23.42.44-12829964' is not valid according to the validation procedure. The tracking id is '6034f949-16f5-468c-9b3b-8c903a291384'. See inner errors for details.",
			Details: []map[string]interface{}{
				map[string]interface{}{
					"code": "InvalidTemplateDeployment",
				},
			},
		}

		requestErr := &azure.RequestError{
			ServiceError: svErr,
		}

		resp := &http.Response{
			StatusCode: 400,
		}

		cat, subCode := getCategoryAndSubCodeFromError(resp, requestErr)

		assert.Equal(t, ErrorSubCode("InvalidTemplateDeployment"), subCode)
		assert.Equal(t, apierror.ClientError, cat)
	})

	t.Run("subcode is empty in service error details", func(t *testing.T) {
		svErr := &azure.ServiceError{
			Code:    "InvalidTemplateDeployment",
			Message: "The template deployment 'agent-21-06-09T23.42.44-12829964' is not valid according to the validation procedure. The tracking id is '6034f949-16f5-468c-9b3b-8c903a291384'. See inner errors for details.",
			Details: []map[string]interface{}{
				map[string]interface{}{
					"notcode": "someotherdetails",
				},
			},
		}

		requestErr := &azure.RequestError{
			ServiceError: svErr,
		}

		resp := &http.Response{
			StatusCode: 400,
		}

		cat, subCode := getCategoryAndSubCodeFromError(resp, requestErr)

		assert.Equal(t, ErrorSubCode("InvalidTemplateDeployment"), subCode)
		assert.Equal(t, apierror.ClientError, cat)
	})
}

func TestGetCategoryAndSubCodeFromADALError(t *testing.T) {
	t.Run("EOF error", func(t *testing.T) {
		err := fmt.Errorf("azure.BearerAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/b57894d0-9e3c-4123-9bde-61b3c2889a37/providers/Microsoft.Compute?api-version=2019-08-01: StatusCode=0 -- Original Error: adal: Failed to execute the refresh request. Error = 'Post \"https://login.microsoftonline.com/68e770fa-f42f-412c-bce0-9e4c228045f4/oauth2/token?api-version=1.0\": EOF'")
		cat, subCode := getCategoryAndSubCodeFromADALError(err)

		assert.Equal(t, EOF, subCode)
		assert.Equal(t, apierror.InternalError, cat)
	})

	t.Run("contextCanceled", func(t *testing.T) {
		err := fmt.Errorf("azure.BearerAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/91abb4fd-9c89-45ee-9108-cbd7a3affbb7/providers/Microsoft.Compute?api-version=2019-08-01: StatusCode=0 -- Original Error: adal: Failed to execute the refresh request. Error = 'Post \"https://login.microsoftonline.com/9148d540-88e0-4508-9710-873898a483f1/oauth2/token?api-version=1.0\": context canceled'")
		cat, subCode := getCategoryAndSubCodeFromADALError(err)

		assert.Equal(t, ContextCanceled, subCode)
		assert.Equal(t, apierror.InternalError, cat)
	})

	t.Run("service unavaiable", func(t *testing.T) {
		err := fmt.Errorf("azure.BearerAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/1a099dea-d42b-4376-9207-a38e91798f46/providers/Microsoft.Network?api-version=2019-08-01: StatusCode=503 -- Original Error: adal: Refresh request failed. Status Code = '503'. Response body: The service is unavailable. Endpoint https://login.microsoftonline.com/d8bbd0f9-4a38-4818-ace4-79356c30d69d/oauth2/token?api-version=1.0")
		cat, subCode := getCategoryAndSubCodeFromADALError(err)

		assert.Equal(t, ErrorSubCode("ServiceUnavailable"), subCode)
		assert.Equal(t, apierror.InternalError, cat)
	})

	t.Run("AADSTS700016 applicaiont not found", func(t *testing.T) {
		err := fmt.Errorf("azure.BearerAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/b4adf4bf-1609-42aa-b86e-31dad729b971/resourceGroups/rg-privatelink.canadacentral.azmk8s.io-vzyx/providers/Microsoft.Network/privateDnsZones/privatelink.canadacentral.azmk8s.io?api-version=2018-09-01: StatusCode=400 -- Original Error: adal: Refresh request failed. Status Code = '400'. Response body: {\"error\":\"unauthorized_client\",\"error_description\":\"AADSTS700016: Application with identifier 'd2fee990-4a1c-45f0-b72e-7838146e17f0' was not found in the directory 'f1821f56-8d20-4902-bf65-bc7502e367e7'. This can happen if the application has not been installed by the administrator of the tenant or consented to by any user in the tenant. You may have sent your authentication request to the wrong tenant.\r\nTrace ID: 16621d64-ce75-42eb-b13c-e5549da06800\r\nCorrelation ID: a3945bef-08c1-46c8-9883-a385ab8a0d59\r\nTimestamp: 2021-03-09 19:34:55Z\",\"error_codes\":[700016],\"timestamp\":\"2021-03-09 19:34:55Z\",\"trace_id\":\"16621d64-ce75-42eb-b13c-e5549da06800\",\"correlation_id\":\"a3945bef-08c1-46c8-9883-a385ab8a0d59\",\"error_uri\":\"https://login.microsoftonline.com/error?code=700016\"} Endpoint https://login.microsoftonline.com/f1821f56-8d20-4902-bf65-bc7502e367e7/oauth2/token?api-version=1.0")
		cat, subCode := getCategoryAndSubCodeFromADALError(err)

		assert.Equal(t, ErrorSubCode("AADSTS700016"), subCode)
		assert.Equal(t, apierror.ErrorCategory(""), cat)
	})

	t.Run("extra status code from error msg", func(t *testing.T) {
		err := fmt.Errorf("azure.multiTenantSPTAuthorizer#WithAuthorization: Failed to refresh one or more Tokens for request to https://management.azure.com/subscriptions/a5adcb22-9d5c-4a8c-935f-e0bc583f6302/resourceGroups/MC_CP-CBB-CLC01-rg_akosaks01_westeurope/providers/Microsoft.Compute/virtualMachineScaleSets/aks-agentpool-92533872-vmss?api-version=2020-12-01:  StatusCode=0 -- Original Error: failed to refresh primary token: adal: Refresh request failed. Status Code = '504'. Response body: Endpoint https://login.microsoftonline.com/da823195-14f1-4c80-b72d-8e6d18cf5a0a/oauth2/token?api-version=1.0")
		cat, subCode := getCategoryAndSubCodeFromADALError(err)

		assert.Equal(t, ErrorSubCode("GatewayTimeout"), subCode)
		assert.Equal(t, apierror.InternalError, cat)
	})
}

func TestHandleErrorToCategorizedError(t *testing.T) {
	ctx := log.NewContextWithTeam(context.Background(), log.AKSTeamRP)
	t.Run("with RequestError", func(t *testing.T) {
		err := &azure.RequestError{
			DetailedError: autorest.DetailedError{
				Original: errors.New("RequestDisallowedByPolicy"),
				Method:   "GET",
			},
			ServiceError: &azure.ServiceError{
				Code:    "RequestDisallowedByPolicy",
				Message: "resource 'f7b40021-7ddc-4589-8d85-8a2b0e2fed6f' was disallowed by policy",
			},
		}
		cerr := HandleErrorToCategorizedError(ctx, nil, err)
		assert.Equal(t, ARM, cerr.Dependency)
		assert.Equal(t, RequestDisallowedByPolicy, cerr.SubCode)
		assert.Equal(t, apierror.ClientError, cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)
	})

	t.Run("with detailedError", func(t *testing.T) {
		err := &autorest.DetailedError{
			Original: errors.New("DiskEncryptionSet"),
			Method:   "GET",
		}

		cerr := HandleErrorToCategorizedError(ctx, nil, err)
		assert.Equal(t, ARM, cerr.Dependency)
		assert.Equal(t, DiskEncryptionSetError, cerr.SubCode)
		assert.Equal(t, apierror.ClientError, cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)
	})

	t.Run("with ServiceError", func(t *testing.T) {
		err := &azure.ServiceError{
			Code:    "Conflict",
			Message: "Operation group '/operations/groups/id/|virtualNetworkLinks|3f807987-8acd-4b30-8bbf-16ccb4804602|mc_e2erg-e2emtrcs-21021205-stolr_e2eaks-ugs_centraluseuap|a3abae6b-93e1-4b0f-8cf7-847414c92172.privatelink.centraluseuap.azmk8s.io|e2edns-e2emtrcs-21021205-xqj-aa20b93a' already has 1 operations like '/operations/type/UpsertVirtualNetworkLink/id/b86a2625-9a4d-42f0-af27-cec172157797' queued.",
		}

		cerr := HandleErrorToCategorizedError(ctx, nil, err)
		assert.Equal(t, ARM, cerr.Dependency)
		assert.Equal(t, Conflict, cerr.SubCode)
		assert.Equal(t, apierror.InternalError, cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)
	})

	t.Run("with adalError", func(t *testing.T) {
		err := fmt.Errorf("azure.BearerAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/b57894d0-9e3c-4123-9bde-61b3c2889a37/providers/Microsoft.Compute?api-version=2019-08-01: StatusCode=0 -- Original Error: adal: Failed to execute the refresh request. Error = 'Post \"https://login.microsoftonline.com/68e770fa-f42f-412c-bce0-9e4c228045f4/oauth2/token?api-version=1.0\": EOF'")
		cerr := HandleErrorToCategorizedError(ctx, nil, err)

		assert.Equal(t, ADAL, cerr.Dependency)
		assert.Equal(t, EOF, cerr.SubCode)
		assert.Equal(t, apierror.InternalError, cerr.Category)

		err = fmt.Errorf("azure.BearerAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/465d8b4d-4d84-4514-a5c3-6c59d97391e3/resourceGroups/MC_e2erg-e2emtrcs-21030919-zMqb3_e2eaks-dCZ_koreasouth/providers/Microsoft.Network/loadBalancers/kubernetes?api-version=2018-11-01: StatusCode=0 -- Original Error: adal: Failed to execute the refresh request. Error = 'Post \"https://login.microsoftonline.com/72f988bf-86f1-41af-91ab-2d7cd011db47/oauth2/token?api-version=1.0\": read tcp 172.31.23.6:45410->40.126.35.144:443: read: connection reset by peer")
		cerr = HandleErrorToCategorizedError(ctx, nil, err)
		assert.Equal(t, ADAL, cerr.Dependency)
		assert.Equal(t, ConnectionResetByPeer, cerr.SubCode)
		assert.Equal(t, apierror.InternalError, cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)

		err = fmt.Errorf("azure.BearerAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/b4adf4bf-1609-42aa-b86e-31dad729b971/resourceGroups/rg-privatelink.canadacentral.azmk8s.io-vzyx/providers/Microsoft.Network/privateDnsZones/privatelink.canadacentral.azmk8s.io?api-version=2018-09-01: StatusCode=400 -- Original Error: adal: Refresh request failed. Status Code = '400'. Response body: {\"error\":\"unauthorized_client\",\"error_description\":\"AADSTS700016: Application with identifier 'd2fee990-4a1c-45f0-b72e-7838146e17f0' was not found in the directory 'f1821f56-8d20-4902-bf65-bc7502e367e7'. This can happen if the application has not been installed by the administrator of the tenant or consented to by any user in the tenant. You may have sent your authentication request to the wrong tenant.\r\nTrace ID: 16621d64-ce75-42eb-b13c-e5549da06800\r\nCorrelation ID: a3945bef-08c1-46c8-9883-a385ab8a0d59\r\nTimestamp: 2021-03-09 19:34:55Z\",\"error_codes\":[700016],\"timestamp\":\"2021-03-09 19:34:55Z\",\"trace_id\":\"16621d64-ce75-42eb-b13c-e5549da06800\",\"correlation_id\":\"a3945bef-08c1-46c8-9883-a385ab8a0d59\",\"error_uri\":\"https://login.microsoftonline.com/error?code=700016\"} Endpoint https://login.microsoftonline.com/f1821f56-8d20-4902-bf65-bc7502e367e7/oauth2/token?api-version=1.0")
		cerr = HandleErrorToCategorizedError(ctx, nil, err)
		assert.Equal(t, ADAL, cerr.Dependency)
		assert.Equal(t, ErrorSubCode("AADSTS700016"), cerr.SubCode)
		assert.Equal(t, apierror.ErrorCategory(""), cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)

		err = fmt.Errorf("azure.BearerAuthorizer#WithAuthorization: Failed to refresh the Token for request to https://management.azure.com/subscriptions/465d8b4d-4d84-4514-a5c3-6c59d97391e3/resourceGroups/MC_e2erg-e2emtrcs-21030919-zMqb3_e2eaks-dCZ_koreasouth/providers/Microsoft.Network/loadBalancers/kubernetes?api-version=2018-11-01: StatusCode=502 -- Original Error: adal: Failed to execute the refresh request. Error = 'Post \"https://login.microsoftonline.com/72f988bf-86f1-41af-91ab-2d7cd011db47/oauth2/token?api-version=1.0\": bad gateway")
		cerr = HandleErrorToCategorizedError(ctx, nil, err)
		assert.Equal(t, ADAL, cerr.Dependency)
		assert.Equal(t, ErrorSubCode("BadGateway"), cerr.SubCode)
		assert.Equal(t, apierror.InternalError, cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)
	})

	t.Run("without adalError", func(t *testing.T) {
		err := errors.New("VMSSAgentPoolReconciler retry failed: VMSS 'vmsstest' reported failure details={instance 1 has no instance views available, vmssInstanceErrorCode=NoVMSSInstanceView;}")
		cerr := HandleErrorToCategorizedError(ctx, nil, err)
		assert.Equal(t, ARM, cerr.Dependency)
		assert.Equal(t, ErrorSubCode("NoVMSSInstanceView"), cerr.SubCode)
		assert.Equal(t, apierror.InternalError, cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)
	})

	t.Run("with provided resp", func(t *testing.T) {
		resp := &http.Response{
			StatusCode: 409,
		}
		err := errors.New("Code=\"Conflict\" Message=\"The request failed due to conflict with a concurrent request. To resolve it, please refer to https://aka.ms/activitylog to get more details on the conflicting requests.")
		cerr := HandleErrorToCategorizedError(ctx, resp, err)
		assert.Equal(t, ARM, cerr.Dependency)
		assert.Equal(t, Conflict, cerr.SubCode)
		assert.Equal(t, apierror.ClientError, cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)

		resp = &http.Response{
			StatusCode: 409,
		}

		err = &azure.ServiceError{
			Code:    "OperationNotAllowed",
			Message: "Operation could not be completed as it results in exceeding approved Total Regional Cores quota. Additional details - Deployment Model: Resource Manager, Location: westeurope, Current Limit: 40, Current Usage: 38, Additional Required: 4, (Minimum) New Limit Required: 42.",
		}

		cerr = HandleErrorToCategorizedError(ctx, resp, err)
		assert.Equal(t, ARM, cerr.Dependency)
		assert.Equal(t, ErrorSubCode("OperationNotAllowed"), cerr.SubCode)
		assert.Equal(t, apierror.ClientError, cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)

		resp = &http.Response{
			StatusCode: 429,
		}

		err = errors.New("too many requests")
		cerr = HandleErrorToCategorizedError(ctx, resp, err)
		assert.Equal(t, ARM, cerr.Dependency)
		assert.Equal(t, ErrorSubCode("TooManyRequests"), cerr.SubCode)
		assert.Equal(t, apierror.ClientError, cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)
	})

	t.Run("InvalidParameter error is treated as InternalError", func(t *testing.T) {
		err := &autorest.DetailedError{
			Original: errors.New("InvalidParameter"),
			Method:   "GET",
		}

		cerr := HandleErrorToCategorizedError(ctx, nil, err)
		assert.Equal(t, ARM, cerr.Dependency)
		assert.Equal(t, InvalidParameter, cerr.SubCode)
		assert.Equal(t, apierror.InternalError, cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)
	})

	t.Run("InvalidParameter error with \"vm size not available\" is treated as ClientError", func(t *testing.T) {
		err := &autorest.DetailedError{
			Original: errors.New("Code=\"InvalidParameter\" Message=\"The requested VM size Standard_DS2_v2 is not available in the current region. The sizes available in the current region are: Standard_DC1s_v3,Standard_DC2s_v3,Standard_DC4s_v3,Standard_DC8s_v3,Standard_DC16s_v3,Standard_DC24s_v3,Standard_DC32s_v3,Standard_DC48s_v3,Standard_DC1ms_v3,Standard_DC2ms_v3,Standard_DC4ms_v3,Standard_DC8ms_v3,Standard_DC16ms_v3,Standard_DC24ms_v3,Standard_DC32ms_v3,Standard_DC1dms_v3,Standard_DC2dms_v3,Standard_DC4dms_v3,Standard_DC8dms_v3,Standard_DC16dms_v3,Standard_DC24dms_v3,Standard_DC32dms_v3.\r\nFind out more on the available VM sizes in each region at https://aka.ms/azure-regions."),
			Method:   "GET",
		}

		cerr := HandleErrorToCategorizedError(ctx, nil, err)
		assert.Equal(t, ARM, cerr.Dependency)
		assert.Equal(t, InvalidParameter, cerr.SubCode)
		assert.Equal(t, apierror.ClientError, cerr.Category)
		assert.Equal(t, log.AKSTeamRP, cerr.AKSTeam)
	})
}

func TestSetRetriableBasedOnCategorizedError(t *testing.T) {
	t.Run("client error", func(t *testing.T) {
		cerr := &CategorizedError{Category: apierror.ClientError}
		cerr = setRetriableBasedOnCategorizedError(cerr)
		assert.Equal(t, false, *cerr.Retriable)
	})

	t.Run("invalid parameter", func(t *testing.T) {
		cerr := &CategorizedError{
			Category: apierror.InternalError,
			SubCode:  ErrorSubCode("InvalidParameter"),
		}
		cerr = setRetriableBasedOnCategorizedError(cerr)
		assert.Equal(t, false, *cerr.Retriable)
	})

	t.Run("internal error", func(t *testing.T) {
		cerr := &CategorizedError{Category: apierror.InternalError}
		cerr = setRetriableBasedOnCategorizedError(cerr)
		assert.Equal(t, true, *cerr.Retriable)
	})
}
