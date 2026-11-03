package main

import "fmt"

func mySqrt(x int) int {
	left := 0
	right := x
	res := 0
	for left <= right {
		mid := left + (right-left)/2
		switch {
		case mid*mid <= x:
			left = mid + 1
			res = mid
		default:
			right = mid - 1
		}
	}
	return res
}

func main() {

	x1 := 6
	x2 := 4
	x3 := 25
	x4 := 10
	fmt.Println(mySqrt(x1), mySqrt(x2), mySqrt(x3), mySqrt(x4))
}
