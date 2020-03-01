package class

import (
)

type Timers struct {
	Values []int32
}

func (timers *Timers) Init(count int32) {
	timers.Values = make([]int32, count)
}

func (timers *Timers) SetTimer(index int32, value int32) {
	timers.Values[index] = value
}

func (timers *Timers) Decrement() []int32 {
	notEnded := make([]int32, 0)
	for index, timer := range timers.Values {
		if timer > 0 {
			timers.Values[index]--
			if timer > 0 {
				notEnded = append(notEnded, int32(index))
			}
		}
	}
	return notEnded
}