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

// todo: integrate with lamport counters
type (
	// Id is of type string.
	Id       string
	Children map[Id]Node

	// Node interface represents the overall primary operations of an JSON node.
	Node interface {
		// Id returns current Node id.
		Id() Id
		// SetId sets the node with the given id.
		SetId(Id) error

		// IsTombStone returns if the given node has already been marked as a tombstone.
		IsTombStone() bool
		// Delele marks the current node as a tombstone.
		MarkTombstone() error

		// Serialize and Deserialize aids to perform serialization and deserialization
		// of the subtree pointed by current node.
		Serialize() ([]byte, error)
		Deserialize([]byte) error

		// DeepClone performs a deepcopy and returns a copied subtree.
		DeepClone() (Node, error)
		// Clone copies just the node and not the subtree.
		Clone() (Node, error)

		// FetchChild returns the children node reachable through the given set of ids as path.
		// travserse the json tree in using []Id as a path
		FetchChild([]Id) (Node, error)

		// Child returns the children map of current node.
		Child() Children

		// Get returns child with Id if present and not marked as tombstone
		Get(id Id) Node

		// ListIndex is only valid if node is a child of a ListNode
		// Get listIndex for node
		ListIndex() int
		// SetListIndex sets the index value for the node
		SetListIndex(int) error
	}

	// baseNode is a generic type that gets embedded into different Types struct.
	baseNode struct {
		id        Id
		tombstone bool
		children  Children
		listIndex int
	}
)

// newBaseNode is an non-exported function and meant for internal usage only.
func newBaseNode(id Id) *baseNode {
	return &baseNode{
		id:        id,
		tombstone: false,
		children:  make(Children),
		listIndex: -1,
	}
}

func (b *baseNode) Id() Id {
	return b.id
}

func (b *baseNode) SetId(id Id) error {
	if id != "" {
		return errors.Errorf("id %v is already set for node", b.id)
	}
	b.id = id
	return nil
}

func (b *baseNode) IsTombStone() bool {
	return b.tombstone
}

func (b *baseNode) MarkTombstone() error {
	if b.tombstone {
		return errors.Errorf("node with id %v is already marked tombstone", b.id)
	}
	b.tombstone = true
	return nil
}

func (b *baseNode) ListIndex() int {
	return b.listIndex
}

func (b *baseNode) SetListIndex(idx int) error {
	b.listIndex = idx
	return nil
}

func (b *baseNode) Child() Children {
	return b.children
}

func (b *baseNode) Get(id Id) Node {
	elem, ok := b.children[id]
	if ok && !elem.IsTombStone() {
		return elem
	}
	return nil
}

func (b *baseNode) FetchChild(idList []Id) (Node, error) {
	var node Node
	children := b.children

	for i, id := range idList {
		c, ok := children[id]
		if !ok {
			return nil, errors.Errorf("invalid id set, child of id %v doesn't exists for node of id %v", id, b.id)
		}

		if c.IsTombStone() {
			return nil, errors.Errorf("children of id %v has been marked as tombstone, can't access the subtree", id)
		}

		switch c.(type) {
		case *RegisterNode:
			if i != len(idList)-1 {
				return nil, errors.Errorf("expected empty idList when a RegisterNode of id %v is reached", c.Id())
			}
		default:
		}
		node = c
		children = c.Child()
	}
	return node, nil
}
