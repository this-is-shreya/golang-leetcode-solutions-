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
			printLinkedList(l)

			break
		}
		l = insertNode(l, ele)

	}
	// fmt.Println(l.head)
	middleElement(l)

}
func printLinkedList(l linkedlist) {

	temp := l.head
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

func middleElement(l linkedlist) {

	fast := l.head
	slow := l.head
	if fast == nil {
		fmt.Println("empty linked list,", fast)
	} else if fast.next == nil {
		fmt.Println(fast.data)
	} else {
		for fast != nil && fast.next != nil {
			fast = fast.next.next
			slow = slow.next
		}
		fmt.Println("data of the middle element: ", slow.data)
	}

}
