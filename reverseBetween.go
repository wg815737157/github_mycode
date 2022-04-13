package main

import "fmt"

func reverseBetween(head *ListNode, m, n int) *ListNode {
	dummy := &ListNode{-1, head}
	cur := dummy
	var pre *ListNode
	for i := 0; i < n && cur != nil; i++ {
		if i == m-1 {
			pre = cur
		}
		if i >= m {
			next := cur.Next
			cur.Next = cur.Next.Next
			next.Next = pre.Next
			pre.Next = next
			continue
		}
		cur = cur.Next
	}
	return dummy.Next
}

func reverseBetween2(head *ListNode, m int, n int) *ListNode {
	dummy := &ListNode{0, head}
	pre := dummy
	for i := 1; i < m; i++ {
		pre = pre.Next
	}
	cur := pre.Next
	for i := 0; i < n-m; i++ {
		next := cur.Next
		pre.Next = next
		cur.Next = next.Next
		next.Next = cur
	}
	return dummy.Next
}
func main() {
	//node4 := &ListNode{4, nil}
	node3 := &ListNode{3, nil}
	node2 := &ListNode{2, node3}
	head := &ListNode{1, node2}
	head = reverseBetween2(head, 1, 3)
	i := 0
	for ; i < 4 && head != nil; i++ {
		fmt.Println(head.Val)
		head = head.Next
	}
}
