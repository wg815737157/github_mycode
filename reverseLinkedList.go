package main

import "fmt"

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newRoot := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newRoot
}

func solution(root *ListNode) *ListNode {
	_, newRoot := reverse(root)
	root.Next = nil
	return newRoot
}

func reverse(node *ListNode) (next *ListNode, root *ListNode) {
	if node == nil || node.Next == nil {
		return node, node
	}
	next, root = reverse(node.Next)
	next.Next = node
	return node, root
}

func reverse3(node *ListNode) *ListNode {
	var pre *ListNode
	cur := node
	if cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
func main() {
	node := &ListNode{2, nil}
	root := &ListNode{1, node}
	newRoot := reverse3(root)
	for newRoot != nil {
		fmt.Println(newRoot.Val)
		newRoot = newRoot.Next
	}
}
