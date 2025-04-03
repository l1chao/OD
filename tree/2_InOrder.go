package main

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inorder func(*TreeNode) // 函数里面声明递归函数，要将声明和初始化分开写。
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}

	inorder(root)

	return res
}
