package huawei

func ReverseList(pHead *ListNode) *ListNode {
	if pHead == nil || pHead.Next == nil {
		return pHead
	}
	var newHead *ListNode
	var temp *ListNode
	for pHead != nil {
		temp = pHead
		pHead = pHead.Next

		temp.Next = newHead
		newHead = temp
	}
	return newHead
}
