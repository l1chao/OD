package main

import "math"

func goodNodes(root *TreeNode) int {
	total := 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, maxN int) {
		if root == nil {
			return
		}

		if root.Val >= maxN {
			total++
			dfs(root.Left, root.Val)
			dfs(root.Right, root.Val)
		} else {
			dfs(root.Left, maxN)
			dfs(root.Right, maxN)
		}
	}
	dfs(root, math.MinInt)
	return total
}
