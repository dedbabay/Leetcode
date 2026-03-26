package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	cp := []byte{}

	fs := strs[0]
	for fsi := 0; fsi < len(fs); fsi++ {
		for i := 1; i < len(strs); i++ {
			if fsi > len(strs[i])-1 || fs[fsi] != strs[i][fsi] {
				return string(cp)
			}
		}
		cp = append(cp, fs[fsi])
	}
	return string(cp)
}

func main() {
	s1 := []string{"flower", "flow", "flossome"}
	s2 := []string{"peter", "bruno"}

	fmt.Println(longestCommonPrefix(s1))
	fmt.Println(longestCommonPrefix(s2))
}
