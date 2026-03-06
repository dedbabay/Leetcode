package main

import "fmt"

// https://leetcode.com/problems/two-sum/
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

func main() {
  a := []int{ 1, 3, 5, 6 }
  t := 8
  fmt.Println ( twoSum ( a, t ) )
}
