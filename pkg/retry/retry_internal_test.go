package retry

import (
	"context"
	"time"

	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"

	"github.com/Azure/aks-deployer/pkg/log"
)

var _ = Describe("retry internal tests", func() {
	var (
		ctx    context.Context
		logger *log.Logger
	)

	BeforeEach(func() {
		logger = log.InitializeTestLogger()
		ctx = log.WithLogger(context.Background(), logger)
		apiTracking := log.NewAPITrackingFromParametersMap(nil)
		ctx = log.WithAPITracking(ctx, apiTracking)
	})

	AfterEach(func() {
	})

	It("interval is not set", func() {
		settings := Settings{
			Timeout:       3 * time.Second,
			RetryMaxCount: 1000,
		}

		settings = prepareSetting(ctx, "base", settings)
		Expect(settings.Interval).To(Equal(5 * time.Second))
		Expect(settings.Timeout).To(Equal(3 * time.Second))
		Expect(settings.RetryMaxCount).To(Equal(1000))
	})

	It("timeout is not set", func() {
		settings := Settings{
			Interval:      3 * time.Second,
			RetryMaxCount: 1000,
		}

		settings = prepareSetting(ctx, "base", settings)
		Expect(settings.Interval).To(Equal(3 * time.Second))
		Expect(settings.Timeout).To(Equal(10 * time.Minute))
		Expect(settings.RetryMaxCount).To(Equal(1000))
	})

	It("retryMaxCount is not set", func() {
		settings := Settings{
			Interval: 5 * time.Second,
			Timeout:  10 * time.Minute,
		}

		settings = prepareSetting(ctx, "base", settings)
		Expect(settings.Interval).To(Equal(5 * time.Second))
		Expect(settings.Timeout).To(Equal(10 * time.Minute))
		Expect(settings.RetryMaxCount).To(Equal(10))
	})
})
