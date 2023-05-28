package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	linkedlist := NewLinkedList()

	assert.True(t, linkedlist.isEmpty())

	a := NewNode("A")
	b := NewNode("B")
	c := NewNode("C")
	d := NewNode("D")
	e := NewNode("E")

	linkedlist.insertStart(&a)
	linkedlist.insertStart(&b)
	linkedlist.insertStart(&c)
	linkedlist.insertAfter(&c, &d)
	linkedlist.insertEnd(&e)

	assert.False(t, linkedlist.isEmpty())

	assert.Nil(t, linkedlist.tail.next)

	assert.Equal(t, linkedlist.head.data, "C")
	assert.Equal(t, linkedlist.head.next.data, "D")
	assert.Equal(t, linkedlist.head.next.next.data, "B")
	assert.Equal(t, linkedlist.head.next.next.next.data, "A")
	assert.Equal(t, linkedlist.head.next.next.next.next.data, "E")

	assert.Equal(t, linkedlist.Search("A").data, "A")
	assert.Nil(t, linkedlist.Search("Z"))

	linkedlist.deleteStart()
	assert.Equal(t, linkedlist.head.data, "D")

	assert.Equal(t, linkedlist.Search("B").data, "B")
	linkedlist.deleteAfter(&d)
	assert.Nil(t, linkedlist.Search("B"))

	linkedlist.deleteEnd()
  assert.Equal(t, linkedlist.tail.data, "A")

}
