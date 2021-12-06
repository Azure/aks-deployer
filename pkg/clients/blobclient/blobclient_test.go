package blobclient

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Azure/go-autorest/autorest"
	"github.com/stretchr/testify/require"
	cgerror "github.com/Azure/aks-deployer/pkg/categorizederror"
)

func TestBlobClient_GetBlobData(t *testing.T) {
	t.Run("returns correct data", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			require.Equal(t, req.Method, http.MethodGet)

			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte("aaa"))
		}))
		defer server.Close()

		client := &BlobClient{}
		_, err := client.GetBlobData(context.Background(), server.URL)
		require.NoError(t, err)
	})
}

func TestBlobClient_ContainerExists(t *testing.T) {
	t.Run("return error for unexpected status code", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodGet)
				resp := &http.Response{StatusCode: http.StatusBadGateway}
				return resp, nil
			})),
		}
		exists, err := client.ContainerExists(ctx, "test-sa", "test-container")
		require.Error(t, err)
		require.False(t, exists)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.UnexpectedStatusCodeCheckContainerExist)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})

	t.Run("return error for nil response", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodGet)
				return nil, nil
			})),
		}
		exists, err := client.ContainerExists(ctx, "test-sa", "test-container")
		require.Error(t, err)
		require.False(t, exists)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.FailedGetResponse)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})
}

func TestBlobClient_DeleteContainer(t *testing.T) {
	t.Run("return error for unexpected status", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodDelete)
				resp := &http.Response{StatusCode: http.StatusBadGateway}
				return resp, nil
			})),
		}
		err := client.DeleteContainer(ctx, "test-sa", "test-container")
		require.Error(t, err)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.UnexpectedStatusCodeDeleteContainer)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})

	t.Run("return error for nil response", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodDelete)
				return nil, nil
			})),
		}
		err := client.DeleteContainer(ctx, "test-sa", "test-container")
		require.Error(t, err)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.FailedGetResponse)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})
}

func TestBlobClient_BlobExists(t *testing.T) {
	t.Run("return not exists for not found status code", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodHead)
				resp := &http.Response{StatusCode: http.StatusNotFound}
				return resp, nil
			})),
		}
		exists, err := client.BlobExists(ctx, "test-sa", "test-container", "test-blob")
		require.NoError(t, err)
		require.False(t, exists)
	})

	t.Run("return exist for OK status code", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodHead)
				resp := &http.Response{StatusCode: http.StatusOK}
				return resp, nil
			})),
		}
		exists, err := client.BlobExists(ctx, "test-sa", "test-container", "test-blob")
		require.NoError(t, err)
		require.True(t, exists)
	})

	t.Run("return error for unknown status code", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodHead)
				resp := &http.Response{StatusCode: http.StatusBadGateway}
				return resp, nil
			})),
		}
		exists, err := client.BlobExists(ctx, "test-sa", "test-container", "test-blob")
		require.Error(t, err)
		require.False(t, exists)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.UnexpectedStatusCodeCheckBlobExist)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})

	t.Run("return error for nil response", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodHead)
				return nil, nil
			})),
		}
		exists, err := client.BlobExists(ctx, "test-sa", "test-container", "test-blob")
		require.Error(t, err)
		require.False(t, exists)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.FailedGetResponse)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})
}

func TestBlobClient_CreateContainer(t *testing.T) {
	ctx := context.Background()
	client := &BlobClient{
		restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
			require.NotContains(t, req.Header, headerNameBlobPublicAccess, "should not contain acl header")

			resp := &http.Response{StatusCode: http.StatusCreated}
			return resp, nil
		})),
	}
	err := client.CreateContainer(ctx, "test-sa", "test-container")
	require.NoError(t, err)
}

func TestBlobClient_CreateContainerWithACL(t *testing.T) {
	t.Run("return error for unexpected status code", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodPut)
				resp := &http.Response{StatusCode: http.StatusBadGateway}
				return resp, nil
			})),
		}
		err := client.CreateContainerWithACL(ctx, "test-sa", "test-container", BlobPublicAccessLevelContainer)
		require.Error(t, err)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.UnexpectedStatusCodeCreateContainer)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})

	t.Run("return error for nil response", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodPut)
				return nil, nil
			})),
		}
		err := client.CreateContainerWithACL(ctx, "test-sa", "test-container", BlobPublicAccessLevelContainer)
		require.Error(t, err)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.FailedGetResponse)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})
}

func TestBlobClient_CreateContainerWithACL_SetACLHeader(t *testing.T) {
	cases := []struct {
		acl               BlobPublicAccessLevel
		expectedSetHeader bool
	}{
		{
			acl:               BlobPublicAccessLevelNotSet,
			expectedSetHeader: false,
		},
		{
			acl:               BlobPublicAccessLevelContainer,
			expectedSetHeader: true,
		},
		{
			acl:               BlobPublicAccessLevelBlob,
			expectedSetHeader: true,
		},
	}

	for idx := range cases {
		t.Run(fmt.Sprintf("case #%d", idx), func(t *testing.T) {
			c := cases[idx]

			ctx := context.Background()
			client := &BlobClient{
				restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
					if c.expectedSetHeader {
						require.NotEmpty(t, req.Header.Get(headerNameBlobPublicAccess))
						require.Equal(t, req.Header.Get(headerNameBlobPublicAccess), string(c.acl))
					} else {
						require.Empty(t, req.Header.Get(headerNameBlobPublicAccess))
					}

					resp := &http.Response{StatusCode: http.StatusCreated}
					return resp, nil
				})),
			}
			err := client.CreateContainerWithACL(ctx, "test-sa", "test-container", c.acl)
			require.NoError(t, err)
		})
	}
}

func TestBlobClient_AcquireLease(t *testing.T) {
	t.Run("return error for unexpected status code", func(t *testing.T) {
		ctx := context.Background()
		leaseDuration := "30"
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodPut)
				require.Equal(t, req.Header.Get("x-ms-lease-action"), "acquire")
				require.Equal(t, req.Header.Get("x-ms-lease-duration"), leaseDuration)
				resp := &http.Response{StatusCode: http.StatusBadGateway}
				return resp, nil
			})),
		}
		_, err := client.AcquireLease(ctx, "test-sa", "test-container", "test-blob", 30)
		require.Error(t, err)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.UnexpectedStatusCodeAcquireLease)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})

	t.Run("return error for nil response", func(t *testing.T) {
		ctx := context.Background()
		leaseDuration := "30"
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodPut)
				require.Equal(t, req.Header.Get("x-ms-lease-action"), "acquire")
				require.Equal(t, req.Header.Get("x-ms-lease-duration"), leaseDuration)
				return nil, nil
			})),
		}
		_, err := client.AcquireLease(ctx, "test-sa", "test-container", "test-blob", 30)
		require.Error(t, err)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.FailedGetResponse)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})
}

func TestBlobClient_ReleaseLease(t *testing.T) {
	t.Run("set expected header", func(t *testing.T) {
		ctx := context.Background()
		testLeaseID := "test-lease-id"
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodPut)
				require.Equal(t, req.Header.Get("x-ms-lease-action"), "release")
				require.Equal(t, req.Header.Get(HeaderLeaseID), testLeaseID)
				resp := &http.Response{StatusCode: http.StatusOK}
				return resp, nil
			})),
		}
		err := client.ReleaseLease(ctx, "test-sa", "test-container", "test-blob", testLeaseID)
		require.NoError(t, err)
	})
}

func TestBlobClient_BreakLease(t *testing.T) {
	t.Run("set expected header", func(t *testing.T) {
		ctx := context.Background()
		testLeaseID := "test-lease-id"
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodPut)
				require.Equal(t, req.Header.Get("x-ms-lease-action"), "break")
				require.Equal(t, req.Header.Get(HeaderLeaseID), testLeaseID)
				resp := &http.Response{StatusCode: http.StatusAccepted}
				return resp, nil
			})),
		}
		err := client.BreakLease(ctx, "test-sa", "test-container", "test-blob", testLeaseID)
		require.NoError(t, err)
	})
}

func TestBlobClient_PutBlobDataV1(t *testing.T) {
	t.Run("put failed", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodPut)
				resp := &http.Response{StatusCode: http.StatusServiceUnavailable}
				return resp, nil
			})),
		}
		err := client.PutBlobDataV1(ctx,
			"test-sa", "test-container", "test-blob", []byte("test-data"),
			nil,
		)
		require.Error(t, err)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.UnexpectedStatusCodePutBlob)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})

	t.Run("put without extra headers", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodPut)
				resp := &http.Response{StatusCode: http.StatusCreated}
				return resp, nil
			})),
		}
		err := client.PutBlobDataV1(ctx,
			"test-sa", "test-container", "test-blob", []byte("test-data"),
			nil,
		)
		require.NoError(t, err)
	})

	t.Run("put with extra headers", func(t *testing.T) {
		ctx := context.Background()
		testLeaseID := "test-lease-id"
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodPut)
				require.Equal(t, req.Header.Get(HeaderLeaseID), testLeaseID)
				resp := &http.Response{StatusCode: http.StatusCreated}
				return resp, nil
			})),
		}
		err := client.PutBlobDataV1(ctx,
			"test-sa", "test-container", "test-blob", []byte("test-data"),
			map[string]interface{}{
				HeaderLeaseID: testLeaseID,
			},
		)
		require.NoError(t, err)
	})

	t.Run("put without response", func(t *testing.T) {
		ctx := context.Background()
		client := &BlobClient{
			restClient: newMockAutoRestClient(autorest.SenderFunc(func(req *http.Request) (*http.Response, error) {
				require.Equal(t, req.Method, http.MethodPut)
				return nil, nil
			})),
		}
		err := client.PutBlobDataV1(ctx,
			"test-sa", "test-container", "test-blob", []byte("test-data"),
			nil,
		)
		require.Error(t, err)
		cerr, ok := err.(*cgerror.CategorizedError)
		require.True(t, ok)
		require.Equal(t, cerr.SubCode, cgerror.FailedGetResponse)
		require.Equal(t, cerr.Dependency, cgerror.BlobStorage)
	})
}
