func reverseList(head *ListNode) (revertHead *ListNode) {
	if head == nil {
		return nil
	}
	if head.nextNode == nil {
		return head
	}
	newHead := reverseList(head.nextNode)
	head.nextNode.nextNode = head
	head.nextNode = nil
	return newHead
}
