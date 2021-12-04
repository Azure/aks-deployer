package categorizederror

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	k8sapierrors "k8s.io/apimachinery/pkg/api/errors"
	//metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"github.com/Azure/aks-deployer/pkg/apierror"
	"github.com/Azure/aks-deployer/pkg/log"
)

func TestSetDependency(t *testing.T) {
	t.Run("setVMSS", func(t *testing.T) {
		err := &CategorizedError{
			Category:    apierror.InternalError,
			SubCode:     Conflict,
			Dependency:  ARM,
			OriginError: errors.New("conflict"),
		}

		cerr := err.SetDependency(VMSS)
		assert.Equal(t, VMSS, cerr.Dependency)
	})

	t.Run("setDependency if not ADAL", func(t *testing.T) {
		err := &CategorizedError{
			Category:    apierror.InternalError,
			SubCode:     Conflict,
			Dependency:  ARM,
			OriginError: errors.New("conflict"),
		}
		cerr := err.SetDependencyIfNotADAL(VMSS)
		assert.Equal(t, VMSS, cerr.Dependency)

		err.Dependency = ADAL
		cerr = err.SetDependencyIfNotADAL(VMSS)
		assert.Equal(t, ADAL, cerr.Dependency)
	})
}

func TestError(t *testing.T) {
	t.Run("no retriable", func(t *testing.T) {
		cerr := &CategorizedError{}
		assert.Equal(t, false, strings.Contains(cerr.Error(), "Retriable"))
	})
	t.Run("retriable", func(t *testing.T) {
		retriable := true
		cerr := &CategorizedError{
			Retriable: &retriable,
		}
		assert.Equal(t, true, strings.Contains(cerr.Error(), "Retriable: true"))
	})
}

func TestString(t *testing.T) {
	t.Run("retriable", func(t *testing.T) {
		retriable := true
		cerr := &CategorizedError{
			Retriable: &retriable,
		}
		assert.Equal(t, true, strings.Contains(fmt.Sprintf("cerr %s", cerr), "Retriable"))
	})
}

func TestSetRetriable(t *testing.T) {
	t.Run("set true", func(t *testing.T) {
		cerr := &CategorizedError{}
		cerr.SetRetriable(true)
		assert.Equal(t, true, *cerr.Retriable)
	})
	t.Run("set false", func(t *testing.T) {
		cerr := &CategorizedError{}
		cerr.SetRetriable(false)
		assert.Equal(t, false, *cerr.Retriable)
	})
}

func TestToCategorizedError(t *testing.T) {
	t.Run("convert to categorized error", func(t *testing.T) {
		err := &CategorizedError{
			Category:    apierror.InternalError,
			SubCode:     Conflict,
			Dependency:  ARM,
			OriginError: errors.New("conflict"),
		}

		cerr := ToCategorizedError(err)
		assert.Equal(t, ARM, cerr.Dependency)
		assert.Equal(t, apierror.InternalError, cerr.Category)
		assert.Equal(t, Conflict, cerr.SubCode)
	})

	t.Run("convert to categorized error", func(t *testing.T) {
		err := errors.New("conflict")

		cerr := ToCategorizedError(err)
		assert.Equal(t, ResourceType(""), cerr.Dependency)
		assert.Equal(t, apierror.ErrorCategory(""), cerr.Category)
		assert.Equal(t, ErrorSubCode(""), cerr.SubCode)
		assert.Equal(t, err, cerr.OriginError)
	})
}

func TestGRPCStatus(t *testing.T) {
	t.Run("GRPCStatus", func(t *testing.T) {
		assert.Equal(t, status.Code(nil), codes.OK)

		err := &CategorizedError{
			Category: apierror.InternalError,
		}
		assert.Equal(t, status.Code(err), codes.OK)

		err = &CategorizedError{
			Category:    apierror.InternalError,
			OriginError: status.Error(codes.NotFound, "not found"),
		}

		assert.Equal(t, status.Code(err), codes.NotFound)

		err = &CategorizedError{
			Category:    apierror.InternalError,
			OriginError: errors.New("new error"),
		}

		assert.Equal(t, status.Code(err), codes.Unknown)
	})
}

func TestStatus(t *testing.T) {
	t.Run("k8s status", func(t *testing.T) {
		assert.Equal(t, k8sapierrors.IsConflict(nil), false)
		err := &CategorizedError{
			Category: apierror.InternalError,
		}
		assert.Equal(t, k8sapierrors.IsNotFound(err), false)

		err = &CategorizedError{
			Category:    apierror.InternalError,
			OriginError: k8sapierrors.NewNotFound(schema.GroupResource{}, "name"),
		}

		assert.Equal(t, k8sapierrors.IsNotFound(err), true)
		assert.Equal(t, k8sapierrors.IsConflict(err), false)

		err = &CategorizedError{
			Category:    apierror.InternalError,
			OriginError: k8sapierrors.NewServiceUnavailable("reason"),
		}

		assert.Equal(t, k8sapierrors.IsNotFound(err), false)
		assert.Equal(t, k8sapierrors.IsConflict(err), false)
		assert.Equal(t, k8sapierrors.IsServiceUnavailable(err), true)
	})
}

func TestToAPIErrorResponse(t *testing.T) {
	t.Run("ToAPIErrorResponse", func(t *testing.T) {
		cerr := &CategorizedError{
			Category:    apierror.InternalError,
			SubCode:     "FakeSubCode",
			Dependency:  HCP,
			OriginError: errors.New("fake original error"),
			AKSTeam:     log.AKSTeamNodeProvisioning,
		}

		apiErrorResponse := cerr.ToAPIErrorResponse("FakeErrorCode", "fake error message")

		err := apiErrorResponse.Body
		assert.Equal(t, err.Code, apierror.ErrorCode("FakeErrorCode"))
		assert.Equal(t, err.Message, "fake error message")
		assert.Equal(t, err.Target, "")
		assert.Equal(t, err.Category, apierror.InternalError)
		assert.Equal(t, err.Subcode, "FakeSubCode")
		assert.Equal(t, err.InnerMessage, "Category: InternalError; SubCode: FakeSubCode; Dependency: HostedControlPlaneDataStore; OrginalError: fake original error; AKSTeam: NodeProvisioning")
	})
}
