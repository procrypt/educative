package main

import "fmt"

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	mid := len(nums)/2
	leftHalf := mergeSort(nums[:mid])
	righthalf := mergeSort(nums[mid:])
	return mergeSortMerge(leftHalf, righthalf)
}

func mergeSortMerge(leftHalf, rightHalf []int) []int  {
	sortedArray := make([]int, len(rightHalf)+len(leftHalf))
	k,i,j := 0,0,0

	for i < len(leftHalf) && j < len(rightHalf) {
		if leftHalf[i] <= rightHalf[j] {
			sortedArray[k] = leftHalf[i]
			i++
		} else {
				sortedArray[k] = rightHalf[j]
				j++
		}
		k++
	}
	for i < len(leftHalf) {
		sortedArray[k] = leftHalf[i]
		i++
		k++
	}
	for j < len(rightHalf) {
		sortedArray[k] = rightHalf[j]
		j++
		k++
	}
	return sortedArray
}

func main() {
	fmt.Println(mergeSort([]int{9,8,7,6,5,4,3,2,1}))
}
