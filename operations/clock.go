package operations

import (
"sync/atomic"
)

type Clock struct {
	counter uint64
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
