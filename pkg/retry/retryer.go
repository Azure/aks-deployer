package retry

import (
	"context"
	"fmt"
	"time"

	"go.opencensus.io/trace"

	"github.com/Azure/aks-deployer/pkg/log"
)

type retryerImpl struct {
	name            string
	settings        Settings
	singleIteration SingleIterationInterface
}

func NewRetry(name string, settings Settings, singleIteration SingleIterationInterface) RetryerInterface {
	return &retryerImpl{
		name:            name,
		settings:        settings,
		singleIteration: singleIteration,
	}
}

func prepareSetting(ctx context.Context, name string, settings Settings) Settings {
	if settings.Interval <= 0 {
		settings.Interval = 5 * time.Second
		log.GetLogger(ctx).Warningf(ctx, "Retry '%s' doesn't have valid interval supplied. Using the default value: %v", name, settings.Interval)
	}

	log.GetLogger(ctx).Infof(ctx, "The interval value: %v for retry %s", settings.Interval, name)

	if settings.Timeout <= 0 {
		settings.Timeout = 10 * time.Minute
		log.GetLogger(ctx).Warningf(ctx, "Retry '%s' doesn't have valid timeout supplied. Using the default value: %v.", name, settings.Timeout)
	}
	log.GetLogger(ctx).Infof(ctx, "The timeout value: %v for retry %s", settings.Timeout, name)

	if settings.RetryMaxCount <= 0 {
		settings.RetryMaxCount = 10
		log.GetLogger(ctx).Warningf(ctx, "Retry '%s' doesn't have valid RetryMaxCount supplied. Using the default value: %v.", name, settings.RetryMaxCount)
	}
	log.GetLogger(ctx).Infof(ctx, "The RetryMaxCount %v for retry %s", settings.RetryMaxCount, name)

	return settings
}

func (r *retryerImpl) Run(ctx context.Context) (interface{}, error) {
	if r.name == "" {
		return nil, fmt.Errorf("The retry name cannot be empty")
	}

	if r.singleIteration == nil {
		return nil, fmt.Errorf("The singleIteration cannot be nil")
	}

	r.settings = prepareSetting(ctx, r.name, r.settings)

	type internalResult struct {
		returnValue interface{}
		err         error
	}
	retryAttempt := 0
	ctx, span := log.StartSpan(ctx, fmt.Sprintf("%s.Run", r.name), log.AKSTeamSameAsParent)
	defer span.End()
	result, err := DoFixedRetryWithMaxCount(
		func() Result {
			// check if context is canceled
			err := ctx.Err()
			if err != nil {
				log.GetLogger(ctx).Infof(ctx, "%s stop retry due to %s", r.name, err)
				return Result{
					Status: Failed,
					Body: &internalResult{
						returnValue: nil,
						err:         err,
					},
				}
			}

			retryAttempt++
			log.GetLogger(ctx).Infof(ctx, "%s retry attempt %d", r.name, retryAttempt)
			ctx, span := log.StartSpan(ctx, fmt.Sprintf("%s.RunOnce.%d", r.name, retryAttempt), log.AKSTeamSameAsParent)
			defer span.End()
			status, returnValue, err := r.singleIteration.RunOnce(ctx)
			log.WithTags(span,
				map[string]string{
					"ReconcilerName": r.name,
					"RunOnceStatus":  status.String(),
				})
			if err != nil {
				span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
			}
			return Result{
				Status: status,
				Body: &internalResult{
					returnValue: returnValue,
					err:         err,
				},
			}
		},
		r.settings.Interval,
		r.settings.Timeout,
		r.settings.RetryMaxCount)

	log.WithTags(span,
		map[string]string{
			"ReconcilerName": r.name,
			"RunStatus":      result.Status.String(),
		})

	if err != nil {
		span.SetStatus(trace.Status{Code: trace.StatusCodeInternal, Message: err.Error()})
		log.GetLogger(ctx).Errorf(ctx, "%s retry failed: '%v'", r.name, err)
		return nil, err
	}

	ir := result.Body.(*internalResult)
	if result.Status == Failed {
		log.GetLogger(ctx).Errorf(ctx, "%s retry failed", r.name)
		return ir.returnValue, fmt.Errorf("%s retry failed: %w", r.name, ir.err)
	}

	if result.Status == Timedout {
		return ir.returnValue, fmt.Errorf("%s retry timed out: %w", r.name, ir.err)
	}

	log.GetLogger(ctx).Infof(ctx, "%s retry succeeded.", r.name)

	return ir.returnValue, nil
}
