package retry

import (
	"errors"
	"fmt"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type dummy string

type MockTimex struct{}

func (dt *MockTimex) Sleep(d time.Duration) {
	fmt.Println("Sleep: ", d)
}

func (dt *MockTimex) Now() time.Time {
	return time.Now()
}

var _ = Describe("Retry", func() {

	mocktime := &MockTimex{}

	It("should have linear delay", func() {
		Expect(getLinearDelay(1, time.Second)).Should(Equal(time.Second))
		Expect(getLinearDelay(2, time.Second)).Should(Equal(time.Second * 2))
		Expect(getLinearDelay(3, time.Second)).Should(Equal(time.Second * 3))
		Expect(getFixedDelay(3, time.Second)).Should(Equal(time.Second))
	})

	It("should have jittery delay", func() {
		Expect(getJitteryDelay(1, 5*time.Second)).Should(BeNumerically("<=", 5*time.Second))
		Expect(getJitteryDelay(5, time.Second)).Should(BeNumerically("<=", time.Second))
	})

	It("retry_ success and no retries", func() {
		count := 0
		res, err := doLinearRetry(
			func() Result {
				count++
				return Result{Success, dummy("hello")}
			},
			time.Millisecond,
			time.Second,
			mocktime,
		)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(count).Should(Equal(1))
		Expect(res.Status).Should(Equal(Success))
		Expect(res.Body).ShouldNot(BeNil())
		Expect(res.Body).Should(BeAssignableToTypeOf(dummy("hello")))
		Expect(res.Body.(dummy)).Should(Equal(dummy("hello")))
	})

	It("retry_ failed and no retries", func() {
		count := 0
		res, err := doLinearRetry(
			func() Result {
				count++
				return Result{Failed, dummy("hello")}
			},
			time.Millisecond,
			time.Second,
			mocktime,
		)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(count).Should(Equal(1))
		Expect(res.Status).Should(Equal(Failed))
		Expect(res.Body).ShouldNot(BeNil())
		Expect(res.Body).Should(BeAssignableToTypeOf(dummy("hello")))
		Expect(res.Body.(dummy)).Should(Equal(dummy("hello")))
	})

	It("retry_ with retries", func() {
		count := 0
		res, err := doLinearRetry(
			func() Result {
				if count++; count == 3 {
					return Result{Status: Success}
				}
				return Result{Status: NeedRetry}
			},
			time.Millisecond,
			time.Second,
			mocktime,
		)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(res.Status).Should(Equal(Success))
		Expect(count).Should(Equal(3))
	})

	It("retry_ timedout by duration", func() {
		count := 0
		res, err := doLinearRetry(
			func() Result {
				count++
				time.Sleep(time.Millisecond * 50)
				return Result{Status: NeedRetry}
			},
			time.Millisecond,
			time.Millisecond*20,
			mocktime,
		)

		Expect(err).ShouldNot(HaveOccurred())

		time.Sleep(time.Millisecond * 100)
		Expect(res.Status).Should(Equal(Timedout))
		Expect(count).Should(Equal(1))
	})

	It("retry_ timedout by max count", func() {
		count := 0
		res, err := doLinearRetryWithMaxCount(
			func() Result {
				count++
				return Result{Status: NeedRetry, Body: errors.New("error")}
			},
			time.Millisecond,
			time.Second*10,
			3,
			mocktime,
		)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(res.Status).Should(Equal(Timedout))
		Expect(res.Body).ShouldNot(BeNil())
		e := res.Body.(error)
		Expect(e).ShouldNot(BeNil())
		Expect(e.Error()).Should(Equal("error"))
		Expect(count).Should(Equal(3))
	})

	It("retry_ negative tests", func() {
		_, err := doLinearRetryWithMaxCount(nil, 1, time.Second*10, 3, mocktime)
		Expect(err).Should(HaveOccurred())

		_, err = doLinearRetry(nil, 1, time.Second*10, mocktime)
		Expect(err).Should(HaveOccurred())

		count := 0
		_, err = doLinearRetryWithMaxCount(
			func() Result {
				count++
				return Result{Status: NeedRetry}
			},
			-1,
			time.Second*10,
			3,
			mocktime,
		)

		Expect(err).Should(HaveOccurred())
		Expect(count).Should(Equal(0))

		_, err = doLinearRetryWithMaxCount(
			func() Result {
				count++
				return Result{Status: NeedRetry}
			},
			time.Second,
			-1,
			3,
			mocktime,
		)

		Expect(err).Should(HaveOccurred())
		Expect(count).Should(Equal(0))

		_, err = doLinearRetryWithMaxCount(
			func() Result {
				count++
				return Result{Status: NeedRetry}
			},
			10,
			1,
			-1,
			mocktime,
		)

		Expect(err).Should(HaveOccurred())
		Expect(count).Should(Equal(0))

		_, err = DoLinearRetry(
			func() Result {
				count++
				return Result{Status: NeedRetry}
			},
			time.Second*10,
			-1)

		Expect(err).Should(HaveOccurred())
		Expect(count).Should(Equal(0))

		_, err = doLinearRetry(
			func() Result {
				count++
				return Result{Status: NeedRetry}
			},
			-1,
			time.Second*10,
			mocktime,
		)

		Expect(err).Should(HaveOccurred())
		Expect(count).Should(Equal(0))
	})

	It("flat retry_ timedout", func() {
		count := 0
		res, err := doFixedRetryWithMaxCount(
			func() Result {
				count++
				return Result{Status: NeedRetry, Body: errors.New("error")}
			},
			time.Millisecond,
			50*time.Millisecond,
			10,
			mocktime,
		)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(res.Status).Should(Equal(Timedout))
		Expect(res.Body).Should(HaveOccurred())
		Expect(count).Should(BeNumerically(">=", 3))
	})

	It("linear retry_ timedout", func() {
		count := 0
		res, err := doLinearRetryWithMaxCount(
			func() Result {
				count++
				return Result{Status: NeedRetry, Body: errors.New("error")}
			},
			time.Millisecond,
			5*time.Millisecond,
			10,
			mocktime,
		)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(res.Status).Should(Equal(Timedout))
		Expect(count).Should(BeNumerically("<=", 10))
	})

	It("should wait the expected interval between attempts", func() {
		interval := time.Millisecond

		firstTry := time.Now()
		tryCount := 0
		var calls []time.Time
		fn := func() Result {
			calls = append(calls, firstTry.Add(interval*time.Duration(tryCount)))
			return Result{Status: NeedRetry, Body: nil}
		}

		delayFunc := func(i int, iv time.Duration) time.Duration {
			tryCount++
			Expect(i).To(BeNumerically("<", 2))
			Expect(iv).To(Equal(interval))
			return interval
		}

		_, err := doRetryWithMaxCount(fn, interval, time.Hour, 2, delayFunc, mocktime)
		Expect(err).ToNot(HaveOccurred())
		Expect(calls).To(HaveLen(2))
		Expect(calls[0].Add(interval)).To(BeTemporally("<=", calls[1]))
	})

	Context("the timeout has expired", func() {
		When("the function returns success", func() {
			It("should not return an error", func() {
				fn := func() Result {
					// Return after timeout by waiting 1ms for the 1ns timeout
					time.Sleep(time.Millisecond)
					return Result{Status: Success, Body: nil}
				}

				result, err := doLinearRetry(fn, time.Hour, time.Nanosecond, mocktime)
				Expect(err).ToNot(HaveOccurred())
				Expect(result.Status).To(Equal(Success))
				Expect(result.Body).To(BeNil())
			})
		})
	})

	Context("the max retries has been reached", func() {
		When("the function returns success", func() {
			It("should not return an error", func() {
				fn := func() Result {
					return Result{Status: Success, Body: nil}
				}

				result, err := doLinearRetryWithMaxCount(fn, time.Nanosecond, time.Hour, 1, mocktime)
				Expect(err).ToNot(HaveOccurred())
				Expect(result.Status).To(Equal(Success))
				Expect(result.Body).To(BeNil())
			})
		})
	})
})
