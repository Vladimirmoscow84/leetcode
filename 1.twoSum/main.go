package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, v := range nums {
		needed := target - v
		if _, ok := seen[needed]; ok {
			return []int{seen[needed], i}
		}
		seen[v] = i

	}
	return nil
}

func main() {
	nums := []int{3, 6, 32, 67, 31, 9, 87, 45}
	target := 51
	fmt.Println(twoSum(nums, target))
}
