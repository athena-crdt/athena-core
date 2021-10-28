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

package lamport

import (
	"sync/atomic"
)

type Clock struct {
	counter uint64
	hostId  string
}

func (clock *Clock) IsGreaterThan(obj *Clock) bool {
	if atomic.LoadUint64(&obj.counter) == atomic.LoadUint64(&clock.counter) {
		return clock.hostId > obj.hostId
	} else {
		return atomic.LoadUint64(&clock.counter) > atomic.LoadUint64(&obj.counter)
	}
}

func (clock *Clock) IsLessThan(obj *Clock) bool {
	if atomic.LoadUint64(&obj.counter) == atomic.LoadUint64(&clock.counter) {
		return clock.hostId < obj.hostId
	} else {
		return atomic.LoadUint64(&clock.counter) < atomic.LoadUint64(&obj.counter)
	}
}

func (clock *Clock) Increment() uint64 {
	return atomic.AddUint64(&clock.counter, 1)
}

func (clock *Clock) CompareAndUpdate(value uint64) {
	for {
		cur := atomic.LoadUint64(&clock.counter)

		if value < cur {
			return
		}
		if atomic.CompareAndSwapUint64(&clock.counter, cur, value+1) {
			break
		}
	}
}

func (clock *Clock) GetTime() uint64 {
	return atomic.LoadUint64(&clock.counter)
}

func (clock *Clock) SetTime(value uint64) {
	atomic.StoreUint64(&clock.counter, value)
}
