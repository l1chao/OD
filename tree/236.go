package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var res *TreeNode
	var helper func(a, b, c *TreeNode) bool
	helper = func(root, p, q *TreeNode) bool {
		if root == nil {
			return false
		}

		leftBool := helper(root.Left, p, q)
		rightBool := helper(root.Right, p, q)

		if (root.Val == p.Val || root.Val == q.Val) && (leftBool || rightBool) ||
			leftBool && rightBool {
			res = root
		}

		if leftBool || rightBool {
			return true
		}
		return false
	}

	helper(root, p, q)
	return res
}
