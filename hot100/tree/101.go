package main

func isSymmetric(root *TreeNode) bool {
	reverse(root.Right)
	return isSameTree(root.Left, root.Right)
}

func reverse(root *TreeNode) {
	if root == nil {
		return
	}

	tmp := root.Left
	root.Left = root.Right
	root.Right = tmp
	reverse(root.Left)
	reverse(root.Right)
}

// 如何检测两个树的值是一样的？
// https://leetcode.cn/problems/same-tree/
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q // 这一步省略了好几步，注意细细分析
	}

	return p.Val == q.Val && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
