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

type MapNode struct {
	*baseNode
}

// NewMapNode is an exported function used to create a Node of type MapNode.
func NewMapNode(id ID) *MapNode {
	return &MapNode{baseNode: newBaseNode(id)}
}

func (m *MapNode) Clone() (Node, error) {
	return deepCopy(m)
}

func (m *MapNode) Serialize() ([]byte, error) {
	panic("implement me")
}

func (m *MapNode) Deserialize(bytes []byte) error {
	panic("implement me")
}
