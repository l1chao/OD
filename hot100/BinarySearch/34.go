package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	head, tail := 0, len(nums)-1
	mid := 0
	isFound := -1
	for head <= tail {
		mid = (head + tail) / 2
		if nums[mid] == target {
			isFound = mid
			break
		} else if nums[mid] < target {
			head = mid + 1
		} else {
			tail = mid - 1
		}
	}
	if isFound != -1 {
		i, j := 0, 0
		for i = isFound - 1; i >= 0 && nums[i] == nums[isFound]; i-- {
		}
		for j = isFound + 1; j < len(nums) && nums[j] == nums[isFound]; j++ {
		}
		return []int{i + 1, j - 1}
	}

	return []int{-1, -1}
}
