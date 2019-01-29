package main

import "fmt"

type Interval struct {
	Start int
	End   int
}

func quickSort(interval []Interval) []Interval  {
	return helper(interval, 0, len(interval)-1)
}

func helper(array []Interval, startIdx, endIdx int) []Interval {
	if startIdx >= endIdx {
		return array
	}
	pivot := startIdx
	leftIdx := startIdx+1
	rightIdx := endIdx

	for leftIdx <= rightIdx {
		if array[pivot].Start < array[leftIdx].Start && array[pivot].Start > array[rightIdx].Start	{
			array[leftIdx], array[rightIdx] = array[rightIdx], array[leftIdx]
		}
		if array[pivot].Start >= array[leftIdx].Start {
			leftIdx++
		}
		if array[pivot].Start < array[rightIdx].Start {
			rightIdx--
		}
	}
	array[pivot], array[rightIdx] = array[rightIdx], array[pivot]
	helper(array, startIdx, rightIdx-1)
	helper(array, rightIdx+1, endIdx)
	return array
}

func mergeOverlapping(intervals []Interval) []Interval {
	interval := quickSort(intervals)
	stack := []Interval{}
	stack = append(stack, interval[0])
	for i:=1; i<len(interval); i++ {
		if stack[len(stack)-1].End >= interval[i].Start {
			stack[len(stack)-1].End = max(stack[len(stack)-1].End, interval[i].End)
		} else {
			stack = append(stack, interval[i])
		}
	}
	return stack
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	d := Interval{1,4}
	c := Interval{2,3}
	b := Interval{8,10}
	a := Interval{15,18}
	interval := []Interval{d,c,b,a}
	fmt.Println(mergeOverlapping(interval))
}
