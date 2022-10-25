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

package fake

import (
	"context"

	v2alpha1 "github.com/kubesphere/kubeocean-api/v2/apis/kubeocean/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterPools implements ClusterPoolInterface
type FakeClusterPools struct {
	Fake *FakeKubeoceanV2alpha1
}

var clusterpoolsResource = schema.GroupVersionResource{Group: "kubeocean.kubesphere.io", Version: "v2alpha1", Resource: "clusterpools"}

var clusterpoolsKind = schema.GroupVersionKind{Group: "kubeocean.kubesphere.io", Version: "v2alpha1", Kind: "ClusterPool"}

// Get takes name of the clusterPool, and returns the corresponding clusterPool object, and an error if there is any.
func (c *FakeClusterPools) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.ClusterPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterpoolsResource, name), &v2alpha1.ClusterPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ClusterPool), err
}

// List takes label and field selectors, and returns the list of ClusterPools that match those selectors.
func (c *FakeClusterPools) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.ClusterPoolList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterpoolsResource, clusterpoolsKind, opts), &v2alpha1.ClusterPoolList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.ClusterPoolList{ListMeta: obj.(*v2alpha1.ClusterPoolList).ListMeta}
	for _, item := range obj.(*v2alpha1.ClusterPoolList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterPools.
func (c *FakeClusterPools) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterpoolsResource, opts))
}

// Create takes the representation of a clusterPool and creates it.  Returns the server's representation of the clusterPool, and an error, if there is any.
func (c *FakeClusterPools) Create(ctx context.Context, clusterPool *v2alpha1.ClusterPool, opts v1.CreateOptions) (result *v2alpha1.ClusterPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterpoolsResource, clusterPool), &v2alpha1.ClusterPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ClusterPool), err
}

// Update takes the representation of a clusterPool and updates it. Returns the server's representation of the clusterPool, and an error, if there is any.
func (c *FakeClusterPools) Update(ctx context.Context, clusterPool *v2alpha1.ClusterPool, opts v1.UpdateOptions) (result *v2alpha1.ClusterPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterpoolsResource, clusterPool), &v2alpha1.ClusterPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ClusterPool), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterPools) UpdateStatus(ctx context.Context, clusterPool *v2alpha1.ClusterPool, opts v1.UpdateOptions) (*v2alpha1.ClusterPool, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterpoolsResource, "status", clusterPool), &v2alpha1.ClusterPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ClusterPool), err
}

// Delete takes name of the clusterPool and deletes it. Returns an error if one occurs.
func (c *FakeClusterPools) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusterpoolsResource, name, opts), &v2alpha1.ClusterPool{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterPools) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterpoolsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.ClusterPoolList{})
	return err
}

// Patch applies the patch and returns the patched clusterPool.
func (c *FakeClusterPools) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.ClusterPool, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterpoolsResource, name, pt, data, subresources...), &v2alpha1.ClusterPool{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.ClusterPool), err
}
