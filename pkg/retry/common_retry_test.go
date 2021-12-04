package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTestActuatorDoNilWaitFunc(t *testing.T) {
	err := (&Actuator{}).Do(context.Background(), func() error {
		return nil
	})

	assert.EqualError(t, err, "wait function is required")
}

func TestActuatorDoNoError(t *testing.T) {
	err := (&Actuator{
		Wait: Linear(0),
	}).Do(context.Background(), func() error {
		return nil
	})

	assert.NoError(t, err)
}

func TestActuatorDoRetryFixed(t *testing.T) {
	probe := &testProbe{}

	calls := 0
	err := (&Actuator{
		Wait:        Fixed(1),
		BeforeSleep: probe.Hook,
	}).Do(context.Background(), func() error {
		calls++
		if calls < 3 {
			return errors.New("transient error")
		}
		return nil
	})

	assert.NoError(t, err, "no error is returned")
	assert.Equal(t, 3, calls, "function was called twice")
	assert.Equal(t, []time.Duration{1, 1}, probe.Intervals, "function was called at expected intervals")
}

func TestActuatorDoRetryLinear(t *testing.T) {
	probe := &testProbe{}

	calls := 0
	err := (&Actuator{
		Wait:        Linear(1),
		BeforeSleep: probe.Hook,
	}).Do(context.Background(), func() error {
		calls++
		if calls < 3 {
			return errors.New("transient error")
		}
		return nil
	})

	assert.NoError(t, err, "no error is returned")
	assert.Equal(t, 3, calls, "function was called twice")
	assert.Equal(t, []time.Duration{1, 2}, probe.Intervals, "function was called at expected intervals")
}

func TestActuatorMaxAttemptsExceeded(t *testing.T) {
	testError := errors.New("persistent error")
	attempts := 2

	err := (&Actuator{
		MaxAttempts: &attempts,
		Wait:        Fixed(0),
	}).Do(context.Background(), func() error {
		return testError
	})

	require.Error(t, err)
	assert.IsType(t, &TimeoutError{}, err)
	assert.Equal(t, testError.Error(), err.Error())
}

func TestActuatorDoTimeout(t *testing.T) {
	timeout := time.Duration(0)

	err := (&Actuator{
		Wait:    Fixed(0),
		Timeout: &timeout,
	}).Do(context.Background(), func() error {
		return errors.New("persistent error")
	})

	assert.EqualError(t, err, "persistent error")
	assert.IsType(t, &TimeoutError{}, err)
}

func TestActuatorDoTerminalError(t *testing.T) {
	testError := &testTerminalErrorType{}

	err := (&Actuator{
		Wait: Fixed(0),
	}).Do(context.Background(), func() error {
		return testError
	})

	assert.Equal(t, testError, err)
}

func TestActuatorDoContextDeadline(t *testing.T) {
	ctx, done := context.WithTimeout(context.Background(), 0)
	defer done()

	err := (&Actuator{
		Wait: Fixed(0),
	}).Do(ctx, func() error {
		return errors.New("persistent error")
	})

	assert.EqualError(t, err, "persistent error")
	assert.IsType(t, &TimeoutError{}, err)
}

func TestActuatorDoContextCanceledBeforeFirstAttempt(t *testing.T) {
	ctx, done := context.WithCancel(context.Background())
	done()

	err := (&Actuator{
		Wait: Fixed(0),
	}).Do(ctx, func() error {
		return nil
	})

	assert.Equal(t, context.Canceled, err)
}

func TestActuatorDoContextCanceledAfterFirstAttempt(t *testing.T) {
	ctx, done := context.WithCancel(context.Background())
	done()

	err := (&Actuator{
		Wait: Fixed(0),
	}).Do(ctx, func() error {
		return errors.New("persistent error")
	})

	assert.Equal(t, context.Canceled, err)
}

type testProbe struct {
	Intervals []time.Duration
}

func (t *testProbe) Hook(ctx context.Context, attempt int, interval time.Duration) context.Context {
	t.Intervals = append(t.Intervals, interval)
	return ctx
}

type testTerminalErrorType struct{}

func (t *testTerminalErrorType) Error() string    { return "it didn't work" }
func (t *testTerminalErrorType) IsTerminal() bool { return true }
