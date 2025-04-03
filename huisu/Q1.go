package main

func permute(nums []int) [][]int { // nums:[2,4,5,6]，要找到nums的全排列
	res := [][]int{} // 每一个里层数组就是一个排列
	backtrackPermute(nums, []int{}, &res)
	return res
}

func backtrackPermute(remaining []int, path []int, res *[][]int) {
	// 终止条件：路径长度等于原始数组长度
	if len(path) == len(remaining) {
		newPath := make([]int, len(path))
		copy(newPath, path) // 将新的数组添加到结果res中，避免内存共享。
		*res = append(*res, newPath)
		return
	}

	// 遍历所有可能的选项
	for i := range remaining {
		exists := false
		for _, v := range path {
			if v == remaining[i] {
				exists = true
				break
			}
		}
		if exists {
			continue
		}

		path = append(path, remaining[i])
		backtrackPermute(remaining, path, res)
		path = path[:len(path)-1]
	}
}
