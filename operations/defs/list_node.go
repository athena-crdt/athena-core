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

type ListNode struct {
	*BaseNode
}

// NewListNode returns a Node of type ListNode.
func NewListNode(id NodeId) *ListNode {
	return &ListNode{
		BaseNode: newBaseNode(id),
	}
}

func (l *ListNode) Assign(child Node, _ bool) error {
	if elm, ok := l.GetChildren()[child.GetId()]; !ok {
		return errors.Errorf("invalid id assignment, child of id %v doesn't exists for listNode of id %v", child.GetId(), l.Id)
	} else if elm.GetListIndex() != child.GetListIndex() {
		return errors.Errorf("invalid id assignment, child of id %v  has invalid index for listNode of id %v", child.GetId(), l.Id)
	} else if elm.IsTombStone() {
		return errors.Errorf("invalid id assignment, child of id %v is marked as tombstone id %v", child.GetId(), l.Id)
	}
	return l.BaseNode.Assign(child, true)
}

func (l *ListNode) InsertAtHead(child Node) error {
	// child.id already present
	_, present := l.GetChildren()[child.GetId()]
	if present {
		return errors.Errorf("invalid head insertion, child of id %v already exists for listNode of id %v", child.GetId(), l.Id)
	}

	// child should have listIndex as 0
	child.SetListIndex(0)

	// increment all other existing elems
	for id := range l.GetChildren() {
		l.GetChildren()[id].SetListIndex(l.GetChildren()[id].GetListIndex() + 1)
	}

	l.GetChildren()[child.GetId()] = child
	return nil
}

func (l *ListNode) InsertAfter(id NodeId, child Node) error {
	// child.id already present
	_, present := l.GetChildren()[child.GetId()]
	if present {
		return errors.Errorf("invalid list insertion, child of id %v already exists for listNode of id %v", child.GetId(), l.Id)
	}

	markedElem, ok := l.GetChildren()[id]
	if !ok {
		return errors.Errorf("invalid list insertion, child of id %v doesn't exists for listNode of id %v", id, l.Id)
	}

	// increment listIndex for all later elements
	for id := range l.GetChildren() {
		if l.GetChildren()[id].GetListIndex() > markedElem.GetListIndex() {
			l.GetChildren()[id].SetListIndex(l.GetChildren()[id].GetListIndex() + 1)
		}
	}
	child.SetListIndex(markedElem.GetListIndex() + 1)
	l.GetChildren()[child.GetId()] = child
	return nil
}

func (l *ListNode) Delete(id NodeId) error {
	child, present := l.GetChildren()[id]
	if present && !child.IsTombStone() {
		child.MarkTombstone()
		// decrement index of elems
		for id := range l.GetChildren() {
			if l.GetChildren()[id].GetListIndex() > child.GetListIndex() {
				l.GetChildren()[id].SetListIndex(l.GetChildren()[id].GetListIndex() - 1)
			}
		}
		return nil
	}
	return errors.Errorf("Cannot delete id %v from listNode of id %v", id, l.GetId())
}

func (l *ListNode) DeepClone() (Node, error) {
	return deepCopy(l, true)
}

func (l *ListNode) Clone() (Node, error) {
	return deepCopy(l, false)
}

func (l *ListNode) Serialize() ([]byte, error) {
	serializer := serializer2.GobSerializer{}
	return serializer.Serialize(l)
}

func (l *ListNode) Deserialize(bytes []byte) error {
	deserializer := serializer2.GobSerializer{}
	err := deserializer.Deserialize(bytes, l)
	return err
}
