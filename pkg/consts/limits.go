//------------------------------------------------------------
// Copyright (c) Microsoft Corporation.  All rights reserved.
//------------------------------------------------------------

package consts

const (
	// LegacyMaxNumberOfAPIServerAuthorizedIPRanges is the legacy default max number of APIServer authorized IP ranges.
	// Some subscriptions keep this number for backward compatibility.
	LegacyMaxNumberOfAPIServerAuthorizedIPRanges = 3500
	// MaxNumberOfAPIServerAuthorizedIPRanges is the default max number of APIServer authorized IP ranges.
	MaxNumberOfAPIServerAuthorizedIPRanges = 200
	// MaxClientSecretLength is the max length of SP client secret length in bytes
	MaxClientSecretLength = 190
)
