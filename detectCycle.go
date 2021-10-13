package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}

func main() {
	s := "aaa"

	b := []byte("ddd")
	fmt.Println(string(b))
	n := len(s)
	for i := 0; i < n; i++ {
		fmt.Printf("%c", s[i])
	}

}
