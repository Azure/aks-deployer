// +build !localbuild

package log

import "github.com/sirupsen/logrus"

func init() {
	logger = newLogger(&logrus.JSONFormatter{})
}
