package main

import "fmt"

func twoMaxNum(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Pre  *LinkedList
	Val  interface{}
	Next *LinkedList
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Queue struct {
	Array []interface{}
}

func (q *Queue) Len() int {
	return len(q.Array)
}

func (q *Queue) Add(i interface{}) {
	if q.Array == nil {
		q.Array = make([]interface{}, 0, 0)
	}
	q.Array = append(q.Array, i)
}
func (q *Queue) Remove() interface{} {
	lenQueue := len(q.Array)
	if lenQueue == 0 {
		return nil
	}
	res := q.Array[0]
	q.Array = q.Array[1:]
	return res
}

func (q *Queue) Peak() interface{} {
	lenQueue := len(q.Array)
	if lenQueue == 0 {
		return nil
	}
	return q.Array[0]
}

type Stack struct {
	Array []interface{}
}

func (s *Stack) Peak() interface{} {
	lenQueue := len(s.Array)
	if lenQueue == 0 {
		return nil
	}
	return s.Array[lenQueue-1]
}

func (s *Stack) Push(i interface{}) {
	if s.Array == nil {
		s.Array = make([]interface{}, 0, 0)
	}
	s.Array = append(s.Array, i)
}
func (s *Stack) Len() int {
	return len(s.Array)
}
func (s *Stack) Pop() interface{} {
	lenQueue := len(s.Array)
	if lenQueue == 0 {
		return nil
	}
	res := s.Array[lenQueue-1]
	s.Array = s.Array[:lenQueue-1]
	return res
}

func revertArray(a []int) {
	n := len(a)
	for i := 0; i < n/2; i++ {
		a[i], a[n-i-1] = a[n-i-1], a[i]
	}
}

func swap(array []int, low, high int) {
	array[low], array[high] = array[high], array[low]
}

func initListNodeByArrayWithHead(a []int) *ListNode {
	head := &ListNode{0, nil}
	r := head
	for i := 0; i < len(a); i++ {
		node := &ListNode{a[i], nil}
		r.Next = node
		r = node
	}
	return head
}

func initListNodeByArray(a []int) *ListNode {
	head := &ListNode{0, nil}
	r := head
	for i := 0; i < len(a); i++ {
		node := &ListNode{a[i], nil}
		r.Next = node
		r = node
	}
	return head.Next
}

func printListNode(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func printListNodeWithHead(node *ListNode) {
	node = node.Next
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func initTreeNodeBySortedArray(a []int) *TreeNode {
	n := len(a)
	if n == 0 {
		return nil
	}
	mid := n / 2
	treeNode := &TreeNode{a[mid], nil, nil}
	treeNode.Left = initTreeNodeBySortedArray(a[:mid])
	treeNode.Right = initTreeNodeBySortedArray(a[mid+1:])
	return treeNode
}

func printTree(treeNode *TreeNode) {
	if treeNode != nil {
		fmt.Println(treeNode.Val)
		printTree(treeNode.Left)
		printTree(treeNode.Right)
	}
}
