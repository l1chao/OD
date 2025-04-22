package main

import "math"

func findMaxAverage(nums []int, k int) float64 {
	cnt := 0
	total := math.MinInt

	for i, v := range nums {
		cnt += v
		if i < k-1 {
			continue
		}

		total = max(total, cnt)

		cnt -= nums[i-k+1]
	}

	return float64(total) / float64(k)
}
