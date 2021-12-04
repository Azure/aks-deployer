// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package deploy

import (
	"os"
	"testing"
)

func TestGetLoggingRegion(t *testing.T) {
	getLoggingRegionTestCase("prod", "eastus", "eastus", t)
	getLoggingRegionTestCase("int", "eastus", "int-eastus", t)
	getLoggingRegionTestCase("staging", "westus2", "staging-westus2", t)
	getLoggingRegionTestCase("", "eastus", "eastus", t)
	getLoggingRegionTestCase("prod", "", "", t)
	getLoggingRegionTestCase("e2e", "westcentralus", "e2e-westcentralus", t)
}

func getLoggingRegionTestCase(environment, region, expected string, t *testing.T) {
	// reset the singleton for this unit test
	Reset()
	_ = os.Setenv(EnvKeyDeployEnv, environment)
	if actual := GetLoggingRegion(region); actual != expected {
		t.Errorf("GetLoggingRegion result was %s, expecting %s", actual, expected)
	}
}

func TestDeploymentConfig_IsE2E(t *testing.T) {
	cases := []struct {
		config   *DeploymentConfig
		expected bool
	}{
		{
			config:   &DeploymentConfig{},
			expected: false,
		},
		{
			config: &DeploymentConfig{
				DeployEnv: E2EEnv,
			},
			expected: true,
		},
	}

	for _, c := range cases {
		if v := c.config.IsE2E(); v != c.expected {
			t.Errorf(
				"for IsE2E %s, expect %t, got %t",
				c.config.DeployEnv, c.expected, v,
			)
		}
	}
}
