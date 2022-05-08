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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/kubesphere/kubeocean-api/apis/kubeocean/v1alpha1"
	scheme "github.com/kubesphere/kubeocean-api/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// DolphinClusterSpecificationsGetter has a method to return a DolphinClusterSpecificationInterface.
// A group's client should implement this interface.
type DolphinClusterSpecificationsGetter interface {
	DolphinClusterSpecifications() DolphinClusterSpecificationInterface
}

// DolphinClusterSpecificationInterface has methods to work with DolphinClusterSpecification resources.
type DolphinClusterSpecificationInterface interface {
	Create(ctx context.Context, dolphinClusterSpecification *v1alpha1.DolphinClusterSpecification, opts v1.CreateOptions) (*v1alpha1.DolphinClusterSpecification, error)
	Update(ctx context.Context, dolphinClusterSpecification *v1alpha1.DolphinClusterSpecification, opts v1.UpdateOptions) (*v1alpha1.DolphinClusterSpecification, error)
	UpdateStatus(ctx context.Context, dolphinClusterSpecification *v1alpha1.DolphinClusterSpecification, opts v1.UpdateOptions) (*v1alpha1.DolphinClusterSpecification, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.DolphinClusterSpecification, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.DolphinClusterSpecificationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DolphinClusterSpecification, err error)
	DolphinClusterSpecificationExpansion
}

// dolphinClusterSpecifications implements DolphinClusterSpecificationInterface
type dolphinClusterSpecifications struct {
	client rest.Interface
}

// newDolphinClusterSpecifications returns a DolphinClusterSpecifications
func newDolphinClusterSpecifications(c *KubeoceanV1alpha1Client) *dolphinClusterSpecifications {
	return &dolphinClusterSpecifications{
		client: c.RESTClient(),
	}
}

// Get takes name of the dolphinClusterSpecification, and returns the corresponding dolphinClusterSpecification object, and an error if there is any.
func (c *dolphinClusterSpecifications) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DolphinClusterSpecification, err error) {
	result = &v1alpha1.DolphinClusterSpecification{}
	err = c.client.Get().
		Resource("dolphinclusterspecifications").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of DolphinClusterSpecifications that match those selectors.
func (c *dolphinClusterSpecifications) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DolphinClusterSpecificationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.DolphinClusterSpecificationList{}
	err = c.client.Get().
		Resource("dolphinclusterspecifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested dolphinClusterSpecifications.
func (c *dolphinClusterSpecifications) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("dolphinclusterspecifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a dolphinClusterSpecification and creates it.  Returns the server's representation of the dolphinClusterSpecification, and an error, if there is any.
func (c *dolphinClusterSpecifications) Create(ctx context.Context, dolphinClusterSpecification *v1alpha1.DolphinClusterSpecification, opts v1.CreateOptions) (result *v1alpha1.DolphinClusterSpecification, err error) {
	result = &v1alpha1.DolphinClusterSpecification{}
	err = c.client.Post().
		Resource("dolphinclusterspecifications").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dolphinClusterSpecification).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a dolphinClusterSpecification and updates it. Returns the server's representation of the dolphinClusterSpecification, and an error, if there is any.
func (c *dolphinClusterSpecifications) Update(ctx context.Context, dolphinClusterSpecification *v1alpha1.DolphinClusterSpecification, opts v1.UpdateOptions) (result *v1alpha1.DolphinClusterSpecification, err error) {
	result = &v1alpha1.DolphinClusterSpecification{}
	err = c.client.Put().
		Resource("dolphinclusterspecifications").
		Name(dolphinClusterSpecification.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dolphinClusterSpecification).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *dolphinClusterSpecifications) UpdateStatus(ctx context.Context, dolphinClusterSpecification *v1alpha1.DolphinClusterSpecification, opts v1.UpdateOptions) (result *v1alpha1.DolphinClusterSpecification, err error) {
	result = &v1alpha1.DolphinClusterSpecification{}
	err = c.client.Put().
		Resource("dolphinclusterspecifications").
		Name(dolphinClusterSpecification.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(dolphinClusterSpecification).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the dolphinClusterSpecification and deletes it. Returns an error if one occurs.
func (c *dolphinClusterSpecifications) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("dolphinclusterspecifications").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *dolphinClusterSpecifications) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("dolphinclusterspecifications").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched dolphinClusterSpecification.
func (c *dolphinClusterSpecifications) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DolphinClusterSpecification, err error) {
	result = &v1alpha1.DolphinClusterSpecification{}
	err = c.client.Patch(pt).
		Resource("dolphinclusterspecifications").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}