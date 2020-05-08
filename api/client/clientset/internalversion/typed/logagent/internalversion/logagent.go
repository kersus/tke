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

package internalversion

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	scheme "tkestack.io/tke/api/client/clientset/internalversion/scheme"
	logagent "tkestack.io/tke/api/logagent"
)

// LogAgentsGetter has a method to return a LogAgentInterface.
// A group's client should implement this interface.
type LogAgentsGetter interface {
	LogAgents() LogAgentInterface
}

// LogAgentInterface has methods to work with LogAgent resources.
type LogAgentInterface interface {
	Create(ctx context.Context, logAgent *logagent.LogAgent, opts v1.CreateOptions) (*logagent.LogAgent, error)
	Update(ctx context.Context, logAgent *logagent.LogAgent, opts v1.UpdateOptions) (*logagent.LogAgent, error)
	UpdateStatus(ctx context.Context, logAgent *logagent.LogAgent, opts v1.UpdateOptions) (*logagent.LogAgent, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*logagent.LogAgent, error)
	List(ctx context.Context, opts v1.ListOptions) (*logagent.LogAgentList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *logagent.LogAgent, err error)
	LogAgentExpansion
}

// logAgents implements LogAgentInterface
type logAgents struct {
	client rest.Interface
}

// newLogAgents returns a LogAgents
func newLogAgents(c *LogagentClient) *logAgents {
	return &logAgents{
		client: c.RESTClient(),
	}
}

// Get takes name of the logAgent, and returns the corresponding logAgent object, and an error if there is any.
func (c *logAgents) Get(ctx context.Context, name string, options v1.GetOptions) (result *logagent.LogAgent, err error) {
	result = &logagent.LogAgent{}
	err = c.client.Get().
		Resource("logagents").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of LogAgents that match those selectors.
func (c *logAgents) List(ctx context.Context, opts v1.ListOptions) (result *logagent.LogAgentList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &logagent.LogAgentList{}
	err = c.client.Get().
		Resource("logagents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested logAgents.
func (c *logAgents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("logagents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a logAgent and creates it.  Returns the server's representation of the logAgent, and an error, if there is any.
func (c *logAgents) Create(ctx context.Context, logAgent *logagent.LogAgent, opts v1.CreateOptions) (result *logagent.LogAgent, err error) {
	result = &logagent.LogAgent{}
	err = c.client.Post().
		Resource("logagents").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(logAgent).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a logAgent and updates it. Returns the server's representation of the logAgent, and an error, if there is any.
func (c *logAgents) Update(ctx context.Context, logAgent *logagent.LogAgent, opts v1.UpdateOptions) (result *logagent.LogAgent, err error) {
	result = &logagent.LogAgent{}
	err = c.client.Put().
		Resource("logagents").
		Name(logAgent.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(logAgent).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *logAgents) UpdateStatus(ctx context.Context, logAgent *logagent.LogAgent, opts v1.UpdateOptions) (result *logagent.LogAgent, err error) {
	result = &logagent.LogAgent{}
	err = c.client.Put().
		Resource("logagents").
		Name(logAgent.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(logAgent).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the logAgent and deletes it. Returns an error if one occurs.
func (c *logAgents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("logagents").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched logAgent.
func (c *logAgents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *logagent.LogAgent, err error) {
	result = &logagent.LogAgent{}
	err = c.client.Patch(pt).
		Resource("logagents").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
