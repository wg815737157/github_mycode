package main

func DeleteN(l *ListNode, n int) *ListNode {
	newHead := &ListNode{-1, l}
	tmp, deleteNode := newHead, newHead
	for i := 0; tmp != nil; {
		if i > n {
			deleteNode = deleteNode.Next
		}
		i++
		tmp = tmp.Next
	}
	deleteNode.Next = deleteNode.Next.Next
	return newHead.Next
}

func main() {
	l := &ListNode{}
	DeleteN(l, 2)
}
