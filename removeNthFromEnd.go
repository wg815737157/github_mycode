package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	sentinel := &ListNode{0, head}
	l, r := sentinel, sentinel
	i := 0
	for ; r != nil; {
		if i > n {
			l = l.Next
		}
		r = r.Next
		i++
	}
	if i > n {
		l.Next = l.Next.Next
	}
	return sentinel.Next
}

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	first, second := head, dummy
	for i := 0; i < n; i++ {
		first = first.Next
	}
	for ; first != nil; first = first.Next {
		second = second.Next
	}
	second.Next = second.Next.Next
	return dummy.Next
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	r := initListNodeByArray(a)
	r = removeNthFromEnd(r, 5)
	printListNode(r)
}
