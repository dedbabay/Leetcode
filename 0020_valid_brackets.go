package main

import (
	"fmt"
)

func isValid(s string) bool {

	a := []rune{}

	for _, c := range s {
		if c == '[' || c == '{' || c == '(' {
			a = append(a, c)
        } else if len ( a ) == 0 {
            return false
        } else {
			l := a[len(a)-1]
			if (c == ']' && l != '[') ||
				(c == ')' && l != '(') ||
				(c == '}' && l != '{') {
				return false
			}
			a = a[:len(a)-1]
		}
	}
	return len(a) == 0
}

func main() {
	s1 := "()"
	fmt.Println(s1, isValid(s1))

	s2 := "({[]})"
	fmt.Println(s2, isValid(s2))

	s3 := "([)])"
	fmt.Println(s3, isValid(s3))

}
