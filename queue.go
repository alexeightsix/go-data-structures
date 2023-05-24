package main

type queue struct {
	head int
	tail int
	data []int
}

const ERROR_QUEUE_EMPTY = "QUEUE_EMPTY"

func NewQueue() queue {
	q := queue{tail: -1, head: -1, data: []int{}}
	return q
}

func (self *queue) isEmpty() bool {
	return self.head == -1
}

func (self *queue) peek() int {
	return self.data[self.head]
}

func (self *queue) enqueue(value int) {
	if self.head == -1 {
		self.head++
	}

	self.tail++
	self.data = append(self.data, value)
}

func (self *queue) dequeue() int {
	if self.isEmpty() {
		panic(ERROR_QUEUE_EMPTY)
	}

	tmp := self.peek()
	self.data = self.data[1 : self.tail+1]
	self.tail--

	if self.tail == -1 {
		self.head = -1
	}
	return tmp 
}
