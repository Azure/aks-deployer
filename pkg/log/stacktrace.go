package log

import (
	"fmt"

	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func getStackTraceStr(err error) string {
	stErr := getBaseStackTracer(err)
	if stErr != nil {
		st := stErr.StackTrace()
		return fmt.Sprintf("%+v", st)
	}

	return ""
}

func getBaseStackTracer(err error) stackTracer {
	type unwrapper interface {
		Unwrap() error
	}

	var baseST stackTracer

	for err != nil {
		if st, ok := err.(stackTracer); ok {
			baseST = st
		}

		unwrappedErr, ok := err.(unwrapper)
		if !ok {
			break
		}
		err = unwrappedErr.Unwrap()
	}
	return baseST
}
