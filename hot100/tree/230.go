package main

func kthSmallest(root *TreeNode, k int) int {
	cnt := 0
	isFound := false
	ans := 0
	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil || isFound == true {
			return
		}

		inorder(root.Left)
		cnt++
		if cnt == k {
			isFound = true
			ans = root.Val
		}
		inorder(root.Right)
	}
	inorder(root)
	return ans
}
