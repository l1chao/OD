package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归写法1
func preorderTraversal(root *TreeNode) []int {
	res := &[]int{}
	PreOrder(root, res) // 这种写法是将递归的部分写到外面去了。
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

// 递归写法2
func preorderTraversal2(root *TreeNode) []int {
	res := []int{}
	var preorder func(*TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		// visit(root)
		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}
	return res
}

// 迭代写法
func preorderTraversal_NoRecursion(root *TreeNode) (res []int) {
	p := root
	stack := []*TreeNode{}

	for p != nil || len(stack) > 0 {
		if p != nil {
			// visit(p)
			res = append(res, p.Val)
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
	return
}
