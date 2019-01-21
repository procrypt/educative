package main

import "fmt"

func lowHigh(array []int, target int) []int {
	finalRange := []int{-1,-1}
	helper(array, target, finalRange, 0, len(array)-1, true)
	helper(array, target, finalRange, 0, len(array)-1, false)
	return finalRange

}

func helper(array []int, target int, finalRange []int, left, right int, goLeft bool) {
	for left <= right {
		mid := (left+right)/2
		if array[mid] > target {
			right = mid - 1
		} else if array[mid] < target {
			left = mid + 1
		} else {
			if goLeft {
				if mid == 0 || array[mid-1] != target {
					finalRange[0] = mid
					return
				} else {
					right = mid -1
				}
			} else {
				if mid == len(array)-1 || array[mid+1] != target {
					finalRange[1] = mid
					return
				} else {
					left = mid + 1
				}
			}
		}
	}
}

func main() {
	fmt.Println(lowHigh([]int{0, 1, 21, 33, 45, 45, 45, 45, 45, 45, 61, 71, 73}, 45))
}
