package main

// 输入：l1 = [2,4,3], l2 = [5,6,4]
// 输出：[7,0,8]

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p1, p2 := l1, l2
	up := 0
	head := new(ListNode)
	p := head

	for p1 != nil && p2 != nil {
		newNode := new(ListNode)
		newNode.Val = (p1.Val + p2.Val + up) % 10
		p.Next = newNode
		p = p.Next

		up = (p1.Val + p2.Val + up) / 10
		p1 = p1.Next
		p2 = p2.Next
	}

	var pRest *ListNode
	if p1 != nil {
		pRest = p1
	} else {
		pRest = p2
	}

	for up != 0 && pRest != nil {
		newNode := new(ListNode)
		newNode.Val = (pRest.Val + up) % 10
		p.Next = newNode
		p = p.Next

		up = (pRest.Val + up) / 10 // 进位为1 or 0
		pRest = pRest.Next
	}
	if up != 0 {
		newNode := new(ListNode)
		newNode.Val = up
		p.Next = newNode
	} else {
		p.Next = pRest
	}

	return head.Next
}
