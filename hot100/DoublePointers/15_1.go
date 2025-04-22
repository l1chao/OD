package main

import (
	"slices"
)

// 15题是三数之和，是1题两数之和的进阶版本。先复习1题，然后写15题

// 1题
func twoSum(nums []int, target int) []int {
	cnt := make(map[int]int)
	for i, v := range nums {
		if index, ok := cnt[target-v]; ok {
			return []int{index, i}
		}
		cnt[v] = i
	}
	return nil
}

// 15题
// 输入：nums = [-1,0,1,2,-1,-4]
// 输出：[[-1,-1,2],[-1,0,1]]
func threeSum(nums []int) [][]int {
	ans := [][]int{}

	slices.Sort(nums)
	n := len(nums)
	for i, x := range nums[:n-2] {
		if i > 0 && x == nums[i-1] {
			continue
		}
		if x+nums[i+1]+nums[i+2] > 0 {
			break
		}
		if x+nums[n-1]+nums[n-2] < 0 {
			continue
		}

		j, k := i+1, n-1
		for j < k {
			s := x + nums[j] + nums[k]
			if s > 0 {
				k--
			} else if s < 0 {
				j++
			} else {
				ans = append(ans, []int{x, nums[j], nums[k]})
				for j++; j < k && nums[j] == nums[j-1]; j++ {
				}
				for k--; j < k && nums[k] == nums[k+1]; k-- {
				}
			}
		}
	}
	return ans
}
