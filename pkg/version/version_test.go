package version

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("version package test", func() {
	AfterEach(func() {
		Version = "unset"
		GitBranch = "unset"
		GitHash = "unset"
		BuildTime = "unset"
	})

	Context("version string", func() {

		It("should return correct version string", func() {
			Version = "3.14.15"
			GitBranch = "master"
			GitHash = "abc123"
			BuildTime = "Some point in time"
			Expect(String()).Should(Equal("Version: 3.14.15 - Branch: master - git SHA: abc123 - Build date / time: Some point in time"))
		})

		It("should return correct short version", func() {
			Version = "3.14.15"
			GitBranch = "master"
			GitHash = "abc123"
			BuildTime = "Some point in time"

			Expect(ShortString()).Should(Equal("3.14.15 abc123 master (Some point in time)"))
		})
	})
})
