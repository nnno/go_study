package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry, sum, x, y int
	dummyHead := &ListNode{}
	current := dummyHead

	for l1 != nil || l2 != nil {
		if l1 == nil {
			x = 0
		} else {
			x = l1.Val
		}
		if l2 == nil {
			y = 0
		} else {
			y = l2.Val
		}

		sum = carry + x + y
		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10, Next: nil}
		current = current.Next

		if l1 != nil {
			l1 = l1.Next
		} else {
			l1 = nil
		}

		if l2 != nil {
			l2 = l2.Next
		} else {
			l2 = nil
		}
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry, Next: nil}
	}

	return dummyHead.Next
}
