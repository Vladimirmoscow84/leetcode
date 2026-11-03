package main

import "fmt"

/*
Свойства XOR (^):
a ^ a = 0 (одинаковые числа обнуляют друг друга)
a ^ 0 = a (ноль не влияет на результат)
a ^ b ^ a = (a ^ a) ^ b = 0 ^ b = b (порядок не важен)
Пример с массивом [4, 1, 2, 1, 2]:
result = 0 ^ 4 = 4
result = 4 ^ 1 = 5
result = 5 ^ 2 = 7
result = 7 ^ 1 = 6 (так как 7 это 4 ^ 1 ^ 2, а 1 ^ 1 дает 0)
result = 6 ^ 2 = 4 (так как 2 ^ 2 дает 0, остаётся только 4)
В итоге все парные числа «гасят» друг друга, и остаётся только то, которое без пары.
*/

func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res = res ^ v
	}
	return res
}
func main() {
	nums := []int{2, 3, 4, 5, 6, 7, 8, 9, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(singleNumber(nums))
}
