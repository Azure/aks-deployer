// This is actually moved from aks/common/retry
// and we should switch to use the retry in the aks/rp/core/retry
// filed one work item to track this:
// https://dev.azure.com/msazure/CloudNativeCompute/_workitems/edit/7483207
package retry

import (
	"context"
	"errors"
	"time"
)

// WaitFunc is responsible for determining how long to wait before retrying
// an operation.
type WaitFunc func(i int) time.Duration

// Fixed returns a WaitFunc that always waits for the given duration.
func Fixed(interval time.Duration) WaitFunc {
	return func(i int) time.Duration {
		return interval
	}
}

// Linear returns a WaitFunc that always waits for one given duration per attempt.
func Linear(interval time.Duration) WaitFunc {
	return func(i int) time.Duration {
		return time.Duration(i) * interval
	}
}

// Actuator is capable of invoking an operation with retry logic.
type Actuator struct {
	// MaxAttempts specifies an optional limit of failed attempts.
	MaxAttempts *int

	// Timeout is an optional time limit for the entire operation including
	// every attempt.
	Timeout *time.Duration

	// Wait is used to determine the amount of time to sleep between attempts.
	// Required.
	Wait WaitFunc

	// BeforeSleep is an optional hook to update the context before the next
	// attempt. Can be used to add the attempt number to a logger.
	BeforeSleep func(ctx context.Context, attempt int, interval time.Duration) context.Context
}

// Do retries the given function until err == nil, MaxAttempts is reached, context
// timeout, or context cancelation.
//
// Retry loop is broken if the returned error satisfies the Error interface and
// IsTerminal returns true. Otherwise, returning an error will result in another
// attempt. The caller is responsible for logging or otherwise handling errors.
func (a *Actuator) Do(ctx context.Context, fn func() error) error {
	if a.Timeout != nil {
		var done context.CancelFunc
		ctx, done = context.WithTimeout(ctx, *a.Timeout)
		defer done()
	}

	if a.Wait == nil {
		return errors.New("wait function is required")
	}

	var err error
	for i := 1; a.MaxAttempts == nil || i < *a.MaxAttempts; i++ {
		if ctx.Err() == context.Canceled {
			return ctx.Err()
		}

		err = fn()
		if err == nil {
			return nil
		}
		if e, ok := err.(CommonError); ok && e.IsTerminal() {
			return err
		}

		interval := a.Wait(i)
		if a.BeforeSleep != nil {
			ctx = a.BeforeSleep(ctx, i, interval)
		}

		timer := time.NewTimer(interval)
		select {
		case <-timer.C:
		case <-ctx.Done():
			timer.Stop()

			if ctx.Err() == context.DeadlineExceeded {
				return &TimeoutError{Original: err}
			}
			return ctx.Err()
		}
	}

	return &TimeoutError{Original: err}
}

// Error can be implemented by error types returned by lambdas provided to
// Do functions in order to influence the retry logic.
type CommonError interface {
	// IsTerminal indicates that an operation should not be retried
	// because it cannot succeed.
	IsTerminal() bool
}

// TimeoutError is returned when a retry operation has timed out.
type TimeoutError struct {
	Original error
}

var _ error = &TimeoutError{}

func (t *TimeoutError) Error() string { return t.Original.Error() }
