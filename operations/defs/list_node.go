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

type ListNode struct {
	*baseNode
}

func init() {
	gob.Register(&ListNode{})
}

// NewListNode returns a Node of type ListNode.
func NewListNode(id NodeId) *ListNode {
	return &ListNode{
		baseNode: newBaseNode(id),
	}
}

func (l *ListNode) Assign(child Node, _ bool) error {
	if elm, ok := l.Children()[child.Id()]; !ok {
		return errors.Errorf("invalid id assignment, child of id %v doesn't exists for listNode of id %v", child.Id(), l.id)
	} else if elm.ListIndex() != child.ListIndex() {
		return errors.Errorf("invalid id assignment, child of id %v  has invalid index for listNode of id %v", child.Id(), l.id)
	} else if elm.IsTombStone() {
		return errors.Errorf("invalid id assignment, child of id %v is marked as tombstone id %v", child.Id(), l.id)
	}
	return l.baseNode.Assign(child, true)
}

func (l *ListNode) InsertAtHead(child Node) error {
	// child.id already present
	_, present := l.Children()[child.Id()]
	if present {
		return errors.Errorf("invalid head insertion, child of id %v already exists for listNode of id %v", child.Id(), l.id)
	}

	// child should have listIndex as 0
	child.SetListIndex(0)

	// increment all other existing elems
	for id := range l.Children() {
		l.Children()[id].SetListIndex(l.Children()[id].ListIndex() + 1)
	}

	l.Children()[child.Id()] = child
	return nil
}

func (l *ListNode) InsertAfter(id NodeId, child Node) error {
	// child.id already present
	_, present := l.Children()[child.Id()]
	if present {
		return errors.Errorf("invalid list insertion, child of id %v already exists for listNode of id %v", child.Id(), l.id)
	}

	markedElem, ok := l.Children()[id]
	if !ok {
		return errors.Errorf("invalid list insertion, child of id %v doesn't exists for listNode of id %v", id, l.id)
	}

	// increment listIndex for all later elements
	for id := range l.Children() {
		if l.Children()[id].ListIndex() > markedElem.ListIndex() {
			l.Children()[id].SetListIndex(l.Children()[id].ListIndex() + 1)
		}
	}
	child.SetListIndex(markedElem.ListIndex() + 1)
	l.Children()[child.Id()] = child
	return nil
}

func (l *ListNode) Delete(id NodeId) error {
	child, present := l.Children()[id]
	if present && !child.IsTombStone() {
		child.MarkTombstone()
		// decrement index of elems
		for id := range l.Children() {
			if l.Children()[id].ListIndex() > child.ListIndex() {
				l.Children()[id].SetListIndex(l.Children()[id].ListIndex() - 1)
			}
		}
		return nil
	}
	return errors.Errorf("Cannot delete id %v from listNode of id %v", id, l.Id())
}

func (l *ListNode) DeepClone() (Node, error) {
	return deepCopy(l, true)
}

func (l *ListNode) Clone() (Node, error) {
	return deepCopy(l, false)
}

func (l *ListNode) Serialize() ([]byte, error) {
	panic("implement me")
}

func (l *ListNode) Deserialize(bytes []byte) error {
	panic("implement me")
}
