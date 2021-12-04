// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package deploy

import (
	"os"
	"sync"
	"sync/atomic"
)

// DeploymentConfig contains the deployment configurations.
type DeploymentConfig struct {
	// DeployEnv should be in [prod,e2e,...]
	DeployEnv string
	E2EConfig E2EConfig
}

// IsE2E tells if current deployment environment is E2E.
func (d *DeploymentConfig) IsE2E() bool {
	return d.DeployEnv == E2EEnv
}

// E2EConfig contains the configs for our E2E.
type E2EConfig struct {
	// VersionString is the unique string we are using to separate the different
	// deployments in our E2E environment.
	VersionString string
}

var initialized uint32
var mu sync.Mutex
var instance *DeploymentConfig

const (
	// EnvKeyDeployEnv sets the environment variable name for deoloy env.
	EnvKeyDeployEnv string = "DEPLOY_ENV"
	// EnvKeyE2EVersionString sets the environment variable name for E2E version string.
	EnvKeyE2EVersionString = "E2E_VERSION_STRING"

	// E2EEnv is the const string for the E2E environment.
	E2EEnv string = "e2e"
	// ProdEnv is the const string for the production environment.
	ProdEnv    string = "prod"
	INTEnv     string = "int"
	INTv2Env   string = "intv2"
	StagingEnv string = "staging"
)

// GetDeploymentConfig returns the singleton of the config.
func GetDeploymentConfig() *DeploymentConfig {

	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	mu.Lock()
	defer mu.Unlock()

	if initialized == 0 {
		deployEnv := os.Getenv(EnvKeyDeployEnv)
		if deployEnv == "" {
			deployEnv = ProdEnv
		}
		e2eVersionString := os.Getenv(EnvKeyE2EVersionString)
		instance = &DeploymentConfig{
			DeployEnv: deployEnv,
			E2EConfig: E2EConfig{
				VersionString: e2eVersionString,
			},
		}
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

// Reset is used to allow test to reset the context
func Reset() {
	// reset the singleton for this unit test
	atomic.StoreUint32(&initialized, 0)
}

// GetLoggingRegion returns the aks region for logs (int-eastus; staging-westus2)
func GetLoggingRegion(region string) string {
	loggingRegion := ""
	environment := GetDeploymentConfig().DeployEnv
	if environment == ProdEnv {
		loggingRegion = region
	} else {
		loggingRegion = environment + "-" + region
	}
	return loggingRegion
}
