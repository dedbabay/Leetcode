package main

import "fmt"

func removeElement(nums []int, val int) int {
	k := 0
	cur := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			k++
		} else {
			nums[cur] = nums[i]
			cur++
		}
	}
	return len(nums) - k
}

func removeElement2(nums []int, val int) int {
	cur := 0
	for _, n := range nums {
		if n != val {
			nums[cur] = n
			cur++
		}
	}
	return cur
}

func removeElement3(nums []int, val int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		if nums[l] == val {
			nums[l] = nums[r]
			r--
		} else {
			l++
		}
	}

	return l
}

func main() {
	a := []int{3, 0, 2, 3, 3, 3, 1, 2, 3, 4, 3, 0, 5, 2, 2, 3}
	k := removeElement2(a, 3)
	fmt.Println(a, "k=", k)
	fmt.Println(a[:k])
}
