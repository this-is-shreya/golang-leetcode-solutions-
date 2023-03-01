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
	fmt.Println(intersection(l, m))
}
func intersection(l linkedlist, m linkedlist) *node {

	first := l.head
	second := m.head

	for first != second {
		if first == nil {
			first = m.head
		} else {
			first = first.next
		}
		if second == nil {
			second = l.head
		} else {
			second = second.next
		}
	}
	return first
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
