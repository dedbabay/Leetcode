// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
package main

import (
	"fmt"
)

func findMedianSortedArrays(a1 []int, a2 []int) float64 {
	r := []int{}
	i, j := 0, 0
	for i < len(a1) && j < len(a2) {
		if a1[i] < a2[j] {
			r = append(r, a1[i])
			i++
		} else {
			r = append(r, a2[j])
			j++
		}
	}

	r = append(r, a1[i:]...)
	r = append(r, a2[j:]...)

	left := int(len(r) / 2)
	right := left
	p := len(r) % 2

	if p == 0 {
		left -= 1
	}

	mediana := float64(r[left]+r[right]) / 2.0

	return mediana
}
