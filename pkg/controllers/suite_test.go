package controllers

import (
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	// +kubebuilder:scaffold:imports
)

// +kubebuilder:docs-gen:collapse=Imports

// Setup environment variables
var _ = func() error {
	os.Setenv(identityResourceIDStr, "test-identity-resource-id")
	os.Setenv(keyVaultResourceStr, "test-key-vault-resource")
	return nil
}()

func TestControllers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t,
		"Controllers Suite",
		[]Reporter{reporters.NewJUnitReporter("junit.xml")})
}
