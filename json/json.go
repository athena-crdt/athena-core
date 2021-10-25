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
	id        string
	children  map[string]*JsonNode // invalid for nodeType RegT
	nodeType  JsonNodeType
	value     interface{} // valid only for nodeType RegT, value can be int, float64, string, bool, nil
	listIndex int         // valid only for list elems
	tombstone bool
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
	if node.nodeType == RegT {
		return
	}
	// if nodeType is ListT the id should already exist and have same index, use insertAfter to insert new Node for listT
	if node.nodeType == ListT {
		listElem, ok := node.children[child.id]
		if !ok {
			return
		} else if listElem.listIndex != child.listIndex || listElem.tombstone {
			return
		}
	}
	node.children[child.id] = child
}

/**
 * Insert at head in ListT node
 */
func (node *JsonNode) InsertAtHead(child *JsonNode) {
	if node.nodeType != ListT {
		return
	}

	// child.id already present
	_, present := node.children[child.id]
	if present {
		return
	}

	// child should have listIndex as 0
	child.listIndex = 0

	// increment all other existing elems
	for _id := range node.children {
		node.children[_id].listIndex += 1
	}

	node.children[child.id] = child
}

/**
 * Insert after given _id in ListT node
 */
func (node *JsonNode) InsertAfter(_id string, child *JsonNode) {
	if node.nodeType != ListT {
		return
	}
	// child.id already present
	_, present := node.children[child.id]
	if present {
		return
	}

	markedElem, ok := node.children[_id]
	if ok {
		// increment listIndex for all later elements
		for _id := range node.children {
			if node.children[_id].listIndex > markedElem.listIndex {
				node.children[_id].listIndex += 1
			}
		}
		child.listIndex = markedElem.listIndex + 1
		node.children[child.id] = child
	}
}

/**
 * setValue for a RegT node
 * args supported nil, string, boolean, int and float64
 */
func (node *JsonNode) SetValue(value interface{}) {
	if node.nodeType != RegT {
		return
	}
	switch value.(type) {
	case int:
		node.value = value
	case float64:
		node.value = value
	case string:
		node.value = value
	case bool:
		node.value = value
	case nil:
		node.value = value
	default:
	}
}

/**
 * Delete a id from MapT or ListT node
 */
func (node *JsonNode) Delete(id string) {
	if node.nodeType == RegT {
		return
	}
	elem, ok := node.children[id]
	if ok {
		node.children[id].tombstone = true

		// if nodeType is listT decrement later indices
		if node.nodeType == ListT {
			for _id := range node.children {
				if node.children[_id].listIndex > elem.listIndex {
					node.children[_id].listIndex -= 1
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
	for k, v := range node.children {
		childrenClone[k] = v
	}
	nodeClone := JsonNode{
		node.id,
		childrenClone,
		node.nodeType,
		node.value,
		node.listIndex,
		node.tombstone,
	}
	return &nodeClone
}
