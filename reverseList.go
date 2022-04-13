package main

import "fmt"

func reverseList(listNode *ListNode) *ListNode {
	cur := listNode
	if cur != nil && cur.Next != nil {
		next := cur.Next
		newHead := reverseList(next)
		next.Next = cur
		cur.Next = nil
		return newHead
	}
	return cur
}
func reverseList2(head *ListNode) (pre *ListNode) {
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
func PrintList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

func main() {
	listNode := &ListNode{}
	reverseList(listNode)
}
