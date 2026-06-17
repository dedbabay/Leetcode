package main

import "fmt"

func lengthOfLastWord(s string) int {
	l := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			l++
		} else if l > 0 {
			break
		}
	}
	return l
}

func main() {
	s := "   fly me   to   the moon  "
	fmt.Println(lengthOfLastWord(s))
}
