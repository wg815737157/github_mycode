package main

func main() {
	root := &ListNode{}
	deleteDuplicates(root)
}
func deleteDuplicates(root *ListNode) *ListNode {
	dummy := &ListNode{-1, root}
	pre := dummy
	cur := root
	if cur == nil {
		return root
	}
	next := cur.Next
	if next == nil {
		return root
	}
	for next != nil {
		if cur.Val != next.Val {
			pre = cur
			cur = next
			next = next.Next
			continue
		}
		for next != nil && cur.Val == next.Val {
			cur = next
			next = next.Next
		}

		// 先结束
		if next == nil {
			pre.Next = next
			return dummy.Next
		}
		// 不相等
		pre.Next = next
		cur = next
		next = next.Next
	}
	return dummy.Next
}
