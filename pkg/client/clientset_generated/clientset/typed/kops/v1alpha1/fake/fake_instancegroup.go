/*
Copyright 2018 The Kubernetes Authors.

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

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1alpha1 "k8s.io/kops/pkg/apis/kops/v1alpha1"
)

// FakeInstanceGroups implements InstanceGroupInterface
type FakeInstanceGroups struct {
	Fake *FakeKopsV1alpha1
	ns   string
}

var instancegroupsResource = schema.GroupVersionResource{Group: "kops", Version: "v1alpha1", Resource: "instancegroups"}

var instancegroupsKind = schema.GroupVersionKind{Group: "kops", Version: "v1alpha1", Kind: "InstanceGroup"}

// Get takes name of the instanceGroup, and returns the corresponding instanceGroup object, and an error if there is any.
func (c *FakeInstanceGroups) Get(name string, options v1.GetOptions) (result *v1alpha1.InstanceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(instancegroupsResource, c.ns, name), &v1alpha1.InstanceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceGroup), err
}

// List takes label and field selectors, and returns the list of InstanceGroups that match those selectors.
func (c *FakeInstanceGroups) List(opts v1.ListOptions) (result *v1alpha1.InstanceGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(instancegroupsResource, instancegroupsKind, c.ns, opts), &v1alpha1.InstanceGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.InstanceGroupList{}
	for _, item := range obj.(*v1alpha1.InstanceGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested instanceGroups.
func (c *FakeInstanceGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(instancegroupsResource, c.ns, opts))

}

// Create takes the representation of a instanceGroup and creates it.  Returns the server's representation of the instanceGroup, and an error, if there is any.
func (c *FakeInstanceGroups) Create(instanceGroup *v1alpha1.InstanceGroup) (result *v1alpha1.InstanceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(instancegroupsResource, c.ns, instanceGroup), &v1alpha1.InstanceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceGroup), err
}

// Update takes the representation of a instanceGroup and updates it. Returns the server's representation of the instanceGroup, and an error, if there is any.
func (c *FakeInstanceGroups) Update(instanceGroup *v1alpha1.InstanceGroup) (result *v1alpha1.InstanceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(instancegroupsResource, c.ns, instanceGroup), &v1alpha1.InstanceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceGroup), err
}

// Delete takes name of the instanceGroup and deletes it. Returns an error if one occurs.
func (c *FakeInstanceGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(instancegroupsResource, c.ns, name), &v1alpha1.InstanceGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeInstanceGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(instancegroupsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.InstanceGroupList{})
	return err
}

// Patch applies the patch and returns the patched instanceGroup.
func (c *FakeInstanceGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.InstanceGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(instancegroupsResource, c.ns, name, data, subresources...), &v1alpha1.InstanceGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.InstanceGroup), err
}
