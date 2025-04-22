package main

// 输入：head = [1,2,2,1]
// 输出：true

// 要点分解：1.获取中间节点（中左、中右） 2.头插法逆转链表
func isPalindrome(head *ListNode) bool {
	slow, quick := head, head

	for quick != nil && quick.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
	}

	var p *ListNode
	if quick == nil {
		p = slow
	} else {
		p = slow.Next
	}
	blank := new(ListNode)
	for p != nil {
		tmp := p
		p = p.Next
		tmp.Next = blank.Next
		blank.Next = tmp
	}

	p = blank.Next
	p1 := head
	for p != nil {
		if p.Val != p1.Val {
			return false
		}
		p = p.Next
		p1 = p1.Next
	}
	return true
}
