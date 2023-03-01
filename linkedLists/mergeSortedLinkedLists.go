package main

import "fmt"

type node struct {
	data int
	next *node
}
type linkedlist struct {
	head *node
	tail *node
}

func main() {

	l := linkedlist{}
	m := linkedlist{}
	for {
		var ele int
		fmt.Println("Enter element:")

		fmt.Scan(&ele)
		if ele == -999 {
			printLinkedList(l.head)

			break
		}
		l = insertNode(l, ele)

	}
	for {
		var ele int
		fmt.Println("Enter element for other linked list:")

		fmt.Scan(&ele)
		if ele == -999 {
			printLinkedList(m.head)

			break
		}
		m = insertNode(m, ele)

	}
	mHead := mergeSortedArray(l, m)
	printLinkedList(mHead)
}
func mergeSortedArray(l linkedlist, m linkedlist) *node {

	curr := &node{}
	first := l.head
	second := m.head

	result := curr

	for first != nil && second != nil {
		if first.data <= second.data {
			curr.next = first
			curr = curr.next
			first = first.next
		} else {
			curr.next = second
			curr = curr.next
			second = second.next
		}
	}

	if second != nil {
		curr.next = second
	}
	if first != nil {
		curr.next = first
	}

	return result.next
}
func printLinkedList(head *node) {

	temp := head
	for temp != nil {
		fmt.Printf("%v ", temp.data)
		temp = temp.next
	}
	fmt.Println()
}
func insertNode(l linkedlist, ele int) linkedlist {
	new_node := &node{
		data: ele,
		next: nil,
	}
	if l.head != nil {
		l.tail.next = new_node
		l.tail = new_node

	} else {
		l.head = new_node
		l.tail = new_node

	}
	return l
}
