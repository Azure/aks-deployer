package panics

import (
	"context"
	"crypto/sha256"
	"fmt"
	"net/http"
	"runtime"
	"runtime/debug"
	"strings"

	"github.com/Azure/aks-deployer/pkg/httphandler"
	"github.com/Azure/aks-deployer/pkg/log"
)

// Recovery recovers from a panic and prints the logs as an error in the correct format. This should
// be deferred at the top of each go routine in order to not crash the entire process
// executes custom handlers so errors can be populated back to channels and other cleanup
func Recovery(ctx context.Context, name string, handlers ...func()) {
	if r := recover(); r != nil {
		Log(ctx, name, r)
		for _, handler := range handlers {
			handler()
		}
	}
}

// Identify returns the frame of the panic and hash of the calls stack which is useful for grouping panics
// it skips the stack that comes from the runtime package and recovery code run after the panic
// modeled after https://golang.org/pkg/runtime/#CallersFrames
func Identify() (runtime.Frame, string) {

	panicFrame := runtime.Frame{
		Function: "UNKNOWN",
		File:     "UNKNOWN",
	}
	pc := make([]uintptr, 16)               //how deep should we look? 16 a deep enough stack?
	n := runtime.Callers(2, pc)             // 3 == runtime.Callers, IdentifyPanic.
	frames := runtime.CallersFrames(pc[:n]) // pass only valid pcs to runtime.CallersFrames

	stackhash := sha256.New()
	foundPanic := false
	passedPanic := false
	for {
		frame, more := frames.Next()

		//ignore runtime at beginning and end of stack
		if strings.HasPrefix(frame.Function, "runtime.") {
			passedPanic = true
			continue
		}

		//could be other non runtime funcions before panic
		if !passedPanic {
			continue
		}

		if !foundPanic {
			//frame that raised the panic will be first after runtime
			//but the number of runtime lines can vary so can't just
			panicFrame = frame
			foundPanic = true
		}
		_, _ = stackhash.Write([]byte(fmt.Sprintf("%s%d\n", frame.Function, frame.Line)))

		if !more {
			break
		}
	}
	return panicFrame, fmt.Sprintf("%x", stackhash.Sum(nil))
}

//Log will use structured logging so we can find/group panics
func Log(ctx context.Context, name string, recovery interface{}) {
	logger := log.GetLogger(ctx)
	frame, hash := Identify()
	// using extra fields to group panics in kusto
	logger = logger.WithField("PanicStackHash", hash)
	logger = logger.WithField("PanicFileName", frame.File)
	logger = logger.WithField("PanicLineNumber", frame.Line)
	logger = logger.WithField("PanicFunctionName", frame.Function)
	logger.Errorf(ctx, "%s Panic: %+v, stacktrace: '%s'", name, recovery, string(debug.Stack()))
}

// HttpHandler handler wrapper to recover from panics
func HttpHandler(name string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.WriteHeader(http.StatusInternalServerError)
				Log(req.Context(), name, r)
			}
		}()
		next.ServeHTTP(w, req)
	})
}

func HttpHandlerWithMetrics(routeName string, next http.Handler, emitter httphandler.MetricEmitter) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				//unexpected http result
				w.WriteHeader(http.StatusInternalServerError)
				Log(req.Context(), routeName, r)
				emitter.UnexpectedResult()
			} else {
				// expected result
				emitter.ExpectedResult()
			}
		}()
		next.ServeHTTP(w, req)
	})
}
