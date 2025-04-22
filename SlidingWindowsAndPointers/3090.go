package main

// 输入： s = "bcbbbcba"
// 输出： 4

func maximumLengthSubstring(s string) int {
	left := 0
	cnt := make([]int, 0)
	maxNum := 0
	for right, v := range s {
		cnt[v]++
		for cnt[v] > 2 {
			cnt[s[left]]--
			left++
		}
		maxNum = max(maxNum, right-left+1)
	}
	return maxNum
}
