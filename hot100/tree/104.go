package main

func maxDepth(root *TreeNode) int {
	ans := 0

	var f func(*TreeNode) int
	f = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		return max(f(root.Left), f(root.Right)) + 1
	}
	return ans
}
