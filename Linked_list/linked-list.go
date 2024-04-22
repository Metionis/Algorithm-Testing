package main

import (
	"fmt"
)

type Node struct {
	value interface{}
	next  *Node
}

type LinkedList struct {
	head *Node
}

func NewNode(value interface{}, next *Node) *Node {
	return &Node{value: value, next: next}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{head: nil}
}

func (ll *LinkedList) InsertAtBeginning(data interface{}) {
	node := NewNode(data, ll.head)
	ll.head = node
}

func (ll *LinkedList) PrintList() {
	if ll.head == nil {
		fmt.Println("Linked List is empty")
		return
	}

	itr := ll.head
	listStr := ""
	for itr != nil {
		listStr += fmt.Sprintf("%v--", itr.value)
		itr = itr.next
	}
	fmt.Println(listStr)
}

func (ll *LinkedList) InsertAtEnd(data interface{}) {
	if ll.head == nil {
		ll.head = NewNode(data, nil)
		return
	}

	itr := ll.head
	for itr.next != nil {
		itr = itr.next
	}

	itr.next = NewNode(data, nil)
}

func (ll *LinkedList) InsertValues(dataList []interface{}) {
	ll.head = nil
	for _, data := range dataList {
		ll.InsertAtEnd(data)
	}
}

func (ll *LinkedList) GetLength() int {
	count := 0
	itr := ll.head
	for itr != nil {
		count++
		itr = itr.next
	}
	return count
}

func (ll *LinkedList) RemoveAt(index int) {
	if index < 0 || index >= ll.GetLength() {
		panic("Invalid value of index")
	}

	if index == 0 {
		ll.head = ll.head.next
		return
	}

	count := 0
	itr := ll.head
	for itr != nil {
		if count == index-1 {
			itr.next = itr.next.next
			break
		}
		count++
		itr = itr.next
	}
}

func (ll *LinkedList) InsertAt(index int, data interface{}) {
	if index < 0 || index > ll.GetLength() {
		panic("Invalid value of index")
	}

	if index == 0 {
		ll.InsertAtBeginning(data)
		return
	}

	count := 0
	itr := ll.head
	for itr != nil {
		if count == index-1 {
			node := NewNode(data, itr.next)
			itr.next = node
			return
		}
		count++
		itr = itr.next
	}
}

func main() {
	li := NewLinkedList()
	li.InsertAtBeginning(1)
	li.InsertAtBeginning(2)
	li.InsertAtBeginning(3)
	li.InsertAtBeginning(4)
	li.InsertAtEnd(9)
	li.InsertAtBeginning(0)
	li.PrintList()

	li2 := NewLinkedList()
	dataList := []interface{}{"Orange", "water", "cocaca", "banana"}
	li2.InsertValues(dataList)
	li2.PrintList()
	fmt.Println("li2 length:", li2.GetLength())
	li2.RemoveAt(2)
	fmt.Println("li2 after remove node")
	li2.PrintList()
	fmt.Println("li2 after insert new value to a specific index")
	li2.InsertAt(1, "cheese")
	li2.PrintList()
}
