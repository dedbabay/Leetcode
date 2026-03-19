package main

import "fmt"

func removeDuplicates(nums []int) int {

	k := 2
	d := 0
	m := nums[0]
	for j := 1; j < len(nums); j++ {
		if nums[j] <= nums[j-1] {
			for i := k; i < len(nums); i++ {
				if nums[i] > nums[j-1] {
					nums[j] = nums[i]
					fmt.Println(nums, i, j)
					k = i + 1
					break
				}
			}
		}
		if nums[j] > m {
			m = nums[j]
			d++
		}
	}
	return d + 1
}

func removeDuplicates2(nums []int) int {
	cur := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[cur] = nums[i]
			cur++
		}
	}
	return cur
}

func main() {
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 3, 3, 4, 4, 4, 5, 6}
	//a := []int{0, 1, 2, 3}
	fmt.Println(a)
	n := removeDuplicates2(a)
	fmt.Println(a, "(n=", n, ")")

}
