package main

import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	for l1 != nil || l2 != nil {
		if l1 == nil {
			if head == nil {
				head = l2
			} else {
				tail.Next = l2
			}
			return head
		}

		if l2 == nil {
			if head == nil {
				head = l1
			} else {
				tail.Next = l1
			}
			return head
		}

		new := &ListNode{}
		if l1.Val < l2.Val {
			new.Val = l1.Val
			l1 = l1.Next
		} else {
			new.Val = l2.Val
			l2 = l2.Next
		}
		if head == nil {
			head = new
			tail = new
		} else {
			tail.Next = new
			tail = tail.Next
		}
	}
	return head
}

func main() {
	l1 := &ListNode{2, &ListNode{6, nil}}
	l2 := &ListNode{4, &ListNode{8, nil}}
	l3 := mergeTwoLists(l1, l2)
	for l3 != nil {
		fmt.Println(l3.Val)
		l3 = l3.Next
	}

}
