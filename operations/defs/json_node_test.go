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

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNodeSetId(t *testing.T) {
	assert := assert.New(t)
	tree := &BaseNode{}

	hehe := tree.SetId(NodeId("cat moment"))
	assert.Nil(hehe)
	assert.Equal(tree.GetId(), NodeId("cat moment"))

	err := tree.SetId(NodeId("uwu"))
	assert.NotNil(err)
	assert.Equal(tree.GetId(), NodeId("cat moment"))
}

func TestMarkTombStone(t *testing.T) {
	assert := assert.New(t)
	tree := NewMapNode("cat1")
	assert.False(tree.IsTombStone())
	tree.MarkTombstone()
	assert.True(tree.IsTombStone())
}

func TestSetListIndex(t *testing.T) {
	assert := assert.New(t)
	doc := NewMapNode("cat")
	assert.Equal(doc.GetListIndex(), -1)
	doc.SetListIndex(0)
	assert.Equal(doc.GetListIndex(), 0)
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	doc := NewMapNode("cat")
	doc.Assign(NewMapNode("uwu"), true)
	doc.Assign(NewRegisterNode("onichan", ":3"), true)
	doc.Assign(NewListNode("araara"), true)
	doc.Delete("uwu")
	uwu, _ := doc.GetChild("uwu")
	assert.Nil(uwu)
	doc.Delete("uwu")
	doc.Delete("onichan")
	assert.Equal(len(doc.GetChildren()), 3)
	araara, _ := doc.GetChild("araara")
	assert.Equal(araara.GetId(), NodeId("araara"))
}

func TestFetchChild(t *testing.T) {
	assert := assert.New(t)
	doc := NewMapNode("doc")
	c1 := NewMapNode("c1")
	l1 := NewListNode("l1")
	lc1 := NewMapNode("lc1")
	lc2 := NewMapNode("lc2")
	r1 := NewRegisterNode("r1", 1)
	r2 := NewRegisterNode("r2", 2)
	r3 := NewRegisterNode("r3", 3)
	r4 := NewRegisterNode("r4", 4)
	del := NewRegisterNode("del", "delete")

	doc.Assign(c1, true)
	c1.Assign(r1, true)
	l1.InsertAtHead(lc1)
	l1.InsertAtHead(lc2)
	lc1.Assign(r2, true)
	lc2.Assign(r3, true)
	doc.Assign(l1, true)
	doc.Assign(r4, true) // doc: { "c1": {r1: 1}, "l1": [{ "r3": 3 }, { "r2": 2 }], "r4": 4 }
	doc.Assign(del, true)
	doc.Delete("del")

	for _, test := range []struct {
		arg      []NodeId
		expected Node
		err      error
	}{
		{[]NodeId{NodeId("c1"), NodeId("r1")}, r1, nil},
		{[]NodeId{NodeId("l1"), NodeId("lc1"), NodeId("r2")}, r2, nil},
		{[]NodeId{NodeId("l1"), NodeId("lc2"), NodeId("r3")}, r3, nil},
		{[]NodeId{NodeId("r4")}, r4, nil},
		{[]NodeId{NodeId("r5")}, nil, errors.New("invalid id set, child of id r5 doesn't exists for node of id doc")},
		{[]NodeId{NodeId("del")}, nil, errors.New("children of id del has been marked as tombstone, can't access the subtree")},
		{[]NodeId{NodeId("r4"), NodeId("r4")}, nil, errors.New("expected empty id[] when a RegisterNode of id r4 is reached")},
	} {
		node, e := doc.FetchChild(test.arg)
		assert.Equal(node, test.expected)
		if test.err != nil {
			assert.Equal(e.Error(), test.err.Error())
		} else {
			assert.Nil(e)
		}
	}
}
