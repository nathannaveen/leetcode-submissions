func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return l1
	} else if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}

	l := &ListNode{}
	l1, l2, l = maxAndExtract(l1, l2, l)
	head := l

	for l1 != nil && l2 != nil {
		l1, l2, l = maxAndExtract(l1, l2, l)
	}
	for l1 != nil {
		l.Next = l1 
		l = l.Next
		l1 = l1.Next
	}
	for l2 != nil {
		l.Next = l2
		l = l.Next
		l2 = l2.Next
	}

	return head
}

func maxAndExtract(l1, l2, l *ListNode) (*ListNode, *ListNode, *ListNode) {
	if l1.Val > l2.Val {
		l.Next = l2 
		l = l.Next
		l2 = l2.Next
		return l1, l2, l
	}

	l.Next = l1 
	l = l.Next
	l1 = l1.Next
	return l1, l2, l
}
