package class

import (
	// "fmt"
)

type TimeCounter struct {
	Active	bool
	Counter int32
	Time		int32
}

func (tc *TimeCounter) Off() {
	tc.Active = false
}

func (tc *TimeCounter) On() {
	tc.Active = true
}

func (tc *TimeCounter) Reset() {
	tc.Counter = 0
}

func (tc *TimeCounter) Init(active bool, time int32) {
	tc.Active = active
	tc.Counter = 0
	tc.Time = time
}

func (tc *TimeCounter) Increment() bool {
	ended := false
	if tc.Active {
		tc.Counter++
		if tc.Counter == tc.Time {
			ended = true
			tc.Active = false
			tc.Counter = 0
		}
	}
	return ended
}

/////////////

var secondCounter int32 = 0

func IncrementSeconds(_map *Map) {
	secondCounter++
	_map.CurrPlayer.EverySecAction()
	if secondCounter % 2 == 0 {
		secondCounter = 0
		_map.CurrPlayer.Every2SecAction()
	}
}

////////////

type Timers struct {
	Values []int32
	Decrements []int32
}

func (timers *Timers) Init(count int32) {
	timers.Values = make([]int32, count)
	timers.Decrements = make([]int32, count)
	for index, _ := range timers.Decrements {
		timers.Decrements[index] = 1
	}
}

func (timers *Timers) Decrement() ([]int32, []int32) {
	notEnded := make([]int32, 0)
	justEnded := make([]int32, 0)
	for index, _ := range timers.Values {
		if timers.Values[index] > 0 {
			timers.Values[index] -= timers.Decrements[index]
			if timers.Values[index] > 0 {
				notEnded = append(notEnded, int32(index))
			}
			if timers.Values[index] == 0 {
				justEnded = append(justEnded, int32(index))
			}
		}
	}
	return notEnded, justEnded
}