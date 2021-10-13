package main

func middleNode(head *ListNode) *ListNode {
	l, r := head, head
	for r != nil && r.Next != nil {
		l = l.Next
		r = r.Next.Next
	}
	return l
}

func main() {
	r := &ListNode{}
	middleNode(r)
}
