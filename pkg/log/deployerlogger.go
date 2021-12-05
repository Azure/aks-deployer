package log

import (
	log "github.com/sirupsen/logrus"
)

// New returns a new logger
func New(serviceName, version string) *log.Entry {
	log.SetFormatter(&log.JSONFormatter{})
	return log.WithFields(log.Fields{
		"service": serviceName,
		"version": version,
	})
}
