// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License").
// You may not use this file except in compliance with the License.
// A copy of the License is located at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kuperiu/k8s-newrelic-adapter/pkg/apis/metrics/v1alpha1"
	scheme "github.com/kuperiu/k8s-newrelic-adapter/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ExternalMetricsGetter has a method to return a ExternalMetricInterface.
// A group's client should implement this interface.
type ExternalMetricsGetter interface {
	ExternalMetrics(namespace string) ExternalMetricInterface
}

// ExternalMetricInterface has methods to work with ExternalMetric resources.
type ExternalMetricInterface interface {
	Create(*v1alpha1.ExternalMetric) (*v1alpha1.ExternalMetric, error)
	Update(*v1alpha1.ExternalMetric) (*v1alpha1.ExternalMetric, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ExternalMetric, error)
	List(opts v1.ListOptions) (*v1alpha1.ExternalMetricList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	ExternalMetricExpansion
}

// externalMetrics implements ExternalMetricInterface
type externalMetrics struct {
	client rest.Interface
	ns     string
}

// newExternalMetrics returns a ExternalMetrics
func newExternalMetrics(c *MetricsV1alpha1Client, namespace string) *externalMetrics {
	return &externalMetrics{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the externalMetric, and returns the corresponding externalMetric object, and an error if there is any.
func (c *externalMetrics) Get(name string, options v1.GetOptions) (result *v1alpha1.ExternalMetric, err error) {
	result = &v1alpha1.ExternalMetric{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("externalmetrics").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(context.Background()).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ExternalMetrics that match those selectors.
func (c *externalMetrics) List(opts v1.ListOptions) (result *v1alpha1.ExternalMetricList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ExternalMetricList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("externalmetrics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(context.Background()).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested externalMetrics.
func (c *externalMetrics) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("externalmetrics").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(context.Background())
}

// Create takes the representation of a externalMetric and creates it.  Returns the server's representation of the externalMetric, and an error, if there is any.
func (c *externalMetrics) Create(externalMetric *v1alpha1.ExternalMetric) (result *v1alpha1.ExternalMetric, err error) {
	result = &v1alpha1.ExternalMetric{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("externalmetrics").
		Body(externalMetric).
		Do(context.Background()).
		Into(result)
	return
}

// Update takes the representation of a externalMetric and updates it. Returns the server's representation of the externalMetric, and an error, if there is any.
func (c *externalMetrics) Update(externalMetric *v1alpha1.ExternalMetric) (result *v1alpha1.ExternalMetric, err error) {
	result = &v1alpha1.ExternalMetric{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("externalmetrics").
		Name(externalMetric.Name).
		Body(externalMetric).
		Do(context.Background()).
		Into(result)
	return
}

// Delete takes name of the externalMetric and deletes it. Returns an error if one occurs.
func (c *externalMetrics) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("externalmetrics").
		Name(name).
		Body(options).
		Do(context.Background()).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *externalMetrics) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("externalmetrics").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do(context.Background()).
		Error()
}
