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

package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewJsonNode(t *testing.T) {
	assert := assert.New(t)

	mapNode := NewJsonNode("1", MapT)
	listNode := NewJsonNode("2", ListT)
	regNode := NewJsonNode("3", RegT)

	assert.Equal(mapNode.id, "1")
	assert.Equal(mapNode.tombstone, false)
	assert.Equal(mapNode.nodeType, MapT)
	assert.Equal(listNode.id, "2")
	assert.Equal(listNode.tombstone, false)
	assert.Equal(listNode.nodeType, ListT)
	assert.Equal(regNode.id, "3")
	assert.Equal(regNode.tombstone, false)
	assert.Equal(regNode.nodeType, RegT)
}

func TestAssignInMapNodes(t *testing.T) {
	assert := assert.New(t)

	root := NewJsonNode("root", MapT)
	child1 := NewJsonNode("registerChild", RegT)
	child1.SetValue(1)
	child2 := NewJsonNode("listChild", ListT)
	child3 := NewJsonNode("mapChild", MapT)

	root.Assign(child1)
	root.Assign(child2)
	root.Assign(child3)

	assert.Equal(len(root.children), 3)
	assert.Equal(root.children[child1.id], child1)
	assert.Equal(root.children[child2.id], child2)
	assert.Equal(root.children[child3.id], child3)
}

func TestAssignAndInsertInListNodes(t *testing.T) {
	assert := assert.New(t)

	root := NewJsonNode("list", ListT)
	child1 := NewJsonNode("reg1", RegT)
	child2 := NewJsonNode("reg2", RegT)
	child3 := NewJsonNode("reg3", RegT)
	child4 := NewJsonNode("reg3", RegT)

	root.InsertAtHead(child1)
	assert.Equal(len(root.children), 1)
	root.InsertAfter(child2.id, child1)
	assert.Equal(len(root.children), 1)
	root.InsertAfter(child1.id, child3)
	root.InsertAfter(child1.id, child2)
	assert.Equal(len(root.children), 3)
	assert.Equal(root.children[child1.id].listIndex, 0)
	assert.Equal(root.children[child2.id].listIndex, 1)
	assert.Equal(root.children[child3.id].listIndex, 2)
	root.Assign(child4)
	assert.Equal(len(root.children), 3)
}

func TestSetValue(t *testing.T) {
	assert := assert.New(t)

	reg := NewJsonNode("reg", RegT)
	assert.Equal(reg.value, nil)

	reg.SetValue(1)
	assert.Equal(reg.value, 1)

	reg.SetValue("hi")
	assert.Equal(reg.value, "hi")

	reg.SetValue(false)
	assert.Equal(reg.value, false)

	reg.SetValue(3.14)
	assert.Equal(reg.value, 3.14)

	reg.SetValue(nil)
	assert.Equal(reg.value, nil)
}

func TestDeleteInNode(t *testing.T) {
	assert := assert.New(t)

	node := NewJsonNode("root", MapT)
	node.Assign(NewJsonNode("child1", ListT))

	node.Delete("child1")
	assert.True(node.children["child1"].tombstone)
}

func TestNodeClone(t *testing.T) {
	assert := assert.New(t)

	node := NewJsonNode("root", MapT)
	node.Assign(NewJsonNode("child1", RegT))
	node.Assign(NewJsonNode("child2", RegT))
	cpy1 := node.Clone()
	cpy2 := node.Clone()
	node = nil

	assert.Equal(cpy1, cpy2)
}
