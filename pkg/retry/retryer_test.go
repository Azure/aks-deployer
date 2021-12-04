package retry_test

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"

	. "github.com/onsi/gomega"

	"github.com/Azure/aks-deployer/pkg/log"
	"github.com/Azure/aks-deployer/pkg/retry"
	"github.com/Azure/aks-deployer/pkg/retry/mock_retry"
)

type result struct {
	x int
}

var _ = Describe("Base retry tests", func() {
	var (
		mockCtrl        *gomock.Controller
		singleIteration *mock_retry.MockSingleIterationInterface
		ctx             context.Context
		logger          *log.Logger
		settings        retry.Settings
		obj             interface{}
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		singleIteration = mock_retry.NewMockSingleIterationInterface(mockCtrl)
		logger = log.InitializeTestLogger()
		ctx = log.WithLogger(context.Background(), logger)
		apiTracking := log.NewAPITrackingFromParametersMap(nil)
		ctx = log.WithAPITracking(ctx, apiTracking)
		settings = retry.Settings{
			Interval:      1 * time.Millisecond,
			Timeout:       5 * time.Millisecond,
			RetryMaxCount: 1,
		}
		obj = &result{x: 3}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	It("Succeeded without retries", func() {
		singleIteration.EXPECT().RunOnce(gomock.Any()).Return(retry.Success, obj, nil).Times(1)
		r := retry.NewRetry("base", settings, singleIteration)
		newobj, err := r.Run(ctx)
		Expect(err).To(BeNil())
		Expect(reflect.DeepEqual(newobj, obj)).To(BeTrue())
	})

	It("Succeeded with retries", func() {
		settings = retry.Settings{
			Interval:      1 * time.Millisecond,
			Timeout:       1000 * time.Millisecond,
			RetryMaxCount: 4,
		}

		before := singleIteration.EXPECT().RunOnce(gomock.Any()).Return(retry.NeedRetry, nil, nil).Times(3)
		singleIteration.EXPECT().RunOnce(gomock.Any()).Return(retry.Success, obj, nil).Times(1).After(before)
		r := retry.NewRetry("base", settings, singleIteration)

		newobj, err := r.Run(ctx)
		Expect(err).To(BeNil())
		Expect(reflect.DeepEqual(newobj, obj)).To(BeTrue())
	})

	It("Failed without retries", func() {
		singleIteration.EXPECT().RunOnce(gomock.Any()).Return(retry.Failed, nil, nil).Times(1)
		r := retry.NewRetry("base", settings, singleIteration)
		newobj, err := r.Run(ctx)
		Expect(err.Error()).To(Equal("base retry failed: %!w(<nil>)"))
		Expect(newobj).To(BeNil())
	})

	It("Failed with retries", func() {
		settings = retry.Settings{
			Interval:      1 * time.Millisecond,
			Timeout:       1000 * time.Millisecond,
			RetryMaxCount: 4,
		}

		before := singleIteration.EXPECT().RunOnce(gomock.Any()).Return(retry.NeedRetry, nil, nil).Times(3)
		singleIteration.EXPECT().RunOnce(gomock.Any()).Return(retry.Failed, nil, nil).Times(1).After(before)

		r := retry.NewRetry("base", settings, singleIteration)
		newobj, err := r.Run(ctx)
		Expect(err.Error()).To(Equal("base retry failed: %!w(<nil>)"))
		Expect(newobj).To(BeNil())
	})

	It("Time out - exceeded retries number", func() {
		settings = retry.Settings{
			Interval:      1 * time.Millisecond,
			Timeout:       100 * time.Millisecond,
			RetryMaxCount: 2,
		}

		singleIteration.EXPECT().RunOnce(gomock.Any()).Return(retry.NeedRetry, nil, fmt.Errorf("failed")).MinTimes(1)

		r := retry.NewRetry("base", settings, singleIteration)
		newobj, err := r.Run(ctx)
		Expect(err.Error()).To(Equal("base retry timed out: failed"))
		Expect(newobj).To(BeNil())
	})

	It("time out - real time out", func() {
		settings = retry.Settings{
			Interval:      1 * time.Millisecond,
			Timeout:       2 * time.Millisecond,
			RetryMaxCount: 1000,
		}

		singleIteration.EXPECT().RunOnce(gomock.Any()).Return(retry.NeedRetry, nil, nil).MinTimes(1)

		r := retry.NewRetry("base", settings, singleIteration)
		newobj, err := r.Run(ctx)
		Expect(err.Error()).To(Equal("base retry timed out: %!w(<nil>)"))
		Expect(newobj).To(BeNil())
	})

	It("name cannot be empty", func() {
		r := retry.NewRetry("", settings, singleIteration)
		newobj, err := r.Run(ctx)
		Expect(err.Error()).To(Equal("The retry name cannot be empty"))
		Expect(newobj).To(BeNil())
	})

	It("The singleIteration cannot be nil", func() {
		r := retry.NewRetry("base", settings, nil)
		newobj, err := r.Run(ctx)
		Expect(err.Error()).To(Equal("The singleIteration cannot be nil"))
		Expect(newobj).To(BeNil())
	})

	It("should stop retry on context cancel", func() {
		ctx, cancelFunc := context.WithCancel(ctx)
		counter := 0
		// setup 5 time retry
		settings.RetryMaxCount = 5
		settings.Timeout = time.Second
		singleIteration.EXPECT().RunOnce(gomock.Any()).DoAndReturn(func(ctx context.Context) (retry.Status, interface{}, error) {
			counter++
			if counter > 2 {
				cancelFunc()
			}
			if counter > 5 {
				return retry.Success, nil, nil
			}
			return retry.NeedRetry, nil, nil
		}).AnyTimes()

		r := retry.NewRetry("base", settings, singleIteration)
		newobj, err := r.Run(ctx)
		Expect(errors.Is(err, context.Canceled)).To(BeTrue())
		Expect(newobj).To(BeNil())
		// we actually invoked the function.
		Expect(counter > 0).To(BeTrue())
		Expect(counter < 5).To(BeTrue())
	})
})
