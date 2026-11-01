package main

import "fmt"

func plusOne(digits []int) []int {
	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}
	return append([]int{1}, digits...)

}
func main() {
	digits := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	digits1 := []int{9}
	digits2 := []int{3, 7, 9, 0, 1, 8}
	fmt.Println(plusOne(digits))
	fmt.Println(plusOne(digits1))
	fmt.Println(plusOne(digits2))

}
