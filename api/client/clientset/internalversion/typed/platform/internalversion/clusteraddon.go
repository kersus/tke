/*
 * Copyright 2019 THL A29 Limited, a Tencent company.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package internalversion

import (
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	platform "tkestack.io/tke/api/platform"
)

// ClusterAddonsGetter has a method to return a ClusterAddonInterface.
// A group's client should implement this interface.
type ClusterAddonsGetter interface {
	ClusterAddons() ClusterAddonInterface
}

// ClusterAddonInterface has methods to work with ClusterAddon resources.
type ClusterAddonInterface interface {
	Get(name string, options v1.GetOptions) (*platform.ClusterAddon, error)
	List(opts v1.ListOptions) (*platform.ClusterAddonList, error)
	ClusterAddonExpansion
}

// clusterAddons implements ClusterAddonInterface
type clusterAddons struct {
	client rest.Interface
}

// newClusterAddons returns a ClusterAddons
func newClusterAddons(c *PlatformClient) *clusterAddons {
	return &clusterAddons{
		client: c.RESTClient(),
	}
}

// Get takes name of the clusterAddon, and returns the corresponding clusterAddon object, and an error if there is any.
func (c *clusterAddons) Get(name string, options v1.GetOptions) (result *platform.ClusterAddon, err error) {
	result = &platform.ClusterAddon{}
	err = c.client.Get().
		Resource("clusteraddons").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterAddons that match those selectors.
func (c *clusterAddons) List(opts v1.ListOptions) (result *platform.ClusterAddonList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &platform.ClusterAddonList{}
	err = c.client.Get().
		Resource("clusteraddons").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}
