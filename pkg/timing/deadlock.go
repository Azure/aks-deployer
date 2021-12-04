package timing

import (
	"context"
	"os"
	"runtime"
	"time"

	"github.com/Azure/aks-deployer/pkg/log"
)

var noop = func() {}

// IdentifyDeadlocks attempts to find deadlock conditions by waiting a given duration before
// logging the stack of all goroutines. Make sure to defer the returned function to avoid leaks.

// Can generate up to 2mb of console output, so it is only enabled in e2e.
func IdentifyDeadlocks(ctx context.Context, after time.Duration) func() {
	return identifyDeadlocks(after, logDeadlocks(ctx))
}

func logDeadlocks(ctx context.Context) func(string) {
	return func(stack string) {
		log.GetLogger(ctx).Warningf(ctx, "Possible deadlock state with stack: %s", stack)
	}
}

func identifyDeadlocks(after time.Duration, fn func(string)) func() {
	if os.Getenv("DEPLOY_ENV") != "e2e" {
		return noop
	}

	ticker := time.NewTimer(after)
	done := make(chan struct{})

	go func() {
		select {
		case <-ticker.C:
			buf := make([]byte, 1024*1024*2) // limit log output to 2mb
			runtime.Stack(buf, true)
			fn(string(buf))
		case <-done:
		}
	}()

	return func() {
		ticker.Stop()
		close(done)
	}
}
