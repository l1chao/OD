package main

//  输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]

// n>=2 && n<=len
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	front, back := head, head
	for range n {
		back = back.Next
		if back == nil {
			return head.Next
		}
	}
	for back.Next != nil {
		back = back.Next
		front = front.Next
	}

	tmp := front.Next
	front.Next = tmp.Next

	return head
}
