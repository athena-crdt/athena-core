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
