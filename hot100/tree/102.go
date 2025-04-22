package main

// https://leetcode.cn/problems/binary-tree-level-order-traversal/?envType=study-plan-v2&envId=top-100-liked
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		currentLen := len(queue) // 每次取queue中剩余元素个数那么多，这正好是一层的数量。
		levelArr := make([]int, 0, currentLen)
		for range currentLen {
			out := queue[0]
			queue = queue[1:]
			levelArr = append(levelArr, out.Val)

			if out.Left != nil {
				queue = append(queue, out.Left)
			}
			if out.Right != nil {
				queue = append(queue, out.Right)
			}
		}
		ans = append(ans, levelArr)
	}
	return ans
}
