package main

import "fmt"

func subsets(array []int) [][]int {
	sets := [][]int{}
	sets = append(sets, []int{})
	for _, element := range array {
		for _, set := range sets {
			currentSet := set
			currentSet = append(currentSet, element)
			sets = append(sets, currentSet)
		}
	}
	return sets
}

func main() {
	fmt.Println(subsets([]int{1,2,3,4}))
}

