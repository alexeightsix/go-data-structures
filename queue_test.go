package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue()

	assert.True(t, queue.isEmpty())

	queue.enqueue(1)
	queue.enqueue(2)
	queue.enqueue(3)

	assert.False(t, queue.isEmpty())

	assert.Equal(t, 1, queue.peek())
	assert.Equal(t, 1, queue.dequeue())

	queue.dequeue()
	queue.dequeue()

	assert.True(t, queue.isEmpty())

  assert.Panics(t, func() {queue.dequeue()}, ERROR_QUEUE_EMPTY)
}
