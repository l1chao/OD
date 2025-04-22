package main

import "math"

func isValidBST(root *TreeNode) bool {

	last := math.MinInt
	ans := true
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil || ans == false {
			return
		}
		inorder(root.Left)
		if last < root.Val {
			last = root.Val
		} else {
			ans = false
		}
		inorder(root.Right)
	}
	inorder(root)
	return ans
}
