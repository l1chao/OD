package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	first := preorder[0]
	mid := 0
	for i, v := range inorder {
		if v == first {
			mid = i
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:mid+1], inorder[:mid]),
		Right: buildTree(preorder[mid+1:], inorder[mid+1:]),
	}
}
