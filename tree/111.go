package main

import "math"

func minDepth(root *TreeNode) int {
	queue := []*TreeNode{}
	queue = append(queue, root)
	var p *TreeNode

	levels := 0
	for len(queue) > 0 {
		tmpLen := len(queue)

		for range tmpLen {
			p = queue[0]
			queue = queue[1:]

			if p.Left == nil && p.Right == nil {
				return levels + 1
			}
			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
		}
		levels++
	}
	return levels
}

// 自上而下递归
func minDepth1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := math.MaxInt

	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, cnt int) {
		if root == nil {
			return
		}

		cnt++
		if cnt >= ans {
			return
		}

		if root.Left == nil && root.Right == nil {
			ans = min(ans, cnt)
			return
		}
		dfs(root.Left, cnt)
		dfs(root.Right, cnt)
	}
	dfs(root, 0) //表示0是根节点上面的虚空节点，等到遍历根节点的时候才+1成为cnt==1。
	return ans
}

func minDepth1_(root *TreeNode) int {
	if root == nil {
		return 0
	}

	ans := 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, cnt int) {
		if root == nil { // 这个是天然的终止条件
			return
		}
		cnt++
		if cnt >= ans {
			return
		}
		if root.Left == nil && root.Right == nil {
			ans = min(ans, cnt)
			return
		}
		dfs(root.Left, cnt)
		dfs(root.Right, cnt)
	}
	dfs(root, 0)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Left == nil {
			return dfs(root.Right) + 1
		}
		if root.Right == nil {
			return dfs(root.Left) + 1
		}
		return min(dfs(root.Left), dfs(root.Right)) + 1
	}
	return dfs(root)
}
