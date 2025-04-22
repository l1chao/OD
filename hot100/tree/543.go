package main

func diameterOfBinaryTree(root *TreeNode) int {
	maxL := 0
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		maxL = max(maxL, left+right)
		return max(left, right) + 1
	}

	return maxL
}
