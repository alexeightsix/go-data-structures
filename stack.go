package main

import (
	"errors"
)

type stack struct {
	pointer int
	data    []int
	length  int
}

const ERROR_STACK_OVERFLOW = "STACK_OVERFLOW"

func (self *stack) isEmpty() bool {
	return self.pointer == -1
}

func (self *stack) isFull() bool {
	if self.length == 0 {
		return false
	}
	return self.pointer+1 == self.length
}

func (self *stack) peek() int {
	return self.data[self.pointer]
}

func (self *stack) push(value int) error {
	if self.length > 0 && self.pointer+1 == self.length {
		return errors.New(ERROR_STACK_OVERFLOW)
	}

	self.data = append(self.data, value)
	self.pointer++
	return nil
}

func (self *stack) pop() int {
	value := self.data[self.pointer]

	if self.pointer == 0 {
		self.resetData()
	} else {
		self.data = self.data[:self.pointer]
	}

	self.pointer--

	return value
}

func (self *stack) resetData() {
	self.data = []int{}
}

func NewStackWithLength(length int) stack {
	s := stack{pointer: -1, length: length}
	s.resetData()
	return s
}

func NewStack() stack {
	s := stack{pointer: -1, length: 0}
	s.resetData()
	return s
}
