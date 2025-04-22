package main

// https://leetcode.cn/problems/maximum-depth-of-binary-tree/?envType=problem-list-v2&envId=depth-first-search
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
