package retry

import (
	"time"
)

// Settings is settings for retry
type Settings struct {
	Interval      time.Duration
	Timeout       time.Duration
	RetryMaxCount int
}
