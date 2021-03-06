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

	v1beta1 "github.com/kubesphere/kubeocean-api/apis/clusterapiproviderqc/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeQCMachineTemplates implements QCMachineTemplateInterface
type FakeQCMachineTemplates struct {
	Fake *FakeClusterapiproviderqcV1beta1
	ns   string
}

var qcmachinetemplatesResource = schema.GroupVersionResource{Group: "clusterapiproviderqc", Version: "v1beta1", Resource: "qcmachinetemplates"}

var qcmachinetemplatesKind = schema.GroupVersionKind{Group: "clusterapiproviderqc", Version: "v1beta1", Kind: "QCMachineTemplate"}

// Get takes name of the qCMachineTemplate, and returns the corresponding qCMachineTemplate object, and an error if there is any.
func (c *FakeQCMachineTemplates) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.QCMachineTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(qcmachinetemplatesResource, c.ns, name), &v1beta1.QCMachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.QCMachineTemplate), err
}

// List takes label and field selectors, and returns the list of QCMachineTemplates that match those selectors.
func (c *FakeQCMachineTemplates) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.QCMachineTemplateList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(qcmachinetemplatesResource, qcmachinetemplatesKind, c.ns, opts), &v1beta1.QCMachineTemplateList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.QCMachineTemplateList{ListMeta: obj.(*v1beta1.QCMachineTemplateList).ListMeta}
	for _, item := range obj.(*v1beta1.QCMachineTemplateList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested qCMachineTemplates.
func (c *FakeQCMachineTemplates) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(qcmachinetemplatesResource, c.ns, opts))

}

// Create takes the representation of a qCMachineTemplate and creates it.  Returns the server's representation of the qCMachineTemplate, and an error, if there is any.
func (c *FakeQCMachineTemplates) Create(ctx context.Context, qCMachineTemplate *v1beta1.QCMachineTemplate, opts v1.CreateOptions) (result *v1beta1.QCMachineTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(qcmachinetemplatesResource, c.ns, qCMachineTemplate), &v1beta1.QCMachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.QCMachineTemplate), err
}

// Update takes the representation of a qCMachineTemplate and updates it. Returns the server's representation of the qCMachineTemplate, and an error, if there is any.
func (c *FakeQCMachineTemplates) Update(ctx context.Context, qCMachineTemplate *v1beta1.QCMachineTemplate, opts v1.UpdateOptions) (result *v1beta1.QCMachineTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(qcmachinetemplatesResource, c.ns, qCMachineTemplate), &v1beta1.QCMachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.QCMachineTemplate), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeQCMachineTemplates) UpdateStatus(ctx context.Context, qCMachineTemplate *v1beta1.QCMachineTemplate, opts v1.UpdateOptions) (*v1beta1.QCMachineTemplate, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(qcmachinetemplatesResource, "status", c.ns, qCMachineTemplate), &v1beta1.QCMachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.QCMachineTemplate), err
}

// Delete takes name of the qCMachineTemplate and deletes it. Returns an error if one occurs.
func (c *FakeQCMachineTemplates) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(qcmachinetemplatesResource, c.ns, name, opts), &v1beta1.QCMachineTemplate{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeQCMachineTemplates) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(qcmachinetemplatesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.QCMachineTemplateList{})
	return err
}

// Patch applies the patch and returns the patched qCMachineTemplate.
func (c *FakeQCMachineTemplates) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.QCMachineTemplate, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(qcmachinetemplatesResource, c.ns, name, pt, data, subresources...), &v1beta1.QCMachineTemplate{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.QCMachineTemplate), err
}
