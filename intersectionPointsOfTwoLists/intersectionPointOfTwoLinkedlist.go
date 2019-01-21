/*
1) Find lengths of both linked lists: L1 and L2
2) Calculate difference in length of both linked lists: d = |L1 - L2|
3) Move head pointer of longer list 'd' steps forward
4) Now traverse both lists, comparing nodes until we find a match or reach the end of lists
*/

package main

import (
	linkedList "algorithms-in-golang/educative/LinkedList"
	"math"
)

func intersectionPoint(l1, l2 *linkedList.List) *linkedList.Node {
	currentL1 := l1.Head
	currentL2 := l2.Head
	lenOfl1 := lengthOfLinkedList(l1)
	lenOfl2 := lengthOfLinkedList(l2)
	diff := int(math.Abs(float64(lenOfl1)- float64(lenOfl2)))
	counter := 0

	if lenOfl1 > lenOfl2 {
		for counter < diff {
			currentL1 = currentL1.Next
			counter++
		}
		for currentL1 != nil {
			if currentL1 == currentL2 {
				return currentL1
				break
			}
			currentL1 = currentL1.Next
			currentL2 = currentL2.Next
		}
	} else {
		for counter < diff {
			currentL2 = currentL2.Next
			counter++
		}
		for currentL2 != nil {
			if currentL2 == currentL1 {
				return currentL2
				break
			}
			currentL2 = currentL2.Next
			currentL1 =currentL1.Next
		}
	}
	return nil
}

func lengthOfLinkedList(l *linkedList.List) int {
	length := 0
	current := l.Head
	for current != nil {
		length++
		current = current.Next
	}
	return length
}


func main() {
	l :=&linkedList.List{}
	values := []int{1,2,3,4,5,6,7,8,9}
	for _, value := range values {
		l.Append(value)
	}
}
