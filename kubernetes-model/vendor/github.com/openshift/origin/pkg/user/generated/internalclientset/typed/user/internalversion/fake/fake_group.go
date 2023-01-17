/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package fake

import (
	user "github.com/openshift/origin/pkg/user/apis/user"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGroups implements GroupInterface
type FakeGroups struct {
	Fake *FakeUser
}

var groupsResource = schema.GroupVersionResource{Group: "user.openshift.io", Version: "", Resource: "groups"}

var groupsKind = schema.GroupVersionKind{Group: "user.openshift.io", Version: "", Kind: "Group"}

// Get takes name of the group, and returns the corresponding group object, and an error if there is any.
func (c *FakeGroups) Get(name string, options v1.GetOptions) (result *user.Group, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(groupsResource, name), &user.Group{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user.Group), err
}

// List takes label and field selectors, and returns the list of Groups that match those selectors.
func (c *FakeGroups) List(opts v1.ListOptions) (result *user.GroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(groupsResource, groupsKind, opts), &user.GroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &user.GroupList{}
	for _, item := range obj.(*user.GroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested groups.
func (c *FakeGroups) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(groupsResource, opts))
}

// Create takes the representation of a group and creates it.  Returns the server's representation of the group, and an error, if there is any.
func (c *FakeGroups) Create(group *user.Group) (result *user.Group, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(groupsResource, group), &user.Group{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user.Group), err
}

// Update takes the representation of a group and updates it. Returns the server's representation of the group, and an error, if there is any.
func (c *FakeGroups) Update(group *user.Group) (result *user.Group, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(groupsResource, group), &user.Group{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user.Group), err
}

// Delete takes name of the group and deletes it. Returns an error if one occurs.
func (c *FakeGroups) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(groupsResource, name), &user.Group{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGroups) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(groupsResource, listOptions)

	_, err := c.Fake.Invokes(action, &user.GroupList{})
	return err
}

// Patch applies the patch and returns the patched group.
func (c *FakeGroups) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *user.Group, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(groupsResource, name, data, subresources...), &user.Group{})
	if obj == nil {
		return nil, err
	}
	return obj.(*user.Group), err
}
