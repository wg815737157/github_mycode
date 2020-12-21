//删除链表倒数第N个节点
func deleteDescN_2(head *ListNode, n int) error {
	if head == nil || n == 0 {
		return errors.New("空链表或0")
	}
	i := 0
	tmp := head
	deleteNodePre := head
	deleteNode := head
	for tmp != nil {
		i++
		if i > n {
			deleteNodePre = deleteNode
			deleteNode = deleteNode.nextNode
		}
		tmp = tmp.nextNode
	}
	if i <= n {
		return errors.New("删除节点大于等于链表节点数")
	}
	deleteNodePre.nextNode = deleteNode.nextNode
	return nil
}


func deleteDescN_1(head *ListNode, n int) error {
	if head == nil {
		return errors.New("链表空")
	}
	len := 1
	tmp := head
	for tmp.nextNode != nil {
		tmp = tmp.nextNode
		len++
	}
	if len < n {
		return errors.New("超过链表长度")
	}
	m := len - n
	err := deleteN(head, m)
	if err != nil {
		return err
	}
	return nil
}
