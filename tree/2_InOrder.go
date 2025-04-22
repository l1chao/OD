package main

import "fmt"

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

// 中序遍历 非递归版本。任务：按照中序遍历返回一个数组
func inorder_NoRecursion(root *TreeNode) {
	stack := []*TreeNode{}
	p := root

	for p != nil || len(stack) > 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			visit(p)
			p = p.Right
		}
	}
}

// 仅用于模拟对特定树节点操作
func visit(node *TreeNode) {
	fmt.Println(node)
}
