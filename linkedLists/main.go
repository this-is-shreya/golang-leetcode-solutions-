package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type Linkedlist struct {
	head *Node
	tail *Node
}

func main() {

	l := Linkedlist{}

	for {
		fmt.Printf("Do you want to insert a Node?(y/n)\n")
		var ans string
		fmt.Scan(&ans)
		if ans == "y" {
			fmt.Printf("Enter a value:\n")
			var ele int
			fmt.Scan(&ele)
			new_node := &Node{
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
		} else {
			printLinkedList(l)
			break
		}

	}
	fmt.Println("\nwhich node do you want to delete?")
	var ele int
	fmt.Scan(&ele)
	deleteNode(ele, l)

}
func printLinkedList(l Linkedlist) {

	temp := l.head
	for temp != nil {
		fmt.Printf("%v ", temp.data)
		temp = temp.next
	}
	fmt.Println()
}

func reverseLinkedList(l Linkedlist) {
	first := l.head
	if first.next != nil {
		second := first.next
		third := second.next

		first.next = nil
		for {
			second.next = first
			first = second
			if third == nil {
				l.head = second
				break
			}
			second = third

			third = third.next

		}
	}
	printLinkedList(l)
}
func deleteNode(n int, l Linkedlist) {
	curr := l.head
	prev := l.head
	for curr != nil {
		if curr.data == n {
			if curr == l.head {
				l.head = curr.next
			}
			prev.next = curr.next
			curr.next = nil
			break
		}
		prev = curr
		curr = curr.next
	}
	printLinkedList(l)
}
