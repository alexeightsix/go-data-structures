package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()

	assert.False(t, stack.isFull())
	assert.True(t, stack.isEmpty())

	stack.push(1)
	stack.push(2)
	stack.push(3)

	assert.False(t, stack.isEmpty())
	assert.False(t, stack.isFull())

	tmp := stack.pop()
	assert.Equal(t, 3, tmp)

	tmp = stack.pop()
	assert.Equal(t, 2, tmp)

	tmp = stack.pop()
	assert.Equal(t, 1, tmp)

	stackWithLength := NewStackWithLength(5)

	stackWithLength.push(1)
	stackWithLength.push(2)
	stackWithLength.push(3)
	stackWithLength.push(4)
	stackWithLength.push(5)

	assert.True(t, stackWithLength.isFull())

	err := stackWithLength.push(6)
	assert.EqualError(t, err, ERROR_STACK_OVERFLOW)
}
