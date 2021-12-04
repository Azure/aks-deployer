package categorizederror

import (
	"context"
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/aks-deployer/pkg/apierror"
	"github.com/Azure/aks-deployer/pkg/log"
)

// CategorizedError is used for error categorization.
type CategorizedError struct {
	Category    apierror.ErrorCategory
	SubCode     ErrorSubCode
	Dependency  ResourceType
	OriginError error

	// Retriable used for client to try or not. If it is nil or false, we don't retry, and if it
	// is true, client needs to retry. Also, when it is nil, which means we don't know yet if it is retriable.
	Retriable *bool

	// AKSTeam is the sub-team that owns the code that generated the error.
	// This info is helpful for troubleshooting since it tells who is most knowledgeable
	// and thus can help with issue investigation.
	AKSTeam log.AKSTeam
}

func NewCategorizedError(
	ctx context.Context,
	category apierror.ErrorCategory,
	subCode ErrorSubCode,
	dependency ResourceType,
	originError error) *CategorizedError {
	return &CategorizedError{
		Category:    category,
		SubCode:     subCode,
		Dependency:  dependency,
		OriginError: originError,
		AKSTeam:     log.GetTeamFromContext(ctx),
	}
}

var _ error = &CategorizedError{}

func (c *CategorizedError) Error() string {
	return c.String()
}

func (c *CategorizedError) String() string {
	if c.Retriable == nil {
		return fmt.Sprintf("Category: %s; SubCode: %s; Dependency: %s; OrginalError: %s; AKSTeam: %s",
			string(c.Category),
			string(c.SubCode),
			string(c.Dependency),
			c.OriginError,
			string(c.AKSTeam))
	}
	return fmt.Sprintf("Category: %s; SubCode: %s; Dependency: %s; OrginalError: %s; AKSTeam: %s, Retriable: %t",
		string(c.Category),
		string(c.SubCode),
		string(c.Dependency),
		c.OriginError,
		string(c.AKSTeam),
		*c.Retriable)
}

func (c *CategorizedError) Unwrap() error {
	return c.OriginError
}

func (c *CategorizedError) SetDependency(dep ResourceType) *CategorizedError {
	c.Dependency = dep
	return c
}

func (c *CategorizedError) SetDependencyIfNotADAL(dep ResourceType) *CategorizedError {
	if c.Dependency == ADAL {
		return c
	}
	return c.SetDependency(dep)
}

func (c *CategorizedError) SetRetriable(retriable bool) *CategorizedError {
	c.Retriable = &retriable
	return c
}

// implemented GRPCStutus interface to make sure error returned from HCP can be
// cast to codes.Code from status.Code()
func (c *CategorizedError) GRPCStatus() *status.Status {
	if c == nil || c.OriginError == nil {
		return nil
	}
	if se, ok := c.OriginError.(interface {
		GRPCStatus() *status.Status
	}); ok {
		return se.GRPCStatus()
	}
	return status.New(codes.Unknown, c.OriginError.Error())
}

// implement k8s Status interface to make sure error returned from k8s client can be
// cast to metav1.Status
func (c *CategorizedError) Status() metav1.Status {
	if c == nil || c.OriginError == nil {
		return metav1.Status{}
	}
	if se, ok := c.OriginError.(apierrors.APIStatus); ok {
		return se.Status()
	}
	return metav1.Status{
		Reason:  metav1.StatusReasonUnknown,
		Message: c.OriginError.Error(),
	}
}

func (c *CategorizedError) ToAPIErrorResponse(errorCode apierror.ErrorCode, message string) *apierror.ErrorResponse {
	return apierror.NewWithInnerMessage(
		c.Category,
		errorCode,
		string(c.SubCode),
		message,
		c.Error())
}

func ToCategorizedError(err error) *CategorizedError {
	var cerr *CategorizedError
	if errors.As(err, &cerr) {
		return cerr
	}
	return &CategorizedError{
		OriginError: err,
	}
}
