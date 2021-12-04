//------------------------------------------------------------
// Copyright (c) Microsoft Corporation.  All rights reserved.
//------------------------------------------------------------

// fakelogger.go contains the code for making a logger context for a test

package log

import (
	"fmt"

	logrus "github.com/sirupsen/logrus"
)

var counter = 0

// InitializeTestLogger initializes a logger that can be used in a test
func InitializeTestLogger() *Logger {
	counter++
	// stop flood the screen with logs by print fatal level only.
	// but you can change your specific test to trace Info/Error level, for local troubleshoot only.
	// please check in with this default level.
	if counter%100 == 1 {
		fmt.Println("!!      update log level at pkg/core/log/fakelogger.go if you are looking for traces     !!")
	}
	return InitializeTestLoggerWithLevel(logrus.ErrorLevel)
}

// InitializeTestLoggerWithLevel initializes a logger that can be used in a test, but with specified trace level
func InitializeTestLoggerWithLevel(level logrus.Level) *Logger {
	retVal := NewLogger(&APITracking{})

	retVal.QosLogger.Logger.SetLevel(level)
	retVal.TraceLogger.Logger.SetLevel(level)

	return retVal
}
