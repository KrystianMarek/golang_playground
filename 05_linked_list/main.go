package main

import (
	"fmt"
)

type MyType interface {
	int | string
}

type Node[V MyType] struct {
	next     *Node[V]
	previous *Node[V]
	value    V
}

func (node *Node[V]) SetValue(value V) {
	node.value = value
}

func (node *Node[V]) GetValue() V {
	return node.value
}

type LinkedList[V MyType] struct {
	head *Node[V]
	tail *Node[V]
}

func (list *LinkedList[V]) Add(val V) {
	newNode := &Node[V]{value: val}
	if list.head == nil {
		list.head = newNode
	} else if list.tail == list.head {
		list.head.next = newNode
	} else if list.tail != nil {
		list.tail.next = newNode
	}
	list.tail = newNode
}

func (list *LinkedList[V]) String() string {
	output := ""
	for n := list.head; n != nil; n = n.next {
		output += fmt.Sprintf(" {%v} ", n.GetValue())
	}
	return output
}

func main() {
	listOfInt := new(LinkedList[int])
	listOfInt.Add(1)
	listOfInt.Add(2)
	listOfInt.Add(3)
	fmt.Println(listOfInt)

	listOfString := new(LinkedList[string])
	listOfString.Add("one")
	listOfString.Add("two")
	listOfString.Add("three")
	fmt.Println(listOfString)
}
