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
	tc := map[string]interface{}{
		"x":    1,
		"list": []byte{65, 66, 67},
	}
	var obj Serializer
	// Testing Json Serializer and Deserializer
	obj = &JsonSerializer{}
	data, err := obj.Serialize(tc)
	assert.Equal(t, nil, err, err)
	v, err := obj.Deserialize(data)
	assert.Equal(t, nil, err, err)
	newData, err := obj.Serialize(v)
	assert.Equal(t, data, newData, "Json Deserialization failed")
}
