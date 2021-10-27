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
	"github.com/pkg/errors"
)

type ListNode struct {
	*baseNode
}

// NewListNode returns a Node of type ListNode.
func NewListNode(id Id) *ListNode {
	return &ListNode{
		baseNode: newBaseNode(id),
	}
}

func (l *ListNode) Assign(child Node) error {
	// if nodeType is ListT the id should already exist and have same index, use insertAfter to insert new Node for listT
	listElem, ok := l.Child()[child.Id()]
	if !ok {
		return errors.Errorf("invalid id assignment, child of id %v doesn't exists for listNode of id %v", child.Id(), l.id)
	} else if listElem.ListIndex() != child.ListIndex() || listElem.IsTombStone() {
		return errors.Errorf("invalid id assignment, child of id %v is marked as tombstore or has invalid index for listNode of id %v", child.Id(), l.id)
	}

	l.Child()[child.Id()] = child
	return nil
}

func (l *ListNode) InsertAtHead(child Node) error {
	// child.id already present
	_, present := l.Child()[child.Id()]
	if present {
		return errors.Errorf("invalid head insertion, child of id %v already exists for listNode of id %v", child.Id(), l.id)
	}

	// child should have listIndex as 0
	child.SetListIndex(0)

	// increment all other existing elems
	for _id := range l.Child() {
		l.Child()[_id].SetListIndex(l.Child()[_id].ListIndex() + 1)
	}

	l.Child()[child.Id()] = child
	return nil
}

func (l *ListNode) InsertAfter(id Id, child Node) error {
	// child.id already present
	_, present := l.Child()[child.Id()]
	if present {
		return errors.Errorf("invalid list insertion, child of id %v already exists for listNode of id %v", child.Id(), l.id)
	}

	markedElem, ok := l.Child()[id]
	if ok {
		// increment listIndex for all later elements
		for _id := range l.Child() {
			if l.Child()[_id].ListIndex() > markedElem.ListIndex() {
				l.Child()[_id].SetListIndex(l.Child()[_id].ListIndex() + 1)
			}
		}
		child.SetListIndex(markedElem.ListIndex() + 1)
		l.Child()[child.Id()] = child
		return nil
	}
	return errors.Errorf("invalid list insertion, child of id %v doesn't exists for listNode of id %v", id, l.id)
}

func (l *ListNode) Delete(id Id) error {
	child, present := l.Child()[id]
	if present && !child.IsTombStone() {
		child.MarkTombstone()
		// decrement index of elems
		for _id := range l.Child() {
			if l.Child()[_id].ListIndex() > child.ListIndex() {
				l.Child()[_id].SetListIndex(l.Child()[_id].ListIndex() - 1)
			}
		}
	}
	return nil
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
