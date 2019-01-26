package squareRoot

import "fmt"

func mySqrt(x int) int {
	result := 0
	if x == 0 {
		return 0
	}
	if x == 1 {
		return 1
	}
	left, right := 1, x

	for left <= right {
		mid := (left+right)/2
		if mid*mid == x {
			return mid
		}
		if mid*mid < x {
			left = mid+1
			result = mid
		} else {
			right = mid-1
		}
	}
	return result
}

func main() {
	fmt.Println(mySqrt(8))
}
