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

// ClusterClassesGetter has a method to return a ClusterClassInterface.
// A group's client should implement this interface.
type ClusterClassesGetter interface {
	ClusterClasses(namespace string) ClusterClassInterface
}

// ClusterClassInterface has methods to work with ClusterClass resources.
type ClusterClassInterface interface {
	Create(ctx context.Context, clusterClass *v1beta1.ClusterClass, opts v1.CreateOptions) (*v1beta1.ClusterClass, error)
	Update(ctx context.Context, clusterClass *v1beta1.ClusterClass, opts v1.UpdateOptions) (*v1beta1.ClusterClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ClusterClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ClusterClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ClusterClass, err error)
	ClusterClassExpansion
}

// clusterClasses implements ClusterClassInterface
type clusterClasses struct {
	client rest.Interface
	ns     string
}

// newClusterClasses returns a ClusterClasses
func newClusterClasses(c *ClusterapiV1beta1Client, namespace string) *clusterClasses {
	return &clusterClasses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterClass, and returns the corresponding clusterClass object, and an error if there is any.
func (c *clusterClasses) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ClusterClass, err error) {
	result = &v1beta1.ClusterClass{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterClasses that match those selectors.
func (c *clusterClasses) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ClusterClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ClusterClassList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterClasses.
func (c *clusterClasses) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a clusterClass and creates it.  Returns the server's representation of the clusterClass, and an error, if there is any.
func (c *clusterClasses) Create(ctx context.Context, clusterClass *v1beta1.ClusterClass, opts v1.CreateOptions) (result *v1beta1.ClusterClass, err error) {
	result = &v1beta1.ClusterClass{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterClass).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a clusterClass and updates it. Returns the server's representation of the clusterClass, and an error, if there is any.
func (c *clusterClasses) Update(ctx context.Context, clusterClass *v1beta1.ClusterClass, opts v1.UpdateOptions) (result *v1beta1.ClusterClass, err error) {
	result = &v1beta1.ClusterClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterclasses").
		Name(clusterClass.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(clusterClass).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the clusterClass and deletes it. Returns an error if one occurs.
func (c *clusterClasses) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterclasses").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterClasses) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterclasses").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched clusterClass.
func (c *clusterClasses) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ClusterClass, err error) {
	result = &v1beta1.ClusterClass{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterclasses").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
