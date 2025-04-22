package main

import (
	"container/list"
	"fmt"
)

// （二叉树BFS）按BFS“逐个”访问元素
func levelOrder_OneByOne(root *TreeNode) {
	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		p := queue.Front()
		queue.Remove(p)
		node := p.Value.(*TreeNode)
		fmt.Println(node.Val)

		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
	}
}

// (二叉树BFS)按层输出元素
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}

	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		tempLen := queue.Len()

		part := []int{}
		for range tempLen {
			p := queue.Front()
			queue.Remove(p)
			node := p.Value.(*TreeNode)
			part = append(part, node.Val)

			if node.Left != nil {
				queue.PushBack(node)
			}
			if node.Right != nil {
				queue.PushBack(node)
			}
		}
		res = append(res, part)
	}
	return res
}
