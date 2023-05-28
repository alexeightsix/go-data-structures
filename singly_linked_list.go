package main

type node struct {
	data string
	next *node
}

type linkedlist struct {
	head *node
	tail *node
}

func NewLinkedList() *linkedlist {
	return &linkedlist{}
}

func NewNode(value string) node {
	n := node{}
	n.data = value
	return n
}

func (self *linkedlist) isEmpty() bool {
	return self.head == nil
}

func (self *linkedlist) setTail(node *node) {
	self.tail = node
	self.tail.next = nil
}

func (self *linkedlist) clear() {
	self.head = nil
	self.tail = nil
}

func (self *linkedlist) insertStart(node *node) {
	if self.isEmpty() {
		self.head = node
		self.setTail(node)
	} else {
		tmp := self.head

		if self.head == self.tail {
			self.setTail(tmp)
		}

		self.head = node
		self.head.next = tmp
	}
}

func (self *linkedlist) insertAfter(prev *node, node *node) {
	if self.tail == prev {
		self.setTail(node)
	}

	if prev.next != nil {
		node.next = prev.next
	}
	prev.next = node
}

func (self *linkedlist) insertEnd(node *node) {
	if self.isEmpty() {
		self.insertStart(node)
	} else {
		tmp := self.tail
		self.setTail(node)
		tmp.next = self.tail
	}
}

func (self *linkedlist) Search(value string) *node {
	tmp := self.head

	for {
		if tmp == nil {
			break
		}

		if tmp.data == value {
			return tmp
		}

		tmp = tmp.next
	}

	return nil
}

func (self *linkedlist) deleteStart() {
	if self.head == self.tail {
		self.clear()
	}

	self.head = self.head.next
}

func (self *linkedlist) deleteAfter(node *node) {
	if node.next == self.tail {
		self.deleteEnd()
	} else {
		node.next = node.next.next
	}
}

func (self *linkedlist) deleteEnd() {
	tmp := self.head
	for {
		if tmp.next == self.tail {
			self.setTail(tmp)
			break
		}
		tmp = tmp.next
	}
}
