package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := &[]int{}
	PreOrder(root, res)
	return *res
}

func PreOrder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	*res = append(*res, root.Val)
	PreOrder(root.Left, res)
	PreOrder(root.Right, res)
}
