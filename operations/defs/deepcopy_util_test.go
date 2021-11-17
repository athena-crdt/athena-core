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

func TestDeepCopy(t *testing.T) {
	assert := assert.New(t)
	oldValue := 6
	tree := NewMapNode("abc")
	tree.Assign(NewRegisterNode("def", oldValue), true)

	// Cloning
	cloneTree, err := deepCopy(tree, true)

	assert.Nil(err)
	assert.Equal(cloneTree.GetId(), NodeId("abc"))

	ch, ok := cloneTree.GetChildren()["def"]
	assert.True(ok)

	// Updating subtree
	regC, ok := ch.(*RegisterNode)
	assert.True(ok)
	regC.SetValue(9)

	// Should not interfere with original tree
	originalRegC := tree.GetChildren()["def"].(*RegisterNode)
	assert.Equal(originalRegC.GetValue(), oldValue)
}

func TestShallowCopy(t *testing.T) {
	assert := assert.New(t)
	oldValue := 6
	tree := NewMapNode("abc")
	tree.Assign(NewRegisterNode("def", oldValue), true)

	// Cloning
	cloneTree, err := deepCopy(tree, false)

	assert.Nil(err)
	assert.Equal(cloneTree.GetId(), NodeId("abc"))

	ch, ok := cloneTree.GetChildren()["def"]
	assert.True(ok)

	// Updating subtree
	newValue := 9
	regC, ok := ch.(*RegisterNode)
	assert.True(ok)
	regC.SetValue(newValue)

	// original reg should be updated
	originalRegC := tree.GetChildren()["def"].(*RegisterNode)
	assert.Equal(originalRegC.GetValue(), newValue)
}
