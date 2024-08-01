package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := Queue{
		store: make([]int, 0, 10),
	}

	queue.Add(10)
	queue.Add(32)
	f := queue.Add(99)
	assert.Equal(t, f, []int{10, 32, 99})
	queue.Add(44)
	queue.Remove()
	queue.Add(78)
	f = queue.Remove()
	assert.Equal(t, f, []int{99, 44, 78})

}
