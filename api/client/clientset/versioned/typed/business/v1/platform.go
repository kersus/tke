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
	v1 "tkestack.io/tke/api/business/v1"
	scheme "tkestack.io/tke/api/client/clientset/versioned/scheme"
)

// PlatformsGetter has a method to return a PlatformInterface.
// A group's client should implement this interface.
type PlatformsGetter interface {
	Platforms() PlatformInterface
}

// PlatformInterface has methods to work with Platform resources.
type PlatformInterface interface {
	Create(ctx context.Context, platform *v1.Platform, opts metav1.CreateOptions) (*v1.Platform, error)
	Update(ctx context.Context, platform *v1.Platform, opts metav1.UpdateOptions) (*v1.Platform, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.Platform, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.PlatformList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Platform, err error)
	PlatformExpansion
}

// platforms implements PlatformInterface
type platforms struct {
	client rest.Interface
}

// newPlatforms returns a Platforms
func newPlatforms(c *BusinessV1Client) *platforms {
	return &platforms{
		client: c.RESTClient(),
	}
}

// Get takes name of the platform, and returns the corresponding platform object, and an error if there is any.
func (c *platforms) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Platform, err error) {
	result = &v1.Platform{}
	err = c.client.Get().
		Resource("platforms").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Platforms that match those selectors.
func (c *platforms) List(ctx context.Context, opts metav1.ListOptions) (result *v1.PlatformList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.PlatformList{}
	err = c.client.Get().
		Resource("platforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested platforms.
func (c *platforms) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("platforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a platform and creates it.  Returns the server's representation of the platform, and an error, if there is any.
func (c *platforms) Create(ctx context.Context, platform *v1.Platform, opts metav1.CreateOptions) (result *v1.Platform, err error) {
	result = &v1.Platform{}
	err = c.client.Post().
		Resource("platforms").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(platform).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a platform and updates it. Returns the server's representation of the platform, and an error, if there is any.
func (c *platforms) Update(ctx context.Context, platform *v1.Platform, opts metav1.UpdateOptions) (result *v1.Platform, err error) {
	result = &v1.Platform{}
	err = c.client.Put().
		Resource("platforms").
		Name(platform.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(platform).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the platform and deletes it. Returns an error if one occurs.
func (c *platforms) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Resource("platforms").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched platform.
func (c *platforms) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Platform, err error) {
	result = &v1.Platform{}
	err = c.client.Patch(pt).
		Resource("platforms").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
