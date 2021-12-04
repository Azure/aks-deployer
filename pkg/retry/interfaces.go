package retry

//go:generate sh -c "mockgen goms.io/aks/rp/core/retry SingleIterationInterface,RetryerInterface>./mock_$GOPACKAGE/interfaces.go"

import (
	"context"
)

type SingleIterationInterface interface {
	RunOnce(ctx context.Context) (Status, interface{}, error)
}

type Func func(context.Context) (interface{}, *Error)

func (f Func) RunOnce(ctx context.Context) (Status, interface{}, error) {
	val, rerr := f(ctx)
	if rerr == nil {
		return Success, val, nil
	}

	return rerr.Status(), nil, rerr.Error
}

// RetryerInterface is the interface for retryer
type RetryerInterface interface {
	Run(ctx context.Context) (interface{}, error)
}
