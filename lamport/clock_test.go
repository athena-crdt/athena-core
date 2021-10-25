package lamport

import (
	"testing"
)

type compareAndUpdateTest struct {
	arg, expected uint64
}

func TestClock(t *testing.T) {
	clock := &Clock{}
	compareAndUpdateTests := []compareAndUpdateTest {
		{10, 11},
		{16, 17},
		{8, 17},
		{17, 18},
		{1, 18},
	}

	if clock.GetTime() != 0 {
		t.Errorf("Invalid initialisation of clock")
	}

	if clock.Increment() != 1 {
		t.Errorf("Invalid time")
	}

	if clock.GetTime() != 1 {
		t.Errorf("Invalid time")
	}

	for _, test := range compareAndUpdateTests {
		if clock.CompareAndUpdate(test.arg); clock.GetTime() != test.expected {
			t.Errorf("Invalid time")
		}
	}
}