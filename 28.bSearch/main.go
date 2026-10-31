package main

import "fmt"

func searcInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		switch {
		case nums[mid] == target:
			return mid
		case nums[mid] < target:
			left = mid + 1
		default:
			right = mid - 1
		}

	}
	return left
}
func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searcInsert(nums, target))
}
