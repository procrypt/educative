package PowerOfANumber

import "fmt"

func power(x, n int) int {
	if n == 0 {
		return 1
	} else if n < 0 {
		n = -n
		x = 1%x
		fmt.Println(x,n)
		return power(x, n)
	} else if n%2 == 0 {
		temp := power(x, n/2)
		return temp*temp
	} else {
		return x * power(x,n-1)
	}
}

func main() {
	fmt.Println(power(2,-2))
}
