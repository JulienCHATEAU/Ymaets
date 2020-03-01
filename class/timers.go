package class

import (
	// "fmt"
)

type Timers struct {
	Values []int32
}

func (timers *Timers) Init(count int32) {
	timers.Values = make([]int32, count)
}

func (timers *Timers) Decrement() ([]int32, []int32) {
	notEnded := make([]int32, 0)
	justEnded := make([]int32, 0)
	for index, _ := range timers.Values {
		if timers.Values[index] > 0 {
			timers.Values[index]--
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