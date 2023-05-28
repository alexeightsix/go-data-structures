package main

import "fmt"

func main() {

	linkedlist := NewLinkedList()

	a := NewNode("A")
  b := NewNode("B")
  c := NewNode("C")
  d := NewNode("D")


	linkedlist.insertStart(&a)
	linkedlist.insertStart(&b)
	linkedlist.insertStart(&c)
	linkedlist.insertAfter(&a, &d)

  fmt.Println(linkedlist.tail)

	// c , b , a , d

}
