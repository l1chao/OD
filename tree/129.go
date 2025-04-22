package main

import (
	"strconv"
)

// https://leetcode.cn/problems/sum-root-to-leaf-numbers

func sumNumbers(root *TreeNode) int {
	resultsStr := []string{} // 有了这个数组，最后遍历一遍即可得到答案

	var dfs func(*TreeNode, string)
	dfs = func(root *TreeNode, b string) {
		if root == nil { // 基础终止条件
			return
		}
		b = b + strconv.Itoa(root.Val) // 内变量沿一条根叶路径记录沿途变量

		if root.Left == nil && root.Right == nil { // 符合题目的终止条件
			resultsStr = append(resultsStr, b)
			return
		}
		dfs(root.Left, b)
		dfs(root.Right, b)
	}
	dfs(root, "")

	ans := 0
	for _, v := range resultsStr {
		if tmp, err := strconv.Atoi(v); err == nil {
			ans += tmp
		}
	}
	return ans
}
