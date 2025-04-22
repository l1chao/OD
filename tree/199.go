package main

import "fmt"

func rightSideView(root *TreeNode) []int {
	ans := []int{}

	queue := []*TreeNode{}
	queue = append(queue, root)
	for len(queue) > 0 {
		tmpLen := len(queue)

		for i := range tmpLen {
			p := queue[0]
			queue = queue[1:]
			if i == tmpLen-1 {
				ans = append(ans, p.Val)
			}

			if p.Left != nil {
				queue = append(queue, p.Left)
			}
			if p.Right != nil {
				queue = append(queue, p.Right)
			}
		}
	}
	return ans
}

type MyInt int

func my() {
	var p any
	var q MyInt = 32
	p = q
	fmt.Println(p)

}
