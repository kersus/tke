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

package drain

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/client-go/kubernetes"
)

// CordonHelper wraps functionality to cordon/uncordon nodes
type CordonHelper struct {
	node    *corev1.Node
	desired bool
}

// NewCordonHelper returns a new CordonHelper
func NewCordonHelper(node *corev1.Node) *CordonHelper {
	return &CordonHelper{
		node: node,
	}
}

// NewCordonHelperFromRuntimeObject returns a new CordonHelper, or an error if given object is not a
// node or cannot be encoded as JSON
func NewCordonHelperFromRuntimeObject(nodeObject runtime.Object, scheme *runtime.Scheme, gvk schema.GroupVersionKind) (*CordonHelper, error) {
	nodeObject, err := scheme.ConvertToVersion(nodeObject, gvk.GroupVersion())
	if err != nil {
		return nil, err
	}

	node, ok := nodeObject.(*corev1.Node)
	if !ok {
		return nil, fmt.Errorf("unexpected type %T", nodeObject)
	}

	return NewCordonHelper(node), nil
}

// UpdateIfRequired returns true if c.node.Spec.Unschedulable isn't already set,
// or false when no change is needed
func (c *CordonHelper) UpdateIfRequired(desired bool) bool {
	c.desired = desired
	if c.node.Spec.Unschedulable == c.desired {
		return false
	}
	return true
}

// PatchOrReplace uses given clientset to update the node status, either by patching or
// updating the given node object; it may return error if the object cannot be encoded as
// JSON, or if either patch or update calls fail; it will also return a second error
// whenever creating a patch has failed
func (c *CordonHelper) PatchOrReplace(clientset kubernetes.Interface) (err error, patchErr error) {
	client := clientset.CoreV1().Nodes()

	oldData, err := json.Marshal(c.node)
	if err != nil {
		return err, nil
	}

	c.node.Spec.Unschedulable = c.desired

	newData, err := json.Marshal(c.node)
	if err != nil {
		return err, nil
	}

	patchBytes, patchErr := strategicpatch.CreateTwoWayMergePatch(oldData, newData, c.node)
	if patchErr == nil {
		_, err = client.Patch(context.Background(), c.node.Name, types.StrategicMergePatchType, patchBytes, metav1.PatchOptions{})
	} else {
		_, err = client.Update(context.Background(), c.node, metav1.UpdateOptions{})
	}
	return err, patchErr
}
