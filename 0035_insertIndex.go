package main

import "fmt"

func searchInsert(nums []int, target int) int {
	x := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
		if nums[i] > target {
			return x
		}
		x++
	}
	return len(nums)
}

func binarySearchInsert(nums []int, target int) int {

	l := 0
	r := len(nums) - 1

	for l <= r {
		m := l + (r-l)/2
		if target == nums[m] {
			return m
		} else if target < nums[m] {
			r = m - 1
		} else if target > nums[m] {
			l = m + 1
		}
	}
	return l
}

func main() {
	a := []int{1, 3, 5, 7, 9}
	t := 60

	fmt.Println(searchInsert(a, t))
	fmt.Println(binarySearchInsert(a, t))
}
