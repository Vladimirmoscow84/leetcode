package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	//нужно проверить на длину ноль но по условию такой длины слайса не может быть
	//if len(strs)==0{return ""}

	//берем первую строку за эталон, так как общий префикс не может быть болше нее
	for i := 0; i < len(strs[0]); i++ {
		underString := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != underString {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
func main() {

	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))

}
