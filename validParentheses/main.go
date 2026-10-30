package main

import "fmt"

func isvalid(s string) bool {
	stack := make([]byte, 0, len(s))

	parentheses := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(', '{', '[':
			stack = append(stack, s[i])
		case ')', '}', ']':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if parentheses[s[i]] != last {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {
	s1 := "()"
	s2 := "()[]{}"
	s3 := "([)]"
	fmt.Println(isvalid(s1))
	fmt.Println(isvalid(s2))
	fmt.Println(isvalid(s3))
}
