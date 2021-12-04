package timing

//go:generate sh -c "mockgen -package=mock_timing goms.io/aks/rp/core/timing Timex >./mock_timing/mock_timex.go"

import (
	"time"
)

// Timex is the interface used to customize Sleep/Now and the time related mocking thing.
type Timex interface {
	Sleep(d time.Duration)
	Now() time.Time
}

type DefaultTimex struct {
}

func NewDefaultTimex() *DefaultTimex {
	return &DefaultTimex{}
}

func (dt *DefaultTimex) Sleep(d time.Duration) {
	time.Sleep(d)
}

func (dt *DefaultTimex) Now() time.Time {
	return time.Now()
}
