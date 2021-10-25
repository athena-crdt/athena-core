package operations

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	type compareAndUpdateTest struct {
		arg, expected uint64
	}

	clock := &Clock{}
	assert := assert.New(t)
	compareAndUpdateTests := []compareAndUpdateTest {
		{10, 11},
		{16, 17},
		{8, 17},
		{17, 18},
		{1, 18},
	}

	assert.Equal(uint64(0), clock.GetTime(), "Invalid initialisation of clock")
	assert.Equal(uint64(1), clock.Increment(), "Invalid time")
	assert.Equal(uint64(1), clock.GetTime(), "Invalid time")

	for _, test := range compareAndUpdateTests {
		clock.CompareAndUpdate(test.arg)
		assert.Equal(test.expected, clock.GetTime(), "Invalid time")
	}
}