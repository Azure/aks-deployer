package version

import (
	"fmt"
)

// Version ... These variables are populated at build time in the Makefile
var (
	Version   = "unset"
	GitBranch = "unset"
	GitHash   = "unset"
	BuildTime = "unset"
)

// String will return the version info as a raw string instead of a sugared logger.
// Use this in cases where you are not using JSON log formats (like a user-facing CLI tool).
func String() string {
	return fmt.Sprintf("Version: %s - Branch: %s - git SHA: %s - Build date / time: %s", Version, GitBranch, GitHash, BuildTime)
}

// ShortString wil return the version info as a short string instead of a sugared logger.
// In format of "{Version} {GitHash} {GitBranch} ({BuildTime})"
func ShortString() string {
	return fmt.Sprintf("%s %s %s (%s)", Version, GitHash, GitBranch, BuildTime)
}
