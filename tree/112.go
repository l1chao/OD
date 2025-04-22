package main

// 思路和111求二叉树最小深度是一样的
func hasPathSum(root *TreeNode, targetSum int) bool {
	ans := false

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, cnt int) {
		if root == nil {
			return
		}

		cnt += root.Val

		if ans {
			return
		}

		if root.Left == nil && root.Right == nil && cnt == targetSum {
			ans = true
			return
		}

		dfs(root.Left, cnt)
		dfs(root.Right, cnt)
	}
	dfs(root, 0)
	return ans
}
