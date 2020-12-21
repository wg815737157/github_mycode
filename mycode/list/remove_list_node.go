//删除链表节点
func deleteN(head *ListNode, n int) error {
	if head == nil {
		return errors.New("链表空")
	}
	if n == 0 {
		return errors.New("要删除的是链表的头节点")
	}
	tmp := head
	tmpNext := head.nextNode
	for i := 1; i < n; i++ {
		if tmp == nil {
			return errors.New("要删除的节点超过实际节点")
		}
		tmpNext = tmp
		tmp = tmp.nextNode
	}
	if tmp == nil {
		tmpNext.nextNode = nil
		return nil
	}
	tmp.nextNode = tmpNext.nextNode
	return nil
}
