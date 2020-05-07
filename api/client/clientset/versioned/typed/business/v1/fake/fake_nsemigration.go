/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	businessv1 "tkestack.io/tke/api/business/v1"
)

// FakeNsEmigrations implements NsEmigrationInterface
type FakeNsEmigrations struct {
	Fake *FakeBusinessV1
	ns   string
}

var nsemigrationsResource = schema.GroupVersionResource{Group: "business.tkestack.io", Version: "v1", Resource: "nsemigrations"}

var nsemigrationsKind = schema.GroupVersionKind{Group: "business.tkestack.io", Version: "v1", Kind: "NsEmigration"}

// Get takes name of the nsEmigration, and returns the corresponding nsEmigration object, and an error if there is any.
func (c *FakeNsEmigrations) Get(name string, options v1.GetOptions) (result *businessv1.NsEmigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(nsemigrationsResource, c.ns, name), &businessv1.NsEmigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*businessv1.NsEmigration), err
}

// List takes label and field selectors, and returns the list of NsEmigrations that match those selectors.
func (c *FakeNsEmigrations) List(opts v1.ListOptions) (result *businessv1.NsEmigrationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(nsemigrationsResource, nsemigrationsKind, c.ns, opts), &businessv1.NsEmigrationList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &businessv1.NsEmigrationList{ListMeta: obj.(*businessv1.NsEmigrationList).ListMeta}
	for _, item := range obj.(*businessv1.NsEmigrationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nsEmigrations.
func (c *FakeNsEmigrations) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(nsemigrationsResource, c.ns, opts))

}

// Create takes the representation of a nsEmigration and creates it.  Returns the server's representation of the nsEmigration, and an error, if there is any.
func (c *FakeNsEmigrations) Create(nsEmigration *businessv1.NsEmigration) (result *businessv1.NsEmigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(nsemigrationsResource, c.ns, nsEmigration), &businessv1.NsEmigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*businessv1.NsEmigration), err
}

// Update takes the representation of a nsEmigration and updates it. Returns the server's representation of the nsEmigration, and an error, if there is any.
func (c *FakeNsEmigrations) Update(nsEmigration *businessv1.NsEmigration) (result *businessv1.NsEmigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(nsemigrationsResource, c.ns, nsEmigration), &businessv1.NsEmigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*businessv1.NsEmigration), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNsEmigrations) UpdateStatus(nsEmigration *businessv1.NsEmigration) (*businessv1.NsEmigration, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(nsemigrationsResource, "status", c.ns, nsEmigration), &businessv1.NsEmigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*businessv1.NsEmigration), err
}

// Delete takes name of the nsEmigration and deletes it. Returns an error if one occurs.
func (c *FakeNsEmigrations) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(nsemigrationsResource, c.ns, name), &businessv1.NsEmigration{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNsEmigrations) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(nsemigrationsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &businessv1.NsEmigrationList{})
	return err
}

// Patch applies the patch and returns the patched nsEmigration.
func (c *FakeNsEmigrations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *businessv1.NsEmigration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(nsemigrationsResource, c.ns, name, pt, data, subresources...), &businessv1.NsEmigration{})

	if obj == nil {
		return nil, err
	}
	return obj.(*businessv1.NsEmigration), err
}
