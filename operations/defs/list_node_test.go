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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListInsert(t *testing.T) {
	assert := assert.New(t)
	list := NewListNode("list")
	list.InsertAtHead(NewRegisterNode("i1", 1)) // [1]
	list.InsertAtHead(NewRegisterNode("i2", 2)) // [2, 1]
	list.InsertAfter("i2", NewRegisterNode("i3", 3))
	list.InsertAfter("i1", NewRegisterNode("i4", 4)) // [2, 3, 1, 4]

	assert.Equal(list.GetChildren()["i2"].GetListIndex(), 0)
	assert.Equal(list.GetChildren()["i2"].(*RegisterNode).GetValue(), 2)
	assert.Equal(list.GetChildren()["i3"].GetListIndex(), 1)
	assert.Equal(list.GetChildren()["i3"].(*RegisterNode).GetValue(), 3)
	assert.Equal(list.GetChildren()["i1"].GetListIndex(), 2)
	assert.Equal(list.GetChildren()["i1"].(*RegisterNode).GetValue(), 1)
	assert.Equal(list.GetChildren()["i4"].GetListIndex(), 3)
	assert.Equal(list.GetChildren()["i4"].(*RegisterNode).GetValue(), 4)

}

func TestListDelete(t *testing.T) {
	assert := assert.New(t)
	list := NewListNode("list")
	list.InsertAtHead(NewRegisterNode("i1", 1)) // [1]
	list.InsertAtHead(NewRegisterNode("i2", 2)) // [2, 1]
	list.InsertAfter("i2", NewRegisterNode("i3", 3))
	list.InsertAfter("i1", NewRegisterNode("i4", 4)) // [2, 3, 1, 4]

	list.Delete("i3") // [2, 1, 4]
	assert.Equal(list.GetChildren()["i2"].GetListIndex(), 0)
	assert.Equal(list.GetChildren()["i2"].(*RegisterNode).GetValue(), 2)
	assert.Equal(list.GetChildren()["i1"].GetListIndex(), 1)
	assert.Equal(list.GetChildren()["i1"].(*RegisterNode).GetValue(), 1)
	assert.Equal(list.GetChildren()["i4"].GetListIndex(), 2)
	assert.Equal(list.GetChildren()["i4"].(*RegisterNode).GetValue(), 4)

	list.Delete("i4") // [2, 1]
	assert.Equal(list.GetChildren()["i2"].GetListIndex(), 0)
	assert.Equal(list.GetChildren()["i2"].(*RegisterNode).GetValue(), 2)
	assert.Equal(list.GetChildren()["i1"].GetListIndex(), 1)
	assert.Equal(list.GetChildren()["i1"].(*RegisterNode).GetValue(), 1)

	list.Delete("i2") // [1]
	assert.Equal(list.GetChildren()["i1"].GetListIndex(), 0)
	assert.Equal(list.GetChildren()["i1"].(*RegisterNode).GetValue(), 1)

	assert.Nil(list.GetChild("i2"))
	assert.NotNil(list.GetChild("i1"))
}

func TestListAssign(t *testing.T) {
	assert := assert.New(t)
	list := NewListNode("list")

	assert.NotNil(list.Assign(NewRegisterNode("fail1", 1), true))
	assert.Nil(list.InsertAtHead(NewRegisterNode("i1", 1)))
	assert.NotNil(list.Assign(NewRegisterNode("i1", 2), true))

	reg := NewRegisterNode("i1", 2)
	reg.SetListIndex(0)
	assert.Nil(list.Assign(reg, true))

	i1, err := list.GetChild("i1")
	assert.Nil(err)
	assert.Equal(i1.(*RegisterNode).GetValue(), 2)
}

func TestListNode_Serialize_Deserialize(t *testing.T) {
	assert := assert.New(t)
	listNode := NewListNode("test")

	data, err := listNode.Serialize()
	assert.Nil(err, "ListNode Serialization failed")
	newListNode := &ListNode{}
	err = newListNode.Deserialize(data)
	assert.Nil(err, "ListNode Deserialization failed")
	assert.Equal(listNode, newListNode, "ListNode Deserialization failed")
}
