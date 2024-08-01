package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushAndPop(t *testing.T) {
	stack := Stack{
		store: make([]int, 0, 10),
	}
	stack.Push(4)
	stack.Push(54)
	f := stack.Push(12)
	assert.Equal(t, f, []int{4, 54, 12})

	stack.Pop()
	stack.Push(99)
	stack.Push(55)
	f = stack.Pop()

	assert.Equal(t, f, []int{4, 54, 99})

}
