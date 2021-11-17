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

import (
	serializer2 "github.com/athena-crdt/athena-core/operations/serializer"
	"github.com/pkg/errors"
)

type MapNode struct {
	*BaseNode
}

// NewMapNode is an exported function used to create a Node of type MapNode.
func NewMapNode(id NodeId) *MapNode {
	return &MapNode{BaseNode: newBaseNode(id)}
}

func (m *MapNode) Delete(id NodeId) error {
	child, present := m.GetChildren()[id]
	if present && !child.IsTombStone() {
		child.MarkTombstone()
		return nil
	}
	return errors.Errorf("Cannot delete id %v from mapNode of id %v", id, m.GetId())
}

func (m *MapNode) DeepClone() (Node, error) {
	return deepCopy(m, true)
}

func (m *MapNode) Clone() (Node, error) {
	return deepCopy(m, false)
}

func (m *MapNode) Serialize() ([]byte, error) {
	serializer := serializer2.GobSerializer{}
	return serializer.Serialize(m)
}

func (m *MapNode) Deserialize(bytes []byte) error {
	deserializer := serializer2.GobSerializer{}
	err := deserializer.Deserialize(bytes, m)
	return err
}
