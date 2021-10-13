package main

import "fmt"

func revertList(head *ListNode) *ListNode {
	cur := head
	next := head.Next
	cur.Next = nil
	for next != nil {
		tmp := next.Next
		next.Next = cur
		cur = next
		next = tmp
	}
	return cur
}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func revertList2(head *ListNode) *ListNode {
	l := head
	if l == nil {
		return nil
	}
	r := head.Next
	for r != nil {
		tmp := r.Next
		r.Next = l
		l = r
		r = tmp
	}
	head.Next = nil
	return l
}
func main() {
	node1 := &ListNode{1, nil}
	head := &ListNode{0, node1}
	//PrintList(head)
	head = revertList2(head)
	PrintList(head)
}
