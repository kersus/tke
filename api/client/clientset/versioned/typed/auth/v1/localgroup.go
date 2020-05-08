/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
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

package v1

import (
	"context"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1 "tkestack.io/tke/api/auth/v1"
	scheme "tkestack.io/tke/api/client/clientset/versioned/scheme"
)

// LocalGroupsGetter has a method to return a LocalGroupInterface.
// A group's client should implement this interface.
type LocalGroupsGetter interface {
	LocalGroups() LocalGroupInterface
}

// LocalGroupInterface has methods to work with LocalGroup resources.
type LocalGroupInterface interface {
	Create(ctx context.Context, localGroup *v1.LocalGroup, opts metav1.CreateOptions) (*v1.LocalGroup, error)
	Update(ctx context.Context, localGroup *v1.LocalGroup, opts metav1.UpdateOptions) (*v1.LocalGroup, error)
	UpdateStatus(ctx context.Context, localGroup *v1.LocalGroup, opts metav1.UpdateOptions) (*v1.LocalGroup, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.LocalGroup, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.LocalGroupList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LocalGroup, err error)
	LocalGroupExpansion
}

// localGroups implements LocalGroupInterface
type localGroups struct {
	client rest.Interface
}

// newLocalGroups returns a LocalGroups
func newLocalGroups(c *AuthV1Client) *localGroups {
	return &localGroups{
		client: c.RESTClient(),
	}
}

// Get takes name of the localGroup, and returns the corresponding localGroup object, and an error if there is any.
func (c *localGroups) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.LocalGroup, err error) {
	result = &v1.LocalGroup{}
	err = c.client.Get().
		Resource("localgroups").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LocalGroups that match those selectors.
func (c *localGroups) List(ctx context.Context, opts metav1.ListOptions) (result *v1.LocalGroupList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.LocalGroupList{}
	err = c.client.Get().
		Resource("localgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested localGroups.
func (c *localGroups) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("localgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a localGroup and creates it.  Returns the server's representation of the localGroup, and an error, if there is any.
func (c *localGroups) Create(ctx context.Context, localGroup *v1.LocalGroup, opts metav1.CreateOptions) (result *v1.LocalGroup, err error) {
	result = &v1.LocalGroup{}
	err = c.client.Post().
		Resource("localgroups").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localGroup).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a localGroup and updates it. Returns the server's representation of the localGroup, and an error, if there is any.
func (c *localGroups) Update(ctx context.Context, localGroup *v1.LocalGroup, opts metav1.UpdateOptions) (result *v1.LocalGroup, err error) {
	result = &v1.LocalGroup{}
	err = c.client.Put().
		Resource("localgroups").
		Name(localGroup.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localGroup).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *localGroups) UpdateStatus(ctx context.Context, localGroup *v1.LocalGroup, opts metav1.UpdateOptions) (result *v1.LocalGroup, err error) {
	result = &v1.LocalGroup{}
	err = c.client.Put().
		Resource("localgroups").
		Name(localGroup.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(localGroup).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the localGroup and deletes it. Returns an error if one occurs.
func (c *localGroups) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("localgroups").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *localGroups) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("localgroups").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched localGroup.
func (c *localGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.LocalGroup, err error) {
	result = &v1.LocalGroup{}
	err = c.client.Patch(pt).
		Resource("localgroups").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
