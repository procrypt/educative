package main

import (
	linkedlist "algorithms-in-golang/educative/LinkedList"
	"fmt"
)

func removeDuplicates(l *linkedlist.List)  {
	current := l.Head
	for current != nil {
		if current.Next != nil {
			if current.Key == current.Next.Key  {
				current.Next = current.Next.Next
				l.Size--
			}
		}
		current = current.Next
	}
}

func main() {
	l := &linkedlist.List{}
	values := []int{1,2,2,3,3,4,4,5,6}
	for _, value := range values {
		l.Append(value)
	}
	current := l.Head
	removeDuplicates(l)
	for current != nil {
		fmt.Println(current.Key)
		current = current.Next
	}
}
