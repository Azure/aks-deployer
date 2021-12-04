package timing

import (
	"context"
	"os"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("IdentifyDeadlocks", func() {
	When("enabled", func() {
		BeforeEach(func() {
			Expect(os.Setenv("DEPLOY_ENV", "e2e")).To(Succeed())
		})

		AfterEach(func() {
			Expect(os.Setenv("DEPLOY_ENV", "prod")).To(Succeed())
		})

		When("the duration is not reached", func() {
			It("should not call the hook", func() {
				defer identifyDeadlocks(time.Hour, func(stack string) {
					Fail("hook was called")
				})()
			})
		})

		When("the duration has been reached", func() {
			It("should call the hook with the stack trace", func() {
				done := make(chan struct{})
				defer identifyDeadlocks(time.Millisecond, func(stack string) {
					close(done)
				})()

				select {
				case <-time.After(time.Second):
					Fail("hook wasn't called after 1s")
				case <-done:
					// Hook was called!
				}
			})
		})

		It("shouldn't panic", func() {
			defer IdentifyDeadlocks(context.Background(), 1)()
			time.Sleep(time.Millisecond)
		})
	})

	When("disabled", func() {
		When("the duration is not reached", func() {
			It("should not call the hook", func() {
				defer identifyDeadlocks(time.Hour, func(stack string) {
					Fail("hook was called")
				})()
			})
		})

		When("the duration has been reached", func() {
			It("should not call the hook", func() {
				defer identifyDeadlocks(time.Millisecond, func(stack string) {
					Fail("hook was called")
				})()
				time.Sleep(time.Millisecond * 100)
			})
		})
	})
})
