package main

func searchMatrix(matrix [][]int, target int) bool {
	if target < matrix[0][0] {
		return false
	}

	head, tail := 0, len(matrix)-1
	for head <= tail {
		mid := (head + tail) / 2
		if matrix[mid][0] == target {
			return true
		} else if matrix[mid][0] < target {
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}

	targetLine := head - 1
	for head, tail = 0, len(matrix[0])-1; head <= tail; {
		mid := (head + tail) / 2
		if matrix[targetLine][mid] == target {
			return true
		} else if matrix[targetLine][mid] < target {
			head = mid + 1
		} else {
			tail = mid - 1
		}
	}
	return false
}
