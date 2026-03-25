package main

import "fmt"

func generateParenthesis(n int) []string {

	res := []string{}

	var f func(string, int, int, int)
	f = func(s string, o, c, n int) {

		if o == c && o == n {
			res = append(res, s)
			return
		}

		if o < n {
			f(s+"(", o+1, c, n)
		}

		if c < o {
			f(s+")", o, c+1, n)
		}
	}
	f("", 0, 0, n)

	return res

}

func main() {
	for _, b := range generateParenthesis(4) {
		fmt.Println(b)
	}
}
