package main

import "fmt"

func lengthOfLongestSubstring(s string) int {

	m := make(map[rune]int)

	left := 0
	max_len := 0
	len := 0

	for i, c := range s {
		x, e := m[c]
		if e && x >= left {
			left = x + 1
		}
		m[c] = i
		len = i - left + 1
		if len > max_len {
			max_len = len
		}
		fmt.Println(i, string(c), x, e, left, len, max_len, m)
	}

	return max_len
}

func main() {

	s := "abba"

	fmt.Println(lengthOfLongestSubstring(s))
}
