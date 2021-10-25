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

	assert.Equal(mapNode.Id, "1")
	assert.Equal(mapNode.Tombstone, false)
	assert.Equal(mapNode.NodeType, MapT)
	assert.Equal(listNode.Id, "2")
	assert.Equal(listNode.Tombstone, false)
	assert.Equal(listNode.NodeType, ListT)
	assert.Equal(regNode.Id, "3")
	assert.Equal(regNode.Tombstone, false)
	assert.Equal(regNode.NodeType, RegT)
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

	assert.Equal(len(root.Children), 3)
	assert.Equal(root.Children[child1.Id], child1)
	assert.Equal(root.Children[child2.Id], child2)
	assert.Equal(root.Children[child3.Id], child3)
}

func TestAssignAndInsertInListNodes(t *testing.T) {
	assert := assert.New(t)

	root := NewJsonNode("list", ListT)
	child1 := NewJsonNode("reg1", RegT)
	child2 := NewJsonNode("reg2", RegT)
	child3 := NewJsonNode("reg3", RegT)
	child4 := NewJsonNode("reg3", RegT)
	child4.ListIndex = 2

	root.InsertAtHead(child1)
	assert.Equal(len(root.Children), 1)
	root.InsertAfter(child2.Id, child1)
	assert.Equal(len(root.Children), 1)
	root.InsertAfter(child1.Id, child3)
	root.InsertAfter(child1.Id, child2)
	assert.Equal(len(root.Children), 3)
	assert.Equal(root.Children[child1.Id].ListIndex, 0)
	assert.Equal(root.Children[child2.Id].ListIndex, 1)
	assert.Equal(root.Children[child3.Id].ListIndex, 2)
	root.Assign(child4)
	assert.Equal(len(root.Children), 3)
}

func TestSetValue(t *testing.T) {
	assert := assert.New(t)

	reg := NewJsonNode("reg", RegT)
	assert.Equal(reg.Value, nil)

	reg.SetValue(1)
	assert.Equal(reg.Value, 1)

	reg.SetValue("hi")
	assert.Equal(reg.Value, "hi")

	reg.SetValue(false)
	assert.Equal(reg.Value, false)

	reg.SetValue(3.14)
	assert.Equal(reg.Value, 3.14)

	reg.SetValue(nil)
	assert.Equal(reg.Value, nil)
}

func TestDeleteInNode(t *testing.T) {
	assert := assert.New(t)

	node := NewJsonNode("root", MapT)
	node.Assign(NewJsonNode("child1", ListT))

	node.Delete("child1")
	assert.True(node.Children["child1"].Tombstone)
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

func TestSerialization(t *testing.T) {
	assert := assert.New(t)

	node := NewJsonNode("root", MapT)
	child := NewJsonNode("child", MapT)
	node.Assign(child)
	child.Assign(NewJsonNode("child1", RegT))
	child.Children["child1"].SetValue("hello")
	child.Assign(NewJsonNode("child2", RegT))
	child.Children["child2"].SetValue(1)
	child.Assign(NewJsonNode("child3", RegT))
	child.Children["child3"].SetValue(false)

	buf := node.Serialize()
	ret := DeserializeJsonNodeBuffer(buf)

	assert.Equal(ret, node)
	assert.Equal(ret.Children["child"].Children["child1"].Value, "hello")
	assert.Equal(ret.Children["child"].Children["child2"].Value, 1)
	assert.Equal(ret.Children["child"].Children["child3"].Value, false)
}
