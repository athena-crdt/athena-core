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

func TestMapAssign(t *testing.T) {
	assert := assert.New(t)
	tree := NewMapNode("test")
	c1 := NewMapNode("test1")
	c2 := NewMapNode("test2")
	c3 := NewMapNode("test3")

	c1.Assign(c2, true)
	tree.Assign(c1, true)
	tree.Assign(c3, true)

	_, c1Present := tree.Children()["test1"]
	assert.True(c1Present)
	_, c2Present := tree.Children()["test1"].Children()["test2"]
	assert.True(c2Present)
	_, c3Present := tree.Children()["test3"]
	assert.True(c3Present)

	// assert
	assert.Equal(tree.Children()["test1"], c1)
	assert.Equal(tree.Children()["test1"].Id(), NodeId("test1"))
	assert.Equal(tree.Children()["test3"], c3)
	assert.Equal(tree.Children()["test3"].Id(), NodeId("test3"))
	assert.Equal(tree.Children()["test1"].Children()["test2"], c2)
	assert.Equal(tree.Children()["test1"].Children()["test2"].Id(), NodeId("test2"))
	assert.Equal(len(tree.Children()), 2)
}

func TestMapDelete(t *testing.T) {
	assert := assert.New(t)
	tree := NewMapNode("test")
	c1 := NewMapNode("test1")
	c2 := NewMapNode("test2")

	tree.Assign(c1, true)
	tree.Assign(c2, true)

	tree.Delete("test1")
	assert.Nil(tree.Child("test1"))
	test2, _ := tree.Child("test2")
	assert.Equal(test2.Id(), NodeId("test2"))
}
