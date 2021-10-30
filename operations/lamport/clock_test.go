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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	clock := &Clock{0, "xyz"}
	assert := assert.New(t)

	assert.Equal(uint64(0), clock.GetTime(), "Invalid initialisation of clock")
	assert.Equal(uint64(1), clock.Increment(), "Invalid time")
	assert.Equal(uint64(1), clock.GetTime(), "Invalid time")
	for _, test := range []struct {
		arg      Clock
		expected uint64
	}{
		{arg: Clock{10, "abc"}, expected: 11},
		{arg: Clock{10, "def"}, expected: 11},
		{arg: Clock{11, "def"}, expected: 12},
		{arg: Clock{7, "fgh"}, expected: 12},
	} {
		clock.CompareAndUpdate(test.arg.counter)
		assert.Equal(test.expected, clock.GetTime(), "Invalid time")
	}
	clock.SetTime(20)
	assert.Equal(uint64(20), clock.GetTime(), "Invalid time")
	res := clock.IsGreaterThan(&Clock{20, "def"})
	assert.Equal(false, res, "Greater than but concludes lesser")
	res = clock.IsGreaterThan(&Clock{25, "efg"})
	assert.Equal(false, res, "Lesser than but concludes greater")
	res = clock.IsLessThan(&Clock{10, "def"})
	assert.Equal(false, res, "Greater than but concludes lesser")
	res = clock.IsLessThan(&Clock{20, "def"})
	assert.Equal(false, res, "Greater than but concludes lesser")
}
