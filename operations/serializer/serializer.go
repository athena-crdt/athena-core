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
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
)

type (
	JsonSerializer struct{}
	GobSerializer  struct{}
	Serializer     interface {
		Serialize(any interface{}) ([]byte, error)
		Deserialize(data []byte, v *interface{}) error
	}
)

func init() {
	fmt.Println("Register types for gob encoding")
}

func (obj *JsonSerializer) Serialize(any interface{}) ([]byte, error) {
	return json.Marshal(any)
}

func (obj *JsonSerializer) Deserialize(data []byte, v *interface{}) error {
	err := json.Unmarshal(data, v)
	return err
}

func (obj *GobSerializer) Serialize(any interface{}) ([]byte, error) {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(&any)
	return buf.Bytes(), err
}

func (obj *GobSerializer) Deserialize(data []byte, v *interface{}) error {
	buf := bytes.Buffer{}
	buf.Write(data)
	dec := gob.NewDecoder(&buf)
	err := dec.Decode(v)
	return err
}
