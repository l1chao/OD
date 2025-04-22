package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head2 := middleNode(head)

	sortList(head)
	sortList(head2)
	return mergeTowLists(head, head2)
}

// 找中左：1.dummy开始 2.用pre。这里用pre
func middleNode(root *ListNode) *ListNode {
	pre, slow, quick := root, root, root
	for quick != nil && quick.Next != nil {
		pre = slow
		slow = slow.Next
		quick = quick.Next.Next
	}
	pre.Next = nil // 注意要将前后链断开！
	return slow
}

func mergeTowLists(list1, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	p, p1, p2 := dummy, list1, list2
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			tmp := p1
			p1 = p1.Next
			p.Next = tmp
			p = p.Next
		} else {
			tmp := p2
			p2 = p2.Next
			p.Next = tmp
			p = p.Next
		}
	}
	if p1 != nil {
		p.Next = p1
	}
	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}
