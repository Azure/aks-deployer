// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package configmaps

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/onsi/ginkgo/reporters"
)

func TestConfigMaps(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "ConfigMaps Suite", []Reporter{reporters.NewJUnitReporter("junit.xml")})
}
