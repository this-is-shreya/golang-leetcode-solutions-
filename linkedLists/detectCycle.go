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
	fmt.Println(detectCycle(l))

}
func detectCycle(l linkedlist) bool {

	slow := l.head
	fast := l.head

	if fast == nil || fast.next == nil {
		return false
	}
	fast = fast.next.next

	for (fast != nil && fast.next != nil) && fast != slow {
		fast = fast.next.next
		slow = slow.next
	}

	if fast != nil && fast == slow {
		return true
	}
	return false
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
