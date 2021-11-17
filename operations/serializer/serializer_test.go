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

package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSerializer(t *testing.T) {
	x := &map[string]interface{}{
		"X": 1,
		"Y": []byte{1, 2, 2},
	}
	serializer := &GobSerializer{}
	data, err := serializer.Serialize(x)
	assert.Equal(t, nil, err, "Serialization failed")

	newX := &map[string]interface{}{}
	err = serializer.Deserialize(data, newX)
	assert.Equal(t, nil, err, "Deserialization failed")
	data2, err := serializer.Serialize(newX)
	assert.Equal(t, nil, err, "Deserialization failed")
	assert.Equal(t, data, data2, "Deserialization failed")
}
