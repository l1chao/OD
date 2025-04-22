package main

import "fmt"

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	stack := []*TreeNode{root}
	var prev *TreeNode
	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if prev != nil {
			prev.Left = nil
			prev.Right = curr
		}
		if curr.Right != nil {
			stack = append(stack, curr.Right)
		}
		if curr.Left != nil {
			stack = append(stack, curr.Left)
		}

		prev = curr
	}
}

// 复习 迭代前序遍历

func preOrderNoRecursion(root *TreeNode) {
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) > 0 {
		if p != nil {
			visit(p)
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1].Right
			stack = stack[:len(stack)-1]
		}
	}
}

// 顺便复习中序遍历
func inOrderNoRecursion(root *TreeNode) {
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

func visit(root *TreeNode) {
	fmt.Println(root.Val)
}
