package main

import "fmt"

func removeDuplicates(nums []int) int {
	//так как первый всегда уникальный то наинаем со следующего
	idx := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx

}
func main() {
	nums := []int{2, 2, 2, 3, 4, 4, 4, 5, 6, 7, 7, 8, 9, 10, 10}
	answer := removeDuplicates(nums)
	fmt.Println(answer)
}
