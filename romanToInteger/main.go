package main

import "fmt"

// switch
func romanToInt(s string) int {
	answer := 0
	value := func(b byte) int {
		switch b {
		case 'I':
			return 1
		case 'V':
			return 5
		case 'X':
			return 10
		case 'L':
			return 50
		case 'C':
			return 100
		case 'D':
			return 500
		case 'M':
			return 1000
		}
		return 0
	}

	for i := len(s) - 1; i >= 0; i-- {
		if i+1 < len(s) && value(s[i]) < value(s[i+1]) {
			answer -= value(s[i])
		} else {
			answer += value(s[i])
		}
	}
	return answer
}

// hash map (better)
func romanToInt1(s string) int {
	answer := 0
	nums := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && nums[s[i]] < nums[s[i+1]] {
			answer -= nums[s[i]]
		} else {
			answer += nums[s[i]]
		}
	}

	return answer
}
func main() {
	s := "MCMXCIV"
	fmt.Println(romanToInt(s), romanToInt1(s))
}
