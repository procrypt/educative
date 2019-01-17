package main

import "fmt"

func binarSearch(array []int, target int) int {
	left := 0
	right := len(array)-1

	for left <= right {
		mid := (left + right)/2
		if array[mid] == target {
			return mid
		}
		if array[mid] > target {
			right = mid -1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(binarSearch([]int{10,23,34,44,54,67,90}, 100)) //-1
}
