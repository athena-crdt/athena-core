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

type (
	// ID is of type string.
	ID       string
	Children map[ID]Node

	// Node interface represents the overall primary operations of an JSON node.
	Node interface {
		// Id returns current Node id.
		Id() ID

		// SetId sets the node with the given id.
		SetId(ID) error

		// IsTombStone returns if the given node has already been marked as a tombstone.
		IsTombStone() bool
		// MarkTombstone marks the current node as a tombstone.
		MarkTombstone() error

		// Serialize and Deserialize aids to perform serialization and deserialization
		// of the subtree pointed by current node.
		Serialize() error
		Deserialize([]byte) error

		// Clone performs a deepcopy and returns a copied subtree.
		Clone() (Node, error)

		// Assign assigns argument node as a child of the current node.
		Assign(Node, bool) error

		// FetchChild returns the children node reachable through the given set of ids.
		FetchChild([]ID) (Node, error)

		// Child returns the children map of current node.
		Child() Children
	}

	// baseNode is a generic type that gets embedded into different Types struct.
	baseNode struct {
		id        ID
		tombstone bool
		children  Children
	}
)

// newBaseNode is an non-exported function and meant for internal usage only.
func newBaseNode(id ID) *baseNode {
	return &baseNode{
		id:        id,
		tombstone: false,
		children:  make(Children),
	}
}

func (b *baseNode) Id() ID {
	return b.id
}

func (b *baseNode) SetId(id ID) error {
	if id != "" {
		return errors.New("ID has already been set, once set it can't be altered")
	}
	b.id = id
	return nil
}

func (b *baseNode) IsTombStone() bool {
	return b.tombstone
}

func (b *baseNode) MarkTombstone() error {
	if b.tombstone {
		return errors.New("node has already been marked a tombstone")
	}
	b.tombstone = true
	return nil
}

func (b *baseNode) Assign(node Node, override bool) error {
	if _, ok := b.children[node.Id()]; ok && !override {
		return errors.New("failed to assign child to the given node, node exists with the same id")
	}
	b.children[node.Id()] = node

	return nil
}

func (b *baseNode) Child() Children {
	return b.children
}

func (b *baseNode) FetchChild(idList []ID) (Node, error) {
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
