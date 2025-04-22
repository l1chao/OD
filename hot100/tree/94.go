package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var f func(*TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}

		f(root.Left)
		ans = append(ans, root.Val)
		f(root.Right)
	}
	f(root)
	return ans
}

func inorderNoRecursion(root *TreeNode) []int {
	ans := []int{}

	stack := []*TreeNode{root}
	var p *TreeNode
	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = append(ans, p.Val)
			p = p.Right
		}
	}
	return ans
}
