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

func TestSetValue(t *testing.T) {
	assert := assert.New(t)
	reg := NewRegisterNode("test", 1)
	updatedValue := "hello"

	reg.SetValue(updatedValue)

	// assert
	assert.Equal(reg.GetValue(), updatedValue)
}

func TestChild(t *testing.T) {
	assert := assert.New(t)
	reg := NewRegisterNode("test", 1)

	assert.Nil(reg.GetChildren())

	path := []NodeId{NodeId("bruh"), NodeId("bruh2")}
	assert.Nil(reg.FetchChild(path))
}

func TestRegisterNode_Serialize_Deserialize(t *testing.T) {
	assert := assert.New(t)
	reg := NewRegisterNode("test", 1)

	data, err := reg.Serialize()
	assert.Nil(err, "RegisterNode Serialization failed")
	newReg := &RegisterNode{}
	err = newReg.Deserialize(data)
	assert.Nil(err, "RegisterNode Deserialization failed")
	assert.Equal(reg, newReg, "RegisterNode Deserialization failed")
}
