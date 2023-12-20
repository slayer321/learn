package main

import (
	"sync/atomic"
	"testing"
)

// func TestDataRaceCondition(t *testing.T) {
// 	var state int32

// 	//var mu sync.RWMutex
// 	mu := sync.Mutex{}

// 	for i := 0; i < 10; i++ {
// 		go func(i int) {
// 			mu.Lock()
// 			state += int32(i)
// 			mu.Unlock()
// 		}(i)
// 	}

// 	//fmt.Println(state)

// }

func TestDataRaceConditionAtomic(t *testing.T) {
	var state int32

	for i := 0; i < 10; i++ {
		go func(i int) {
			//state += int32(i)
			atomic.AddInt32(&state, int32(i))
		}(i)
	}

	//fmt.Println(state)

}
