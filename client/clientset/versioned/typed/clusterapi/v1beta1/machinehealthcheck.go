/*
Copyright 2020 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/kubesphere/kubeocean-api/apis/clusterapi/v1beta1"
	scheme "github.com/kubesphere/kubeocean-api/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MachineHealthChecksGetter has a method to return a MachineHealthCheckInterface.
// A group's client should implement this interface.
type MachineHealthChecksGetter interface {
	MachineHealthChecks(namespace string) MachineHealthCheckInterface
}

// MachineHealthCheckInterface has methods to work with MachineHealthCheck resources.
type MachineHealthCheckInterface interface {
	Create(ctx context.Context, machineHealthCheck *v1beta1.MachineHealthCheck, opts v1.CreateOptions) (*v1beta1.MachineHealthCheck, error)
	Update(ctx context.Context, machineHealthCheck *v1beta1.MachineHealthCheck, opts v1.UpdateOptions) (*v1beta1.MachineHealthCheck, error)
	UpdateStatus(ctx context.Context, machineHealthCheck *v1beta1.MachineHealthCheck, opts v1.UpdateOptions) (*v1beta1.MachineHealthCheck, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.MachineHealthCheck, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.MachineHealthCheckList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MachineHealthCheck, err error)
	MachineHealthCheckExpansion
}

// machineHealthChecks implements MachineHealthCheckInterface
type machineHealthChecks struct {
	client rest.Interface
	ns     string
}

// newMachineHealthChecks returns a MachineHealthChecks
func newMachineHealthChecks(c *ClusterapiV1beta1Client, namespace string) *machineHealthChecks {
	return &machineHealthChecks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the machineHealthCheck, and returns the corresponding machineHealthCheck object, and an error if there is any.
func (c *machineHealthChecks) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MachineHealthCheck, err error) {
	result = &v1beta1.MachineHealthCheck{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machinehealthchecks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MachineHealthChecks that match those selectors.
func (c *machineHealthChecks) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MachineHealthCheckList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.MachineHealthCheckList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("machinehealthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested machineHealthChecks.
func (c *machineHealthChecks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("machinehealthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a machineHealthCheck and creates it.  Returns the server's representation of the machineHealthCheck, and an error, if there is any.
func (c *machineHealthChecks) Create(ctx context.Context, machineHealthCheck *v1beta1.MachineHealthCheck, opts v1.CreateOptions) (result *v1beta1.MachineHealthCheck, err error) {
	result = &v1beta1.MachineHealthCheck{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("machinehealthchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineHealthCheck).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a machineHealthCheck and updates it. Returns the server's representation of the machineHealthCheck, and an error, if there is any.
func (c *machineHealthChecks) Update(ctx context.Context, machineHealthCheck *v1beta1.MachineHealthCheck, opts v1.UpdateOptions) (result *v1beta1.MachineHealthCheck, err error) {
	result = &v1beta1.MachineHealthCheck{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machinehealthchecks").
		Name(machineHealthCheck.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineHealthCheck).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *machineHealthChecks) UpdateStatus(ctx context.Context, machineHealthCheck *v1beta1.MachineHealthCheck, opts v1.UpdateOptions) (result *v1beta1.MachineHealthCheck, err error) {
	result = &v1beta1.MachineHealthCheck{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("machinehealthchecks").
		Name(machineHealthCheck.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(machineHealthCheck).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the machineHealthCheck and deletes it. Returns an error if one occurs.
func (c *machineHealthChecks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machinehealthchecks").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *machineHealthChecks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("machinehealthchecks").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched machineHealthCheck.
func (c *machineHealthChecks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MachineHealthCheck, err error) {
	result = &v1beta1.MachineHealthCheck{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("machinehealthchecks").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
