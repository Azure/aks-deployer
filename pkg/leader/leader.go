package leader

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/leaderelection"
	"k8s.io/client-go/tools/leaderelection/resourcelock"
)

type RunFunc func(context.Context)

// RunWithLeaderElection comments
func RunWithLeaderElection(fn RunFunc, logger *logrus.Entry, ns, name string) {
	config, err := rest.InClusterConfig()
	if err != nil {
		logger.Errorf("unable to get in cluster config, %s", err.Error())
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Errorf("unable to create clientset, %s", err.Error())
		os.Exit(1)
	}

	id, err := os.Hostname()
	if err != nil {
		logger.Errorf("unable to get hostname, %s", err.Error())
		os.Exit(1)
	}

	rl, err := resourcelock.New(resourcelock.LeasesResourceLock,
		ns,
		name,
		clientset.CoreV1(),
		clientset.CoordinationV1(),
		resourcelock.ResourceLockConfig{
			Identity: id,
		})
	if err != nil {
		logger.Errorf("unable to create a resourcelock, %s", err.Error())
		os.Exit(1)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-ch
		logger.Info("receive termination, signaling shutdown")
		cancel()
	}()

	leaderelection.RunOrDie(ctx, leaderelection.LeaderElectionConfig{
		Lock:          rl,
		LeaseDuration: 30 * time.Second,
		RenewDeadline: 15 * time.Second,
		RetryPeriod:   5 * time.Second,
		Callbacks: leaderelection.LeaderCallbacks{
			OnStartedLeading: fn,
			OnStoppedLeading: func() {
				logger.Warnf("leader lost: %s", id)
				os.Exit(0)
			},
			OnNewLeader: func(identity string) {
				logger.Infof("new leader elected: %s", identity)
			},
		},
	})
}
