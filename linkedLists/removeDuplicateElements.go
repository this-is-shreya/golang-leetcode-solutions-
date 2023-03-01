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
	rHead := removeDuplicate(l)
	printLinkedList(rHead)

}
func removeDuplicate(l linkedlist) *node {

	prev := &node{}
	curr := l.head
	result := curr
	i := 0
	for curr != nil && prev != nil {
		if prev != nil && prev.data == curr.data && i > 0 {

			prev.next = curr.next
			curr = curr.next

		} else {
			prev = curr
			curr = curr.next
			i = i + 1
		}

	}
	return result
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
