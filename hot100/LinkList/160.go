package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	l1, l2 := 0, 0
	p1, p2 := headA.Next, headB.Next

	for p1 != nil {
		l1++
		p1 = p1.Next
	}
	for p2 != nil {
		l2++
		p2 = p2.Next
	}

	maxl := max(l1, l2)
	p1 = headA.Next
	p2 = headB.Next

	if l1 == maxl {
		for range l1 - l2 {
			p1 = p1.Next
		}
		for p1 != nil {
			if p1 == p2 {
				return p1
			}
			p1 = p1.Next
			p2 = p2.Next
		}
	} else {
		for range l2 - l1 {
			p2 = p2.Next
		}
		for p2 != nil {
			if p1 == p2 {
				return p1
			}
			p1 = p1.Next
			p2 = p2.Next
		}
	}

	return nil
}

func main() {
	for range 0 {
		fmt.Println("nihao")
	}
}
