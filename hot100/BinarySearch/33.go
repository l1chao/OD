package main

func search(nums []int, target int) int {
	head, tail := 0, len(nums)-1
	mid := 0
	for head <= tail {
		mid = (head + tail) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[0] {
			if nums[mid] > target && target >= nums[head] {
				tail = mid - 1
			} else {
				head = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[tail] {
				head = mid + 1
			} else {
				tail = mid - 1
			}
		}
	}
	return -1
}
