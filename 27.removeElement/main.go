package main

import "fmt"

func removeElement(nums []int, val int) int {
	idx := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}

	}
	return idx

}
func main() {
	nums := []int{1, 2, 3, 4, 4, 5, 6, 4, 7, 9, 4, 3, 4}
	target := 4
	fmt.Println(removeElement(nums, target))
}
