package main

import (
	"fmt"
)

type MyType interface {
	int | string
}

type Node[V MyType] struct {
	next             *Node[V]
	previous         *Node[V]
	constrainedValue V
	anyValue         interface{}
}

func (node *Node[V]) SetConstrainedValue(value V) {
	node.constrainedValue = value
}

func (node *Node[V]) GetConstrainedValue() V {
	return node.constrainedValue
}

func (node *Node[V]) SetAnyValue(value interface{}) {
	node.anyValue = value
}

func (node *Node[V]) GetAnyValue() interface{} {
	return node.anyValue
}

type LinkedList[V MyType] struct {
	head *Node[V]
	tail *Node[V]
}

func (list *LinkedList[V]) Add(constrained V, any interface{}) {
	newNode := &Node[V]{constrainedValue: constrained, anyValue: any}
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
		output += fmt.Sprintf(" {%v} | {%v}\n", n.GetConstrainedValue(), n.GetAnyValue())
	}
	return output
}

func main() {
	listOfInt := new(LinkedList[int])
	listOfInt.Add(1, 1)
	listOfInt.Add(2, "two")
	listOfInt.Add(3, 3.1)
	fmt.Println(listOfInt)

	listOfString := new(LinkedList[string])
	listOfString.Add("one", [2]int{5, 4})
	listOfString.Add("two", 2.4)
	listOfString.Add("three", 3.14)
	fmt.Println(listOfString)
}
