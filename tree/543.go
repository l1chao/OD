package main

func diameterOfBinaryTree(root *TreeNode) int {
	max := 0

	var helper func(*TreeNode) int
	helper = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left := helper(root.Left)
		right := helper(root.Right)

		if left+right > max {
			max = left + right
		}

		return maxInt(left, right) + 1
	}
	helper(root)
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
