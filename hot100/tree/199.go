package main

func rightSideView(root *TreeNode) []int {
	ans := []int{}

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(ans) == level {
			ans = append(ans, root.Val)
		}
		dfs(root.Right, level+1)
		dfs(root.Left, level+1)
	}
	dfs(root, 0)
	return ans
}
