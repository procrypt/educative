package main

import "fmt"

func quickSort(array []int) []int {
	return helper(array, 0, len(array)-1)
}

func helper(array []int, startIdx, endIdx int) []int {
	if startIdx >= endIdx {
		return array
	}
	pivot := startIdx
	leftPtr, rightPtr := startIdx+1, endIdx

	for leftPtr <= rightPtr {
		if array[leftPtr] > array[pivot] && array[rightPtr] < array[pivot] {
			array[leftPtr], array[rightPtr] = array[rightPtr], array[leftPtr]
		}
		if array[pivot] >= array[leftPtr] {
			leftPtr++
		}
		if array[pivot] < array[rightPtr] {
			rightPtr--
		}
	}
	array[pivot], array[rightPtr] = array[rightPtr], array[pivot]

	helper(array, startIdx, rightPtr-1)
	helper(array, rightPtr+1, endIdx)
	return array
}

func main() {
	fmt.Println(quickSort([]int{10,5,8,4,1,2,9}))
}
