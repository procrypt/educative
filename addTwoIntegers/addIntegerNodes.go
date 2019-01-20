package main

import (
	linkedList "algorithms-in-golang/educative/LinkedList"
	"fmt"
)

func sumList(l1,l2 *linkedList.List) linkedList.List {
	list := linkedList.List{}
	p, q := l1.Head, l2.Head
	i, j, carry := 0,0,0

	for q != nil || q != nil {
		if p == nil {
			i = 0
		} else {
			i = p.Key
			p = p.Next
		}
		if q == nil {
			j = 0
		} else {
			j = q.Key
			q = q.Next
		}

		sum := i+j+carry

		if sum >= 10 {
			carry = sum/10
			list.Append(sum%10)
		} else {
			carry = 0
			list.Append(sum)
		}
	}
	return list
}


func main() {
	l1 := &linkedList.List{}
	values1 := []int{7,1,6}
	for _, value := range values1 {
		l1.Append(value)
	}
	l2 := &linkedList.List{}
	values2 := []int{5,9,2}
	for _, value := range values2{
		l2.Append(value)
	}
	a := sumList(l1,l2)
	for a.Head != nil {
		fmt.Println(a.Head.Key)
		a.Head = a.Head.Next
	}
}
