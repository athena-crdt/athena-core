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

type RegisterNode struct {
	*baseNode
	value interface{}
}

func (r *RegisterNode) Clone() (Node, error) {
	return NewRegisterNode(r.id, r.value), nil
}

// SetValue assigns the current node with the given value.
func (r *RegisterNode) SetValue(val interface{}) {
	r.value = val
}

// Value returns the current RegisterNode value.
func (r *RegisterNode) Value() interface{} {
	return r.value
}

// FetchChild override. RegisterNode is always a leaf node.
func (r *RegisterNode) FetchChild([]ID) (Node, error) {
	return nil, errors.New("RegisterNode doesn't have a children set")
}

// Child returns a nil object, as RegisterNode itself is a leaf node.
func (r *RegisterNode) Child() Children {
	return nil
}

// NewRegisterNode returns a Node of type RegisterNode with given id and value.
func NewRegisterNode(id ID, value interface{}) *RegisterNode {
	return &RegisterNode{
		baseNode: &baseNode{
			id:        id,
			tombstone: false,
			children:  nil,
		},
		value: value,
	}
}

func (r *RegisterNode) Serialize() ([]byte, error) {
	panic("implement me")
}

func (r *RegisterNode) Deserialize(bytes []byte) error {
	panic("implement me")
}
