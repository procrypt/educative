package main

import (
	linkedList "algorithms-in-golang/educative/LinkedList"
	"fmt"
)

func reveseEven(l *linkedList.List) *linkedList.Node {
	current := l.Head
	if current == nil {
		return nil
	}
	odd := current
	even := current.Next
	evenFirst := even

	for {
		if odd == nil || even == nil || even.Next == nil {
			odd.Next = evenFirst
			break
		}
		odd.Next = even.Next
		odd = even.Next

		if odd.Next == nil {
			even.Next = nil
			odd.Next = evenFirst
			break
		}
		even.Next = odd.Next
		even = odd.Next
	}
	return l.Head
}


func main() {
	l := &linkedList.List{}
	values := []int{1,2,3,4,5,6,7,8,9}
	for _, value := range values {
		l.Append(value)
	}
	reveseEven(l)
	current := l.Head
	for current != nil {
		fmt.Println(current.Key)
		current = current.Next
	}
}
