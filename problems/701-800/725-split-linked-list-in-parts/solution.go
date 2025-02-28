func splitListToParts(head *ListNode, k int) []*ListNode {
	res := []*ListNode{}
	length := 0
	cur := head

	for cur != nil {
		cur = cur.Next
		length++
	}
	cur = head

	lengthToAddEveryTime := length / k
	numTimesToAddOne := length % k

	for i := 0; i < k; i++ {
		newlist := &ListNode{}
        appendToResList := newlist
		lengthOfCurentList := lengthToAddEveryTime
        contains := false

		if numTimesToAddOne > 0 {
			numTimesToAddOne--
			lengthOfCurentList++
		}
		if cur != nil {
			newlist.Val = cur.Val
			cur = cur.Next
            contains = true
		}
        
        for j := 1; j < lengthOfCurentList; j++ {
            newlist.Next = &ListNode{ Val: cur.Val, }
            newlist = newlist.Next
			cur = cur.Next
		}

        if contains {
            res = append(res, appendToResList)
        } else {
            res = append(res, nil)
        }
	}

	return res
}