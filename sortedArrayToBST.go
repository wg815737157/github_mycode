package main

import "fmt"

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	mid := n / 2
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6}
	node := sortedArrayToBST(a)
	fmt.Println(inOrderNonRecursionTraversal(node))
	s := &Stack{}
	for node != nil || s.Len() != 0 {
		for node != nil {
			s.Push(node)
			node = node.Left
		}
		node = s.Pop().(*TreeNode)
		fmt.Println(node.Val)
		node = node.Right
	}
}
