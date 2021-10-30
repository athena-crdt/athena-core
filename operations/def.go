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

package operations

import (
	"sort"
	"sync"

	"github.com/pkg/errors"

	"github.com/athena-crdt/athena-core/operations/defs"
	"github.com/athena-crdt/athena-core/operations/lamport"
)

type (
	// Mutation is the type of operation
	Mutation uint8

	// Operation - individual operation entity
	Operation struct {
		Id       *lamport.Clock
		Deps     []*lamport.Clock
		Cursor   []defs.NodeId
		Mutation Mutation
		Value    defs.Node
	}

	// OpStore is a thread safe Operation queue.
	OpStore struct {
		sync.RWMutex
		Store []*Operation
	}
)

const (
	ASSIGN = Mutation(iota)
	INSERT
	DELETE
	GET
)

// NewOpStore creates a new operation store.
func NewOpStore() *OpStore {
	return &OpStore{
		RWMutex: sync.RWMutex{},
		Store:   []*Operation{},
	}
}

func (o *Operation) SortDeps() {
	sort.Slice(o.Deps, func(i, j int) bool {
		return o.Deps[i].GetTime() <= o.Deps[j].GetTime()
	})
}

func (op *OpStore) Push(o *Operation) {
	op.Lock()
	defer op.Unlock()
	op.Store = append(op.Store, o)
}

func (op *OpStore) SortedPush(o *Operation) {
	op.Lock()
	defer op.Unlock()
	op.Store = append(op.Store, o)
	sort.Slice(op.Store, func(i, j int) bool {
		return op.Store[i].Id.GetTime() <= op.Store[j].Id.GetTime()
	})
}

func (op *OpStore) IsEmpty() bool {
	return op.Len() == 0
}

func (op *OpStore) Len() int {
	op.RLock()
	defer op.RUnlock()
	return len(op.Store)
}

func (op *OpStore) Pop() (*Operation, error) {
	op.Lock()
	defer op.Unlock()
	if len(op.Store) == 0 {
		return nil, errors.New("empty queue")
	}

	elm := op.Store[0]
	op.Store = op.Store[1:]
	return elm, nil
}

func (op *OpStore) Sort() {
	sort.Slice(op.Store, func(i, j int) bool {
		return op.Store[i].Id.GetTime() <= op.Store[j].Id.GetTime()
	})
}

func (op *OpStore) Serialize() ([]byte, error) {
	op.RLock()
	defer op.RUnlock()
	panic("implement after PR #12")
}
