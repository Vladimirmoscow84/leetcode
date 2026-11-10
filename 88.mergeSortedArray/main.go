package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
	}
	for i := m; i < n+m; i++ {
		nums1[i] = nums2[i-m]
	}
	sort.Ints(nums1)
}
func main() {
	nums1 := []int{3, 6, 8, 90, 0, 0, 0, 0}
	m := 4
	nums2 := []int{1, 2, 5, 99}
	n := 4
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}
