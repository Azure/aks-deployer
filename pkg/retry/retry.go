//------------------------------------------------------------
// Copyright (c) Microsoft Corporation.  All rights reserved.
//------------------------------------------------------------

// Package retry provides utility functions to execute retriable action with alloted timeout duration.
package retry

import (
	"context"
	"errors"
	"math"
	"math/rand"
	"time"

	"github.com/Azure/aks-deployer/pkg/timing"
)

const (
	timexKey = "retrytimex"
)

func WithTimex(ctx context.Context, timex timing.Timex) context.Context {
	return context.WithValue(ctx, timexKey, timex)
}

func GetTimex(ctx context.Context) timing.Timex {
	retVal, ok := ctx.Value(timexKey).(timing.Timex)
	if !ok {
		return nil
	}
	return retVal
}

// A Result specifies the result of retryable action.
type Result struct {
	Status
	Body interface{}
}

// Status is an arbitrary status code representing the status of a retry
type Status int

// Status definitions
const (
	Success Status = iota
	Failed
	NeedRetry
	Timedout
)

func (s Status) String() string {
	if s == Success {
		return "Success"
	}
	if s == Failed {
		return "Failed"
	}
	if s == NeedRetry {
		return "NeedRetry"
	}
	if s == Timedout {
		return "Timedout"
	}
	return ""
}

// returns (count * interval)
func getLinearDelay(count int, interval time.Duration) time.Duration {
	return time.Duration(count) * interval
}

func getFixedDelay(_ int, interval time.Duration) time.Duration {
	return interval
}

func getJitteryDelay(_ int, interval time.Duration) time.Duration {
	/* #nosec */
	return time.Duration(rand.Int63n(interval.Nanoseconds()))
}

// DoLinearRetryWithMaxCount retries the action for either the maxRetries times or the timeout duration with (interval * attempt) [attempt starts at 0] in between retries
func DoLinearRetryWithMaxCount(action func() Result, interval time.Duration, timeout time.Duration, maxRetries int) (Result, error) {
	return doLinearRetryWithMaxCount(action, interval, timeout, maxRetries, &timing.DefaultTimex{})
}

// DoLinearRetryWithMaxCountWithTimex retries the action for either the maxRetries times or the timeout duration with (interval * attempt) [attempt starts at 0] in between retries
func DoLinearRetryWithMaxCountWithTimex(action func() Result, interval time.Duration, timeout time.Duration, maxRetries int, timex timing.Timex) (Result, error) {
	if timex == nil {
		timex = &timing.DefaultTimex{}
	}
	return doLinearRetryWithMaxCount(action, interval, timeout, maxRetries, timex)
}

// DoFixedRetryWithMaxCount retries the action for either the maxRetries times or the timeout duration with interval duration in between retries
func DoFixedRetryWithMaxCount(action func() Result, interval time.Duration, timeout time.Duration, maxRetries int) (Result, error) {
	return doFixedRetryWithMaxCount(action, interval, timeout, maxRetries, &timing.DefaultTimex{})
}

// DoLinearRetry retries the action for timeout duration with interval duration in between retries
func DoLinearRetry(action func() Result, interval time.Duration, timeout time.Duration) (Result, error) {
	return doLinearRetry(action, interval, timeout, &timing.DefaultTimex{})
}

// DoJitteryRetry retry with random delay at least max times. This avoids multiple clients retrying the same resource at same time
// Jitter is defined here: https://github.com/App-vNext/Polly/wiki/Retry-with-jitter &https://github.com/avast/retry-go
func DoJitteryRetry(action func() Result, interval time.Duration, timeout time.Duration, maxRetries int) (Result, error) {
	return doJitteryRetry(action, interval, timeout, maxRetries, &timing.DefaultTimex{})
}

// DoRetryWithMaxCount retries the action for timeout duration with interval duration in between retries, with custom delayfunc, which could be exponential,fixed etc
// As this function heavily depends on the global function of Time package, which is not easy to mock when testing.
// As a result, replace it with a new method which has an extra parameter, Timex interface.
func DoRetryWithMaxCount(action func() Result, interval time.Duration, timeout time.Duration, maxRetries int, delayfunc func(int, time.Duration) time.Duration) (Result, error) {
	return doRetryWithMaxCount(action, interval, timeout, maxRetries, delayfunc, &timing.DefaultTimex{})
}

// doLinearRetryWithMaxCount retries the action for either the maxRetries times or the timeout duration with (interval * attempt) [attempt starts at 0] in between retries
func doLinearRetryWithMaxCount(action func() Result, interval time.Duration, timeout time.Duration, maxRetries int, timex timing.Timex) (Result, error) {
	return doRetryWithMaxCount(action, interval, timeout, maxRetries, getLinearDelay, timex)
}

// doFixedRetryWithMaxCount retries the action for either the maxRetries times or the timeout duration with interval duration in between retries
func doFixedRetryWithMaxCount(action func() Result, interval time.Duration, timeout time.Duration, maxRetries int, timex timing.Timex) (Result, error) {
	return doRetryWithMaxCount(action, interval, timeout, maxRetries, getFixedDelay, timex)
}

// doLinearRetry retries the action for timeout duration with interval duration in between retries
func doLinearRetry(action func() Result, interval time.Duration, timeout time.Duration, timex timing.Timex) (Result, error) {
	return doLinearRetryWithMaxCount(action, interval, timeout, math.MaxInt32, timex)
}

// doJitteryRetry retry with random delay at least max times. This avoids multiple clients retrying the same resource at same time
// Jitter is defined here: https://github.com/App-vNext/Polly/wiki/Retry-with-jitter &https://github.com/avast/retry-go
func doJitteryRetry(action func() Result, interval time.Duration, timeout time.Duration, maxRetries int, timex timing.Timex) (Result, error) {
	return doRetryWithMaxCount(action, interval, timeout, maxRetries, getJitteryDelay, timex)
}

// doRetryWithMaxCount retries the action for timeout duration with interval duration in between retries, with custom delayfunc, which could be exponential,fixed etc
func doRetryWithMaxCount(action func() Result, interval time.Duration, timeout time.Duration, maxRetries int, delayfunc func(int, time.Duration) time.Duration, timex timing.Timex) (Result, error) {
	if interval <= 0 {
		return Result{}, errors.New("interval cannot be less than or equal to zero")
	}
	if timeout <= 0 {
		return Result{}, errors.New("timeout cannot be less than or equal to zero")
	}
	if maxRetries <= 0 {
		return Result{}, errors.New("maxRetries cannot be less than or equal to zero")
	}
	if action == nil {
		return Result{}, errors.New("action cannot be nil")
	}

	deadline := timex.Now().Add(timeout)

	var ret Result
	for i := 0; i < maxRetries; i++ {
		ret = action()
		if ret.Status != NeedRetry {
			return ret, nil
		}

		if timex.Now().After(deadline) {
			return Result{Status: Timedout, Body: ret.Body}, nil
		}

		timex.Sleep(delayfunc(i, interval))
	}

	return Result{Status: Timedout, Body: ret.Body}, nil
}
