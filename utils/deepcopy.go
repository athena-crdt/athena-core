//  Copyright 2021, athena-crdt authors.
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package utils

import (
	"github.com/pkg/errors"

	"github.com/athena-crdt/athena-core/operations/defs"
)

// DeepCopy creates a deep copy of the json subtree pointed by the current node
func DeepCopy(n defs.Node) (defs.Node, error) {
	var node defs.Node
	switch v := n.(type) {
	case *defs.MapNode:
		m := defs.NewMapNode(v.Id())
		node = m
	case *defs.ListNode:
		l := defs.NewListNode(v.Id())
		l.SetIndex(v.Index())
		node = l
	case *defs.RegisterNode:
		return defs.NewRegisterNode(v.Id(), v.Value()), nil
	default:
		return nil, errors.Errorf("malicious entry of type %T inside json tree", n)
	}

	// recursive deepcopy
	for key, val := range n.Child() {
		deepVal, err := DeepCopy(val)
		if err != nil {
			return nil, err
		}
		node.Child()[key] = deepVal
	}
	return node, nil
}
