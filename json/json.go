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

package json

import (
	"bytes"
	"encoding/gob"
	"log"
)

type JsonNodeType int

/**
 * MapT represents a {} json object
 * ListT represents []
 * RegT represents a value from number, string, boolean, null
 */
const (
	MapT JsonNodeType = iota
	ListT
	RegT
)

/**
 * JsonNode is the basic building block for the
 * internal representation of json in athena-core.
 *
 * todo: add field lastModified LamportTimestamp
 */
type JsonNode struct {
	Id        string
	Children  map[string]*JsonNode // invalid for nodeType RegT
	NodeType  JsonNodeType
	Value     interface{} // valid only for nodeType RegT, value can be int, float64, string, bool, nil
	ListIndex int         // valid only for list elems
	Tombstone bool
}

/**
 * Get an JsonNode with nodeType type
 */
func NewJsonNode(id string, nodeType JsonNodeType) *JsonNode {
	node := JsonNode{
		id,
		make(map[string]*JsonNode),
		nodeType,
		nil,
		-1,
		false,
	}
	return &node
}

/**
 * JsonNode methods
 * A cursor is a vector representing the path in the json tree
 * Each element of cursor is a pair of (key, id)
 * For MapT key and id should be same [mandatory property]
 * For ListT key is index where as id should be a unique identifier
 *
 * todo: Update lastModified field and add LWW rule
 */

/**
 * Assign to child.id for MapT or List node
 */
func (node *JsonNode) Assign(child *JsonNode) {
	if node.NodeType == RegT {
		log.Panicln("Cannot Assign() to JsonNode with NodeType RegT")
		return
	}
	// if nodeType is ListT the id should already exist and have same index, use insertAfter to insert new Node for listT
	if node.NodeType == ListT {
		listElem, ok := node.Children[child.Id]
		if !ok {
			log.Panicln("Cannot Assign() to new Node to JsonNode with NodeType ListT, use InsertAfter or InsertAtHead")
			return
		} else if listElem.ListIndex != child.ListIndex || listElem.Tombstone {
			log.Panicln("Cannot Assign() to new Node to JsonNode with NodeType ListT with invalid ListIndex")
			return
		}
	}
	node.Children[child.Id] = child
}

/**
 * Insert at head in ListT node
 */
func (node *JsonNode) InsertAtHead(child *JsonNode) {
	if node.NodeType != ListT {
		log.Panicln("InsertAtHead() works only for JsonNode with NodeType ListT")
		return
	}

	// child.id already present
	_, present := node.Children[child.Id]
	if present {
		return
	}

	// child should have listIndex as 0
	child.ListIndex = 0

	// increment all other existing elems
	for _id := range node.Children {
		node.Children[_id].ListIndex += 1
	}

	node.Children[child.Id] = child
}

/**
 * Insert after given _id in ListT node
 */
func (node *JsonNode) InsertAfter(_id string, child *JsonNode) {
	if node.NodeType != ListT {
		log.Panicln("InsertAfter() works only for JsonNode with NodeType ListT")
		return
	}
	// child.id already present
	_, present := node.Children[child.Id]
	if present {
		return
	}

	markedElem, ok := node.Children[_id]
	if ok {
		// increment listIndex for all later elements
		for _id := range node.Children {
			if node.Children[_id].ListIndex > markedElem.ListIndex {
				node.Children[_id].ListIndex += 1
			}
		}
		child.ListIndex = markedElem.ListIndex + 1
		node.Children[child.Id] = child
	}
}

/**
 * setValue for a RegT node
 * args supported nil, string, boolean, int and float64
 */
func (node *JsonNode) SetValue(value interface{}) {
	if node.NodeType != RegT {
		log.Panicln("Cannot SetValue() to JsonNode with NodeType ListT or MapT")
		return
	}
	switch value.(type) {
	case int:
		node.Value = value
	case float64:
		node.Value = value
	case string:
		node.Value = value
	case bool:
		node.Value = value
	case nil:
		node.Value = value
	default:
	}
}

/**
 * Delete a id from MapT or ListT node
 */
func (node *JsonNode) Delete(id string) {
	if node.NodeType == RegT {
		log.Panicln("Cannot Delete() child from JsonNode with NodeType RegT")
		return
	}
	elem, ok := node.Children[id]
	if ok {
		node.Children[id].Tombstone = true

		// if nodeType is listT decrement later indices
		if node.NodeType == ListT {
			for _id := range node.Children {
				if node.Children[_id].ListIndex > elem.ListIndex {
					node.Children[_id].ListIndex -= 1
				}
			}
		}
	}
}

/**
 * make a clone of JsonNode
 */
func (node *JsonNode) Clone() *JsonNode {
	// explicitly need to copy the children map
	childrenClone := make(map[string]*JsonNode)
	for k, v := range node.Children {
		childrenClone[k] = v
	}
	nodeClone := JsonNode{
		node.Id,
		childrenClone,
		node.NodeType,
		node.Value,
		node.ListIndex,
		node.Tombstone,
	}
	return &nodeClone
}

/**
 * Recusively flatten the json tree and return a byte array
 */
func (node *JsonNode) Serialize() []byte {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(node); err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

/**
 * Return *JsonNode from binary buffer
 */
func DeserializeJsonNodeBuffer(buf []byte) *JsonNode {
	dec := gob.NewDecoder(bytes.NewBuffer(buf))
	node := new(JsonNode)
	if err := dec.Decode(node); err != nil {
		log.Fatal(err)
	}
	return node
}
