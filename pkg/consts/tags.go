//------------------------------------------------------------
// Copyright (c) Microsoft Corporation.  All rights reserved.
//------------------------------------------------------------

package consts

// Azure resource tag limitations
const (
	MaxTagNameLength    = 512
	MaxTagValueLength   = 256
	MaxNumberOfTags     = 50
	InvalidTagNameChars = "<>%&\\?/"

	K8sVersion            = "orchestrator"
	K8sVersionValuePrefix = "Kubernetes:"
)
