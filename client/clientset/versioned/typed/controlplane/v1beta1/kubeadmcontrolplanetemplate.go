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

	v1beta1 "github.com/kubesphere/kubeocean-api/apis/controlplane/v1beta1"
	scheme "github.com/kubesphere/kubeocean-api/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// KubeadmControlPlaneTemplatesGetter has a method to return a KubeadmControlPlaneTemplateInterface.
// A group's client should implement this interface.
type KubeadmControlPlaneTemplatesGetter interface {
	KubeadmControlPlaneTemplates(namespace string) KubeadmControlPlaneTemplateInterface
}

// KubeadmControlPlaneTemplateInterface has methods to work with KubeadmControlPlaneTemplate resources.
type KubeadmControlPlaneTemplateInterface interface {
	Create(ctx context.Context, kubeadmControlPlaneTemplate *v1beta1.KubeadmControlPlaneTemplate, opts v1.CreateOptions) (*v1beta1.KubeadmControlPlaneTemplate, error)
	Update(ctx context.Context, kubeadmControlPlaneTemplate *v1beta1.KubeadmControlPlaneTemplate, opts v1.UpdateOptions) (*v1beta1.KubeadmControlPlaneTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.KubeadmControlPlaneTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.KubeadmControlPlaneTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KubeadmControlPlaneTemplate, err error)
	KubeadmControlPlaneTemplateExpansion
}

// kubeadmControlPlaneTemplates implements KubeadmControlPlaneTemplateInterface
type kubeadmControlPlaneTemplates struct {
	client rest.Interface
	ns     string
}

// newKubeadmControlPlaneTemplates returns a KubeadmControlPlaneTemplates
func newKubeadmControlPlaneTemplates(c *ControlplaneV1beta1Client, namespace string) *kubeadmControlPlaneTemplates {
	return &kubeadmControlPlaneTemplates{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kubeadmControlPlaneTemplate, and returns the corresponding kubeadmControlPlaneTemplate object, and an error if there is any.
func (c *kubeadmControlPlaneTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.KubeadmControlPlaneTemplate, err error) {
	result = &v1beta1.KubeadmControlPlaneTemplate{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubeadmcontrolplanetemplates").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KubeadmControlPlaneTemplates that match those selectors.
func (c *kubeadmControlPlaneTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.KubeadmControlPlaneTemplateList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.KubeadmControlPlaneTemplateList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kubeadmcontrolplanetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kubeadmControlPlaneTemplates.
func (c *kubeadmControlPlaneTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kubeadmcontrolplanetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kubeadmControlPlaneTemplate and creates it.  Returns the server's representation of the kubeadmControlPlaneTemplate, and an error, if there is any.
func (c *kubeadmControlPlaneTemplates) Create(ctx context.Context, kubeadmControlPlaneTemplate *v1beta1.KubeadmControlPlaneTemplate, opts v1.CreateOptions) (result *v1beta1.KubeadmControlPlaneTemplate, err error) {
	result = &v1beta1.KubeadmControlPlaneTemplate{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kubeadmcontrolplanetemplates").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeadmControlPlaneTemplate).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kubeadmControlPlaneTemplate and updates it. Returns the server's representation of the kubeadmControlPlaneTemplate, and an error, if there is any.
func (c *kubeadmControlPlaneTemplates) Update(ctx context.Context, kubeadmControlPlaneTemplate *v1beta1.KubeadmControlPlaneTemplate, opts v1.UpdateOptions) (result *v1beta1.KubeadmControlPlaneTemplate, err error) {
	result = &v1beta1.KubeadmControlPlaneTemplate{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kubeadmcontrolplanetemplates").
		Name(kubeadmControlPlaneTemplate.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kubeadmControlPlaneTemplate).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kubeadmControlPlaneTemplate and deletes it. Returns an error if one occurs.
func (c *kubeadmControlPlaneTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubeadmcontrolplanetemplates").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kubeadmControlPlaneTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kubeadmcontrolplanetemplates").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kubeadmControlPlaneTemplate.
func (c *kubeadmControlPlaneTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KubeadmControlPlaneTemplate, err error) {
	result = &v1beta1.KubeadmControlPlaneTemplate{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kubeadmcontrolplanetemplates").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
