package main

import (
	"fmt"
	"sort"
)

func threeSum(array []int, target int) [][]int  {
	sort.Ints(array)
	result := [][]int{}
	for i:=0; i< len(array)-2; i++ {
		leftPtr, rightPtr := i+1, len(array)-1
		for leftPtr < rightPtr {
			currentSum := array[i] + array[leftPtr] + array[rightPtr]
			if currentSum == target {
				result = append(result, []int{array[i],array[leftPtr],array[rightPtr]})
				leftPtr++
				rightPtr--
			} else if currentSum > target {
				rightPtr--
			} else {
				leftPtr++
			}
		}
	}
	return result
}

func main() {
	fmt.Println(threeSum([]int{12, 3, 1, 2, -6, 5, -8, 6},0))
}