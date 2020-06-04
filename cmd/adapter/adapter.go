package main

import (
	"flag"
	"os"
	"time"

	"github.com/pkg/errors"
	"k8s.io/component-base/logs"
	"k8s.io/klog"

	clientset "github.com/awslabs/k8s-cloudwatch-adapter/pkg/client/clientset/versioned"
	informers "github.com/awslabs/k8s-cloudwatch-adapter/pkg/client/informers/externalversions"
	"github.com/awslabs/k8s-cloudwatch-adapter/pkg/controller"
	"github.com/awslabs/k8s-cloudwatch-adapter/pkg/metriccache"
	"github.com/awslabs/k8s-cloudwatch-adapter/pkg/newrelic"
	cwprov "github.com/awslabs/k8s-cloudwatch-adapter/pkg/provider"
	basecmd "github.com/kubernetes-incubator/custom-metrics-apiserver/pkg/cmd"
	"github.com/kubernetes-incubator/custom-metrics-apiserver/pkg/provider"
)

// CloudWatchAdapter represents a custom metrics BaseAdapter for Amazon CloudWatch
type NewRelicAdapter struct {
	basecmd.AdapterBase
}

func (a *NewRelicAdapter) makeCloudWatchClient() (newrelic.Client, error) {
	client := newrelic.NewRelicClient()
	return client, nil
}

func (a *NewRelicAdapter) newController(metriccache *metriccache.MetricCache) (*controller.Controller, informers.SharedInformerFactory) {
	clientConfig, err := a.ClientConfig()
	if err != nil {
		klog.Fatalf("unable to construct client config: %v", err)
	}
	adapterClientSet, err := clientset.NewForConfig(clientConfig)
	if err != nil {
		klog.Fatalf("unable to construct lister client to initialize provider: %v", err)
	}

	adapterInformerFactory := informers.NewSharedInformerFactory(adapterClientSet, time.Second*30)
	handler := controller.NewHandler(
		adapterInformerFactory.Metrics().V1alpha1().ExternalMetrics().Lister(),
		metriccache)

	controller := controller.NewController(adapterInformerFactory.Metrics().V1alpha1().ExternalMetrics(), &handler)

	return controller, adapterInformerFactory
}

func (a *NewRelicAdapter) makeProvider(cwClient newrelic.Client, metriccache *metriccache.MetricCache) (provider.ExternalMetricsProvider, error) {
	client, err := a.DynamicClient()
	if err != nil {
		return nil, errors.Wrap(err, "unable to construct Kubernetes client")
	}

	mapper, err := a.RESTMapper()
	if err != nil {
		return nil, errors.Wrap(err, "unable to construct RESTMapper")
	}

	cwProvider := cwprov.NewRelicProvider(client, mapper, cwClient, metriccache)
	return cwProvider, nil
}

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	// set up flags
	cmd := &NewRelicAdapter{}
	cmd.Name = "cloudwatch-metrics-adapter"
	cmd.Flags().AddGoFlagSet(flag.CommandLine) // make sure we get the klog flags
	cmd.Flags().Parse(os.Args)

	stopCh := make(chan struct{})
	defer close(stopCh)

	metriccache := metriccache.NewMetricCache()

	// start and run contoller components
	controller, adapterInformerFactory := cmd.newController(metriccache)
	go adapterInformerFactory.Start(stopCh)
	go controller.Run(2, time.Second, stopCh)

	// create CloudWatch client
	cwClient, err := cmd.makeCloudWatchClient()
	if err != nil {
		klog.Fatalf("unable to construct CloudWatch client: %v", err)
	}

	// construct the provider
	cwProvider, err := cmd.makeProvider(cwClient, metriccache)
	if err != nil {
		klog.Fatalf("unable to construct CloudWatch metrics provider: %v", err)
	}

	cmd.WithExternalMetrics(cwProvider)

	klog.Info("CloudWatch metrics adapter started")

	if err := cmd.Run(stopCh); err != nil {
		klog.Fatalf("unable to run CloudWatch metrics adapter: %v", err)
	}
}
