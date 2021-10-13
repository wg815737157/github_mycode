package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		if head == nil {
			tail = &ListNode{sum % 10, nil}
			head = tail
		} else {
			tail.Next = &ListNode{sum % 10, nil}
			tail = tail.Next
		}
	}
	return
}

func main() {
	l1 := &ListNode{2, nil}
	l2 := &ListNode{3, nil}

	r := addTwoNumbers(l1, l2)
	printListNodeWithHead(r)
}
