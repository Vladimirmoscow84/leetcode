package main

import "fmt"

func lengthOfLastWord(s string) int {
	length := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && length == 0 {
			continue
		} else if s[i] != ' ' {
			length++
		} else {
			break
		}
	}
	return length
}

func main() {
	s := "Hello World   "
	fmt.Println(lengthOfLastWord(s))
}
