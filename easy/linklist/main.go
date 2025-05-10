package main

import "fmt"

type node struct {
	value int
	next  *node
}

type link_list struct {
	head *node
	tail *node
}

func (ll *link_list) add(value int) {
	newNode := &node{value: value}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}
}

func (ll *link_list) display() {
	for node := ll.head; node != nil; node = node.next {
		fmt.Print(node.value, " ")
	}
	fmt.Println()
}

func main() {
	list := link_list{}
	list.add(1)
	list.add(2)
	list.add(3)
	list.display()
}
