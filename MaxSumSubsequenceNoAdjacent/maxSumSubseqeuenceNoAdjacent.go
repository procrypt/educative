package main

import "fmt"

// Dynamic Progamming

func maxSumSubsequenceNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	}
	if len(array) == 1 {
		return array[0]
	}
	first := array[0]
	second := max(array[0], array[1])
	for i:=2; i< len(array); i++ {
		current := max(second, first+array[i])
		first = second
		second = current
	}
	return second
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSumSubsequenceNoAdjacent([]int{75,105,120,75,90,135}))
}
