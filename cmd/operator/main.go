// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package main

import (
	"net/http"
	"flag"
	"context"
	"time"
	"os"

	// Import all Kubernetes client auth plugins (e.g. Azure, GCP, OIDC, etc.)
	// to ensure that exec-entrypoint and run can make use of them.
	_ "k8s.io/client-go/plugin/pkg/client/auth"

	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/Azure/aks-deployer/pkg/configmaps"
	deployerv1 "github.com/Azure/aks-deployer/pkg/api/v1"
	"github.com/Azure/aks-deployer/pkg/log"
	"github.com/Azure/aks-deployer/pkg/leader"
	"github.com/sirupsen/logrus"
	"github.com/Azure/aks-deployer/pkg/version"
	"github.com/Azure/aks-deployer/pkg/controllers"
	"github.com/Azure/aks-deployer/pkg/panics"
	//+kubebuilder:scaffold:imports
)

var (
	scheme   = runtime.NewScheme()

	namespace string
	logger    *logrus.Entry

	metricsPort string

	reconcileSyncPeriod = time.Hour

	useOwnerReference = false
)

func init() {
	flag.StringVar(&namespace, "namespace", configmaps.DefaultDeployerNamespace, "namespace")
	flag.StringVar(&metricsPort, "listen", ":8080", "metrics port")
	flag.BoolVar(&useOwnerReference, "ownerreference", false, "use ownerreference")

	// logger setup
	// note: source field is used in fluentd to match rules and add tags
	logger = log.New("operator", version.String()).WithField("source", "Deployer")

	utilruntime.Must(clientgoscheme.AddToScheme(scheme))
	utilruntime.Must(deployerv1.AddToScheme(scheme))
	//+kubebuilder:scaffold:scheme
}

func main() {
	flag.Parse()

	mgr, err := ctrl.NewManager(ctrl.GetConfigOrDie(), ctrl.Options{
		Scheme:             scheme,
		MetricsBindAddress: "0",	// why "0", not metricsPort? Does this mean metrics are not enabled yet?
		SyncPeriod:         &reconcileSyncPeriod,
	})
	if err != nil {
		logger.Errorf("unable to start manager, %v", err.Error())
		os.Exit(1)
	}

	if useOwnerReference {
		logger.Info("Use OwnerReference to set the value for objects created by deployer")
	} else {
		logger.Info("Use Annotation to set the value for objects created by deployer")
	}

	aksAppReconciler := controllers.NewAksAppReconciler(mgr.GetClient(),
	logger.WithField("controller", "AksApp"), mgr.GetScheme(), namespace, useOwnerReference)

	if err = aksAppReconciler.SetupWithManager(mgr); err != nil {
		logger.Errorf("unable to create AksApp controller, %v", err.Error())
		os.Exit(1)
	}

	configMapReconciler := &controllers.ConfigMapReconciler{
		Client:    mgr.GetClient(),
		Logger:    logger.WithField("controller", "ConfigMap"),
		Scheme:    mgr.GetScheme(),
		Namespace: namespace,
	}

	if err = configMapReconciler.SetupWithManager(mgr); err != nil {
		logger.Errorf("unable to create ConfigMap controller, %v", err.Error())
		os.Exit(1)
	}

	run := func(ctx context.Context) {
		go serveMetrics(logger)

		go aksAppReconciler.MonitorRoutine()
		go aksAppReconciler.MetricResetRoutine()

		// +kubebuilder:scaffold:builder

		logger.Info("starting manager")
		if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
			logger.Errorf("problem running manager, %s", err.Error())
			os.Exit(1)
		}
	}

	leader.RunWithLeaderElection(run, logger, namespace, "operator")
}

func serveMetrics(logger *logrus.Entry) {
	http.Handle("/healthz", panics.HttpHandler("healthz", http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte("OK")); err != nil {
			logger.Errorf("Failed to  ... %s", err.Error())
		}
	})))

	http.Handle("/metrics", panics.HttpHandler("metrics", promhttp.Handler()))
	if err := http.ListenAndServe(metricsPort, nil); err != nil {
		logger.Errorf("problem running metrics server, %s", err.Error())
		os.Exit(1)
	}
}
