package main

import "slices"

// 似乎所有的二叉树遍历对于叶节点来说都是从左到右实现的。
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	vals1 := []int{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			vals1 = append(vals1, root.Val)
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root1)
	// vals2 := append([]int(nil), vals1...)
	vals2 := slices.Clone(vals1)
	vals1 = []int{} // 将前面的值抹去，因为下面会append到vals1后面，如果不抹去，len必定不一样。就这么做也行的。
	dfs(root2)

	if len(vals2) != len(vals1) {
		return false
	}

	for i, v := range vals1 {
		if v != vals2[i] {
			return false
		}
	}
	return true
}

func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor1(root.Left, p, q)
	right := lowestCommonAncestor1(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	return left
}
