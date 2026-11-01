package main

import "fmt"

func addBinary(a, b string) string {
	result := []byte{}
	carry := 0
	i, j := len(a)-1, len(b)-1

	for i >= 0 || j >= 0 || carry > 0 {
		digitA := 0
		if i >= 0 {
			digitA = int(a[i] - '0')
		}
		digitB := 0
		if j >= 0 {
			digitB = int(b[j] - '0')
		}
		total := digitA + digitB + carry
		result = append(result, byte(total%2)+'0')
		carry = total / 2
		i--
		j--
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return string(result)
}
func main() {
	a := "101010"
	b := "01001"
	fmt.Println(addBinary(a, b))
}
