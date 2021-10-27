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
	tree.Child()["def"] = NewRegisterNode("reg", oldValue)

	// Cloning
	cloneTree, err := deepCopy(tree)
	assert.Nil(err)

	assert.Equal(cloneTree.Id(), ID("abc"))

	ch, ok := cloneTree.Child()["def"]
	assert.True(ok)

	// Updating subtree
	regC, ok := ch.(*RegisterNode)
	assert.True(ok)
	regC.SetValue(9)

	// Should not interfere with original tree
	originalRegC := tree.Child()["def"].(*RegisterNode)
	assert.Equal(originalRegC.Value(), oldValue)
}
