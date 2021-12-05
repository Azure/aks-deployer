package controllers

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	serviceVersionVec = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Help:      "The existing versions of AksApp components",
			Name:      "service_targeted_version",
			Namespace: "deployer",
			Subsystem: "deployer",
		},
		[]string{
			"name",
			"type",
			"version",
		},
	)

	secretVersionVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Help:      "The existing versions of secrets used by AksApp",
			Name:      "secret_version",
			Namespace: "deployer",
			Subsystem: "deployer",
		},
		[]string{
			"name",
			"secretName",
			"secretVersion",
		},
	)

	unavailableReplicasVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Help:      "The number of unavailable replicas",
			Name:      "unavailable_replicas",
			Namespace: "deployer",
			Subsystem: "deployer",
		},
		[]string{
			"name",
			"namespace",
		},
	)

	allReplicasVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Help:      "The number of all replicas",
			Name:      "all_replicas",
			Namespace: "deployer",
			Subsystem: "deployer",
		},
		[]string{
			"name",
			"namespace",
		},
	)

	releaseResultVec = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Help:      "The result of aksapp release",
			Name:      "app_release_result",
			Namespace: "deployer",
			Subsystem: "deployer",
		},
		[]string{
			"namespace",
			"name",
			"type",
			"result",
		},
	)
)

func init() {
	prometheus.MustRegister(serviceVersionVec)
	prometheus.MustRegister(secretVersionVec)
	prometheus.MustRegister(unavailableReplicasVec)
	prometheus.MustRegister(allReplicasVec)
	prometheus.MustRegister(releaseResultVec)
}
