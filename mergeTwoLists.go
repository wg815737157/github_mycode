package main

import "fmt"

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{-1, nil}
	tmp := dummy
	for l1 != nil && l2 != nil {
		if l1.Val > l2.Val {
			dummy.Next = l2
			l2 = l2.Next
		} else {
			dummy.Next = l1
			l1 = l1.Next
		}
		dummy = dummy.Next
	}
	for l1 != nil {
		dummy.Next = l1
		l1 = l1.Next
		dummy = dummy.Next
	}
	for l2 != nil {
		dummy.Next = l2
		l2 = l2.Next
		dummy = dummy.Next
	}
	return tmp.Next
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
