# AKS versioning

Standardized version string package for AKS.

Set the semver in the makefile for the service.

Add the following to the Makefile for the service and add `-ldflags $(LD_FLAGS)` to the `go build` call in the Makefile:
```
VERSION         ?= 1.0.0
GIT_HASH        ?= $(shell git rev-parse --short HEAD)
BUILD_TIME      ?= $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
LD_FLAGS        ?= -X goms.io/aks/rp/core/version.Version=$(VERSION) -X goms.io/aks/rp/core/version.GitHash=$(GIT_HASH) -X goms.io/aks/rp/core/version.BuildTime=$(BUILD_TIME)
```

To use this in a CLI application call `version.String()` inside any string printer.

To use this in a JSON logging service call `version.Print(<your_logger_here>)`.