package main

import "fmt"

// O (n^2)
func twoSum(nums []int, target int) []int {
    for i := range nums {
        for j := i + 1; j < len(nums); j++ {
            if nums[i] + nums[j] == target {
                return []int { i, j }
            }
        }
    }
    return nil
}

// O(n)
func twoSum2(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		delta := target - nums[i]
		if j, exists := m[delta]; exists {
			return []int{j, i}
		}
		m[n] = i
	}
	return nil
}

func main() {
  a := []int{ 1, 3, 5, 6 }
  t := 8
  fmt.Println ( twoSum ( a, t ) )
}
