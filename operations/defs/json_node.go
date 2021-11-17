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
	"encoding/gob"

	"github.com/pkg/errors"
)

func init() {
	gob.Register(&ListNode{})
	gob.Register(&MapNode{})
	gob.Register(&RegisterNode{})
}

// todo: integrate with lamport counters
type (
	// NodeId is of type string.
	NodeId     string
	ChildNodes map[NodeId]Node

	// Node interface represents the overall primary operations of an JSON node.
	Node interface {
		// GetId returns current Node id.
		GetId() NodeId
		// SetId sets the node with the given id.
		SetId(NodeId) error

		// IsTombStone returns if the given node has already been marked as a tombstone.
		IsTombStone() bool
		// MarkTombstone marks the current node as a tombstone.
		MarkTombstone() error

		// Serialize and Deserialize aids to perform serialization and deserialization
		// of the subtree pointed by current node.
		Serialize() ([]byte, error)
		Deserialize([]byte) error

		// DeepClone performs a deep copy and returns a copied subtree.
		DeepClone() (Node, error)
		// Clone copies just the node and not the subtree.
		Clone() (Node, error)

		// FetchChild returns the children node reachable through the given set of ids as path.
		// traverse the json tree in using []NodeId as a path
		FetchChild([]NodeId) (Node, error)

		// GetChildren returns the children map of current node.
		GetChildren() ChildNodes

		// GetChild returns child with Id if present and not marked as tombstone
		GetChild(NodeId) (Node, error)

		// Assign assigns argument node as a child of the current node.
		Assign(Node, bool) error

		// GetListIndex is only valid if node is a child of a ListNode
		// Get listIndex for node
		GetListIndex() int
		// SetListIndex sets the index value for the node
		SetListIndex(int) error

		// Delete marks the given id a tombstone.
		Delete(NodeId) error
	}

	// BaseNode is a generic type that gets embedded into different Types struct.
	BaseNode struct {
		Id        NodeId
		Tombstone bool
		Children  ChildNodes
		ListIndex int
	}
)

// newBaseNode is an non-exported function and meant for internal usage only.
func newBaseNode(id NodeId) *BaseNode {
	return &BaseNode{
		Id:        id,
		Tombstone: false,
		Children:  make(ChildNodes),
		ListIndex: -1,
	}
}

func (b *BaseNode) GetId() NodeId {
	return b.Id
}

func (b *BaseNode) SetId(id NodeId) error {
	if b.Id != "" {
		return errors.Errorf("id %v is already set for node", b.Id)
	}
	b.Id = id
	return nil
}

func (b *BaseNode) IsTombStone() bool {
	return b.Tombstone
}

func (b *BaseNode) MarkTombstone() error {
	if b.Tombstone {
		return errors.Errorf("node with id %v is already marked tombstone", b.Id)
	}
	b.Tombstone = true
	return nil
}

func (b *BaseNode) GetListIndex() int {
	return b.ListIndex
}

func (b *BaseNode) SetListIndex(idx int) error {
	b.ListIndex = idx
	return nil
}

func (b *BaseNode) Assign(node Node, override bool) error {
	if _, ok := b.Children[node.GetId()]; ok && !override {
		return errors.New("failed to assign child to the given node, node exists with the same id")
	}
	b.Children[node.GetId()] = node

	return nil
}

func (b *BaseNode) GetChildren() ChildNodes {
	return b.Children
}

func (b *BaseNode) GetChild(id NodeId) (Node, error) {
	elem, ok := b.Children[id]
	if ok && !elem.IsTombStone() {
		return elem, nil
	}
	return nil, errors.Errorf("child with id %v doesn't exist for node %v", id, b.Id)
}

func (b *BaseNode) FetchChild(idList []NodeId) (Node, error) {
	var node Node
	children := b.Children

	for i, id := range idList {
		c, ok := children[id]
		if !ok {
			return nil, errors.Errorf("invalid id set, child of id %v doesn't exists for node of id %v", id, b.Id)
		}

		if c.IsTombStone() {
			return nil, errors.Errorf("children of id %v has been marked as tombstone, can't access the subtree", id)
		}

		switch c.(type) {
		case *RegisterNode:
			if i != len(idList)-1 {
				return nil, errors.Errorf("expected empty id[] when a RegisterNode of id %v is reached", c.GetId())
			}
		default:
		}
		node = c
		children = c.GetChildren()
	}
	return node, nil
}
