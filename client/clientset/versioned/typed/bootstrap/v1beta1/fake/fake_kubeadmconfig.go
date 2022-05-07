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

	v1beta1 "github.com/kubesphere/kubeocean-api/apis/bootstrap/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeKubeadmConfigs implements KubeadmConfigInterface
type FakeKubeadmConfigs struct {
	Fake *FakeBootstrapV1beta1
	ns   string
}

var kubeadmconfigsResource = schema.GroupVersionResource{Group: "bootstrap", Version: "v1beta1", Resource: "kubeadmconfigs"}

var kubeadmconfigsKind = schema.GroupVersionKind{Group: "bootstrap", Version: "v1beta1", Kind: "KubeadmConfig"}

// Get takes name of the kubeadmConfig, and returns the corresponding kubeadmConfig object, and an error if there is any.
func (c *FakeKubeadmConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.KubeadmConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(kubeadmconfigsResource, c.ns, name), &v1beta1.KubeadmConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KubeadmConfig), err
}

// List takes label and field selectors, and returns the list of KubeadmConfigs that match those selectors.
func (c *FakeKubeadmConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.KubeadmConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(kubeadmconfigsResource, kubeadmconfigsKind, c.ns, opts), &v1beta1.KubeadmConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.KubeadmConfigList{ListMeta: obj.(*v1beta1.KubeadmConfigList).ListMeta}
	for _, item := range obj.(*v1beta1.KubeadmConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested kubeadmConfigs.
func (c *FakeKubeadmConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(kubeadmconfigsResource, c.ns, opts))

}

// Create takes the representation of a kubeadmConfig and creates it.  Returns the server's representation of the kubeadmConfig, and an error, if there is any.
func (c *FakeKubeadmConfigs) Create(ctx context.Context, kubeadmConfig *v1beta1.KubeadmConfig, opts v1.CreateOptions) (result *v1beta1.KubeadmConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(kubeadmconfigsResource, c.ns, kubeadmConfig), &v1beta1.KubeadmConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KubeadmConfig), err
}

// Update takes the representation of a kubeadmConfig and updates it. Returns the server's representation of the kubeadmConfig, and an error, if there is any.
func (c *FakeKubeadmConfigs) Update(ctx context.Context, kubeadmConfig *v1beta1.KubeadmConfig, opts v1.UpdateOptions) (result *v1beta1.KubeadmConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(kubeadmconfigsResource, c.ns, kubeadmConfig), &v1beta1.KubeadmConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KubeadmConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeKubeadmConfigs) UpdateStatus(ctx context.Context, kubeadmConfig *v1beta1.KubeadmConfig, opts v1.UpdateOptions) (*v1beta1.KubeadmConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(kubeadmconfigsResource, "status", c.ns, kubeadmConfig), &v1beta1.KubeadmConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KubeadmConfig), err
}

// Delete takes name of the kubeadmConfig and deletes it. Returns an error if one occurs.
func (c *FakeKubeadmConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(kubeadmconfigsResource, c.ns, name, opts), &v1beta1.KubeadmConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeKubeadmConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(kubeadmconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.KubeadmConfigList{})
	return err
}

// Patch applies the patch and returns the patched kubeadmConfig.
func (c *FakeKubeadmConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.KubeadmConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(kubeadmconfigsResource, c.ns, name, pt, data, subresources...), &v1beta1.KubeadmConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.KubeadmConfig), err
}
