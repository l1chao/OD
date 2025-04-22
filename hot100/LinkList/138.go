package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 入：head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
// 输出：[[7,null],[13,0],[11,4],[10,2],[1,0]]

func copyRandomList(head *Node) *Node {
	map1 := make(map[*Node]int)
	map2 := make([]*Node, 1)

	newHead := new(Node)
	pnew := newHead
	for i, p := 1, head; p != nil; i, p = i+1, p.Next {
		map1[p] = i

		newNode := new(Node)
		newNode.Val = p.Val
		pnew.Next = newNode
		pnew = pnew.Next
		map2 = append(map2, newNode)
	}
	map1[nil] = 0
	map2[0] = nil

	pnew = newHead.Next
	for p := head; p != nil; p = p.Next {
		pnew.Random = map2[map1[p.Random]]
		pnew = pnew.Next
	}
	return newHead.Next
}
