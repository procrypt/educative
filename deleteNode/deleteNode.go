package main

import (
	linkedList "algorithms-in-golang/educative/LinkedList"
	"fmt"
)

func deleteNode(l *linkedList.List, key int) {
	current := l.Head
	for current != nil {
		if current.Key == key {
			current.Key = current.Next.Key
			current.Next = current.Next.Next
		}
		current = current.Next
	}
}


func main() {
	l := &linkedList.List{}
	values := []int{1,2,3,4,5,6,7,8,9}
	for _, value := range values {
		l.Append(value)
	}
	current :=  l.Head
	deleteNode(l, 4)
	fmt.Println("---")
	for current != nil {
		fmt.Println(current.Key)
		current = current.Next
	}
}
