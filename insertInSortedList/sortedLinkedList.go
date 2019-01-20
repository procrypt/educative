package main

import (
	linkedList "algorithms-in-golang/educative/LinkedList"
	"fmt"
)

func sortedInsert(l *linkedList.List, key int) *linkedList.List {
	newNode := &linkedList.Node{Key:key}
	current := l.Head
	prev := &linkedList.Node{}

	for current != nil {
		if current.Key > key {
			prev.Next = newNode
			newNode.Next = current
			break
		}
		if current.Key < key && current.Next == nil {
			l.Tail.Next = newNode
			l.Tail = newNode
			break
		}
		prev = current
		current = prev.Next
	}
	return l
}


func main() {
	l := &linkedList.List{}
	values := []int{1,2,3,4,5,7,8,9,20}
	for _, value := range values {
		l.Append(value)
	}
	sortedInsert(l, 100)
	current := l.Head

	for current != nil {
		fmt.Println(current.Key)
		current = current.Next
	}

}
