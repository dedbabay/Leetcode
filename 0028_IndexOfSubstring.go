package main

import "fmt"

func strStr(haystack string, needle string) int {

	if len(needle) > len(haystack) ||
		len(haystack) == 0 || len(needle) == 0 {
		return -1
	}

	for x := 0; x < len(haystack) - len(needle) + 1; x++ {
		eq := true
		for i := x; i < x+len(needle); i++ {
			if haystack[i] != needle[i-x] {
				eq = false
				break
			}
		}
		if eq {
			return x
		}
	}
	return -1
}

func strStr1(haystack string, needle string) int {

	lh := len(haystack)
	ln := len(needle)
	for i := 0; i < lh-ln; i++ {
		leq := 0
		for j := i; j < ln+i; j++ {
			if haystack[j] != needle[j-i] {
				break
			}
			leq++
		}
		if leq == ln {
			return i
		}
	}
	return -1
}

func main() {
	s1 := "sadbuttrueandits sad"
	s2 := "true"
	fmt.Println(strStr1(s1, s2))

}
