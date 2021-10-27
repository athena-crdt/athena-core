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
	tree := &baseNode{}

	hehe := tree.SetId(Id("cat moment"))
	assert.Nil(hehe)
	assert.Equal(tree.Id(), Id("cat moment"))

	err := tree.SetId(Id("uwu"))
	assert.NotNil(err)
	assert.Equal(tree.Id(), Id("cat moment"))
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
	assert.Equal(doc.ListIndex(), -1)
	doc.SetListIndex(0)
	assert.Equal(doc.ListIndex(), 0)
}

func TestGet(t *testing.T) {
	assert := assert.New(t)
	doc := NewMapNode("cat")
	doc.Assign(NewMapNode("uwu"))
	doc.Assign(NewRegisterNode("onichan", ":3"))
	doc.Assign(NewListNode("araara"))
	doc.Delete("uwu")
	assert.Nil(doc.Get("uwu"))
	doc.Delete("uwu")
	doc.Delete("onichan")
	assert.Equal(len(doc.Child()), 3)
	assert.Equal(doc.Get("araara").Id(), Id("araara"))
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

	doc.Assign(c1)
	c1.Assign(r1)
	l1.InsertAtHead(lc1)
	l1.InsertAtHead(lc2)
	lc1.Assign(r2)
	lc2.Assign(r3)
	doc.Assign(l1)
	doc.Assign(r4) // doc: { "c1": {r1: 1}, "l1": [{ "r3": 3 }, { "r2": 2 }], "r4": 4 }
	doc.Assign(del)
	doc.Delete("del")

	for _, test := range []struct {
		arg      []Id
		expected Node
		err      error
	}{
		{[]Id{Id("c1"), Id("r1")}, r1, nil},
		{[]Id{Id("l1"), Id("lc1"), Id("r2")}, r2, nil},
		{[]Id{Id("l1"), Id("lc2"), Id("r3")}, r3, nil},
		{[]Id{Id("r4")}, r4, nil},
		{[]Id{Id("r5")}, nil, errors.New("invalid id set, child of id r5 doesn't exists for node of id doc")},
		{[]Id{Id("del")}, nil, errors.New("children of id del has been marked as tombstone, can't access the subtree")},
		{[]Id{Id("r4"), Id("r4")}, nil, errors.New("expected empty id[] when a RegisterNode of id r4 is reached")},
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
