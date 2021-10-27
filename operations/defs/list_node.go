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

type ListNode struct {
	*baseNode
	listIndex uint64
}

// NewListNode returns a Node of type ListNode.
func NewListNode(id ID) *ListNode {
	return &ListNode{
		baseNode: newBaseNode(id),
	}
}

func (l *ListNode) Clone() (Node, error) {
	return deepCopy(l)
}

// Index returns the current index of ListNode.
func (l *ListNode) Index() uint64 {
	return l.listIndex
}

// SetIndex sets the index of currentNode with the given index.
func (l *ListNode) SetIndex(idx uint64) {
	l.listIndex = idx
}

func (l *ListNode) InsertAtHead(child Node) error {
	panic("implement")
}

func (l *ListNode) InsertAfter(id string, child Node) error {
	panic("implement")
}

func (l *ListNode) Serialize() error {
	panic("implement me")
}

func (l *ListNode) Deserialize(bytes []byte) error {
	panic("implement me")
}
