// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

/*
Package datastructs contains commonly used data structs that don't naturally belong to
any other packages.
*/
package datastructs

// RedactedString represents a string whose Stringer implementation (String() func)
// always returns hard-coded <REDACTED> value instead of the actual string content.
// This is useful when you don't want the content to be accidentally written
// to log by %v or %+v, for example.
type RedactedString string

// Returns hard-coded "<REDACTED>" value.
func (RedactedString) String() string {
	return "<REDACTED>"
}

// Returns real string value
func (v RedactedString) RealString() string {
	return string(v)
}
