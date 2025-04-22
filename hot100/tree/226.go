package main

func invertTree(root *TreeNode) *TreeNode {

	var f func(*TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}

		tmp := root.Left
		root.Left = root.Right
		root.Right = tmp
		f(root.Left)
		f(root.Right)
	}
	f(root)
	return root
}
