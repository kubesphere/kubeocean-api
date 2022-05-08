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

	v1beta1 "github.com/kubesphere/kubeocean-api/apis/clusterapiproviderqc/v1beta1"
	scheme "github.com/kubesphere/kubeocean-api/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// QCClustersGetter has a method to return a QCClusterInterface.
// A group's client should implement this interface.
type QCClustersGetter interface {
	QCClusters(namespace string) QCClusterInterface
}

// QCClusterInterface has methods to work with QCCluster resources.
type QCClusterInterface interface {
	Create(ctx context.Context, qCCluster *v1beta1.QCCluster, opts v1.CreateOptions) (*v1beta1.QCCluster, error)
	Update(ctx context.Context, qCCluster *v1beta1.QCCluster, opts v1.UpdateOptions) (*v1beta1.QCCluster, error)
	UpdateStatus(ctx context.Context, qCCluster *v1beta1.QCCluster, opts v1.UpdateOptions) (*v1beta1.QCCluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.QCCluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.QCClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.QCCluster, err error)
	QCClusterExpansion
}

// qCClusters implements QCClusterInterface
type qCClusters struct {
	client rest.Interface
	ns     string
}

// newQCClusters returns a QCClusters
func newQCClusters(c *ClusterapiproviderqcV1beta1Client, namespace string) *qCClusters {
	return &qCClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the qCCluster, and returns the corresponding qCCluster object, and an error if there is any.
func (c *qCClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.QCCluster, err error) {
	result = &v1beta1.QCCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("qcclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of QCClusters that match those selectors.
func (c *qCClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.QCClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.QCClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("qcclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested qCClusters.
func (c *qCClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("qcclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a qCCluster and creates it.  Returns the server's representation of the qCCluster, and an error, if there is any.
func (c *qCClusters) Create(ctx context.Context, qCCluster *v1beta1.QCCluster, opts v1.CreateOptions) (result *v1beta1.QCCluster, err error) {
	result = &v1beta1.QCCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("qcclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(qCCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a qCCluster and updates it. Returns the server's representation of the qCCluster, and an error, if there is any.
func (c *qCClusters) Update(ctx context.Context, qCCluster *v1beta1.QCCluster, opts v1.UpdateOptions) (result *v1beta1.QCCluster, err error) {
	result = &v1beta1.QCCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("qcclusters").
		Name(qCCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(qCCluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *qCClusters) UpdateStatus(ctx context.Context, qCCluster *v1beta1.QCCluster, opts v1.UpdateOptions) (result *v1beta1.QCCluster, err error) {
	result = &v1beta1.QCCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("qcclusters").
		Name(qCCluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(qCCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the qCCluster and deletes it. Returns an error if one occurs.
func (c *qCClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("qcclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *qCClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("qcclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched qCCluster.
func (c *qCClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.QCCluster, err error) {
	result = &v1beta1.QCCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("qcclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}