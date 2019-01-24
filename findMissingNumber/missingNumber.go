package main

import "fmt"

func missingNumber(array []int) int {
	missingInt := (len(array)+1)*(len(array)+2)/2
	fmt.Println(missingInt)
	for _, num := range array {
		missingInt -= num
	}
	return missingInt
}

func main() {
	fmt.Println(missingNumber([]int{1,2,4,6,3,7,8}))
}
