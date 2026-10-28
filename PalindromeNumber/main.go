package main

import "fmt"

// withArray
func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	count := 0
	y := x
	for y > 0 {
		y /= 10
		count++
	}
	nums := make([]int, count)
	for i := 0; i < len(nums); i++ {
		nums[i] = x % 10
		x /= 10
	}
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		if nums[i] != nums[j] {
			return false
		}
	}
	return true
}

// memory(O1) - best
func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	half := 0
	for x > half {
		half = half*10 + x%10
		x /= 10
	}
	return x == half || x == half/10
}
func main() {
	x := 121
	fmt.Println(isPalindrome1(x), isPalindrome(x))

}
