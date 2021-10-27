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

	assert.Equal(list.Child()["i2"].ListIndex(), 0)
	assert.Equal(list.Child()["i2"].(*RegisterNode).Value(), 2)
	assert.Equal(list.Child()["i3"].ListIndex(), 1)
	assert.Equal(list.Child()["i3"].(*RegisterNode).Value(), 3)
	assert.Equal(list.Child()["i1"].ListIndex(), 2)
	assert.Equal(list.Child()["i1"].(*RegisterNode).Value(), 1)
	assert.Equal(list.Child()["i4"].ListIndex(), 3)
	assert.Equal(list.Child()["i4"].(*RegisterNode).Value(), 4)

}

func TestListDelete(t *testing.T) {
	assert := assert.New(t)
	list := NewListNode("list")
	list.InsertAtHead(NewRegisterNode("i1", 1)) // [1]
	list.InsertAtHead(NewRegisterNode("i2", 2)) // [2, 1]
	list.InsertAfter("i2", NewRegisterNode("i3", 3))
	list.InsertAfter("i1", NewRegisterNode("i4", 4)) // [2, 3, 1, 4]

	list.Delete("i3") // [2, 1, 4]
	assert.Equal(list.Child()["i2"].ListIndex(), 0)
	assert.Equal(list.Child()["i2"].(*RegisterNode).Value(), 2)
	assert.Equal(list.Child()["i1"].ListIndex(), 1)
	assert.Equal(list.Child()["i1"].(*RegisterNode).Value(), 1)
	assert.Equal(list.Child()["i4"].ListIndex(), 2)
	assert.Equal(list.Child()["i4"].(*RegisterNode).Value(), 4)

	list.Delete("i4") // [2, 1]
	assert.Equal(list.Child()["i2"].ListIndex(), 0)
	assert.Equal(list.Child()["i2"].(*RegisterNode).Value(), 2)
	assert.Equal(list.Child()["i1"].ListIndex(), 1)
	assert.Equal(list.Child()["i1"].(*RegisterNode).Value(), 1)

	list.Delete("i2") // [1]
	assert.Equal(list.Child()["i1"].ListIndex(), 0)
	assert.Equal(list.Child()["i1"].(*RegisterNode).Value(), 1)

	assert.Nil(list.Get("i2"))
	assert.NotNil(list.Get("i1"))
}

func TestListAssign(t *testing.T) {
	assert := assert.New(t)
	list := NewListNode("list")

	err1 := list.Assign(NewRegisterNode("fail1", 1))
	list.InsertAtHead(NewRegisterNode("i1", 1))
	err2 := list.Assign(NewRegisterNode("i1", 2))

	reg := NewRegisterNode("i1", 2)
	reg.SetListIndex(0)
	err3 := list.Assign(reg)

	assert.NotNil(err1)
	assert.NotNil(err2)
	assert.NotEqual(err1, err2)
	assert.Nil(err3)

	assert.Equal(list.Get("i1").(*RegisterNode).Value(), 2)
}
