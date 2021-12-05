package panics

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/sirupsen/logrus"
	"github.com/Azure/aks-deployer/pkg/httphandler"
	"github.com/Azure/aks-deployer/pkg/log"
)

const panicLine = 21

//imporant this doesn't move. Can put in a seperate file if that becomes a problem
func createPanic() {
	var s *string
	fmt.Println(*s)
}

func foobar() {
	defer func() {
		fmt.Println("leaving foobar")
	}()
	bar()
}

func bar() {
	defer func() {
		fmt.Println("leaving bar")
	}()
	createPanic()
}

type testdummyemitter struct{}

func (t *testdummyemitter) ExpectedResult()   {}
func (t *testdummyemitter) UnexpectedResult() {}
func (t *testdummyemitter) Heartbeat()        {}

var _ httphandler.MetricEmitter = &testdummyemitter{}

var _ = Describe("Testing Panic Utilities", func() {

	It("Testing Go Routine Recovery Function", func() {
		ctx := context.Background()
		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			defer wg.Done()
			defer Recovery(ctx, "testing-panic")
			panic("should not see this panic")
		}()

		wg.Wait()
	})

	It("Testing Go Routine Recovery Function With Handler", func() {
		handlerExecuted := false
		handler := func() {
			handlerExecuted = true
		}

		ctx := context.Background()
		var wg sync.WaitGroup

		wg.Add(1)

		go func() {
			defer wg.Done()
			defer Recovery(ctx, "testing-panic-and-handle", handler)
			panic("should not see this panic")
		}()

		wg.Wait()

		Expect(handlerExecuted).To(BeTrue())
	})

	It("Can Identify panic Line", func() {
		defer func() {
			r := recover()
			Expect(r).ShouldNot(BeNil())
			frame, hash := Identify()
			Expect(frame.Line).To(Equal(panicLine))
			Expect(frame.File).To(ContainSubstring("panics_test.go"))
			Expect(hash).ToNot(BeEmpty())
		}()
		foobar()
	})

	It("Can differentiate stacks", func() {
		var stackhashes [2]string
		for i := 0; i < 2; i++ {
			stackhashes[i] = GrabStackHash()
		}
		Expect(stackhashes[0]).To(Equal(stackhashes[1]))

		differenthash := GrabStackHash()

		Expect(stackhashes[0]).ToNot(Equal(differenthash))
	})

	It("Testing handler with metrics", func() {
		var handler http.Handler

		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("test")
		})

		handler = HttpHandlerWithMetrics("recoverWithMetrics", handler, &testdummyemitter{})

		req := httptest.NewRequest("GET", "/", nil)
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		Expect(rw.Code).To(Equal(http.StatusInternalServerError))
	})

	It("Handler with Metricsshould not interrupt normal response", func() {
		var handler http.Handler

		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(201)
			w.Write([]byte("test"))
		})

		handler = HttpHandlerWithMetrics("recover", handler, &testdummyemitter{})

		req := httptest.NewRequest("GET", "/", nil)
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		Expect(rw.Code).To(Equal(201))
		Expect(rw.Body).NotTo(BeNil())
	})

	It("should capture panic", func() {
		var handler http.Handler

		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			panic("test")
		})

		handler = HttpHandler("recover", handler)

		req := httptest.NewRequest("GET", "/", nil)
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		Expect(rw.Code).To(Equal(http.StatusInternalServerError))
	})

	It("should not interrupt normal response", func() {
		var handler http.Handler

		handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(201)
			w.Write([]byte("test"))
		})

		handler = HttpHandler("recover", handler)

		req := httptest.NewRequest("GET", "/", nil)
		rw := httptest.NewRecorder()

		handler.ServeHTTP(rw, req)

		Expect(rw.Code).To(Equal(201))
		Expect(rw.Body).NotTo(BeNil())
	})

	It("Log Panic Syntax is legit", func() {
		bufferlogger := logrus.New()
		bufferlogger.Level = logrus.DebugLevel
		buf := &bytes.Buffer{}
		bufferlogger.SetOutput(buf)
		loggerEntry := bufferlogger.WithField("unittest", "true")
		testlogger := &log.Logger{
			loggerEntry,
			loggerEntry,
		}
		ctx := log.WithLogger(context.TODO(), testlogger)

		defer func() {
			r := recover()
			Expect(r).ShouldNot(BeNil())
			Log(ctx, "UNITTEST", r)
			Expect(buf.String()).Should(ContainSubstring("PanicLineNumber=21"))
			Expect(buf.String()).Should(ContainSubstring("PanicFileName="))
			Expect(buf.String()).Should(ContainSubstring("panics_test.go"))
			Expect(buf.String()).Should(ContainSubstring("PanicFunctionName="))
			Expect(buf.String()).Should(ContainSubstring("createPanic"))
			//too fragile? will change if oackage name chagnes
			//Expect(buf.String()).Should(ContainSubstring("PanicStackHash=428d06a41d87780e0ed18f6d2a382f41"))
		}()
		foobar()

	})
})

func GrabStackHash() (hash string) {
	hash = "shouldneverhappen"
	defer func() {
		r := recover()
		Expect(r).ShouldNot(BeNil())
		_, hash = Identify()
	}()
	foobar()
	return hash
}
