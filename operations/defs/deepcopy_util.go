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

package defs

import "github.com/pkg/errors"

// deepCopy creates a deep copy of the json subtree pointed by the current node
// if deep is false only the subtree child references are copied else a recursive deep copy is made
func deepCopy(n Node, deep bool) (Node, error) {
	var node Node
	switch v := n.(type) {
	case *MapNode:
		m := NewMapNode(v.Id())
		node = m
	case *ListNode:
		l := NewListNode(v.Id())
		node = l
	case *RegisterNode:
		return NewRegisterNode(v.Id(), v.Value()), nil
	default:
		return nil, errors.Errorf("malicious entry of type %T inside json tree", n)
	}

	// if deep flag is not set do not copy recursively
	if !deep {
		for _id := range n.Child() {
			node.Child()[_id] = n.Child()[_id]
		}
		return node, nil
	}

	// recursive deepcopy
	for key, val := range n.Child() {
		deepVal, err := deepCopy(val, true)
		if err != nil {
			return nil, err
		}
		node.Child()[key] = deepVal
	}
	return node, nil
}
