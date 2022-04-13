package main

import "fmt"

func binaryTreePost(root *TreeNode) {
	stack := &Stack{}
	var pre *TreeNode
	for root != nil || stack.Len() != 0 {
		for root != nil {
			stack.Push(root)
			root = root.Left
		}
		treeNode := stack.Pop().(*TreeNode)
		if treeNode.Right == nil || treeNode.Right == pre {
			fmt.Println(treeNode.Val)
			pre = treeNode
		} else {
			stack.Push(treeNode)
			root = treeNode.Right
		}
	}
}
func initTree() *TreeNode {
	treeNode2 := &TreeNode{3, nil, nil}
	treeNode1 := &TreeNode{2, nil, nil}
	return &TreeNode{1, treeNode1, treeNode2}
}
func main() {
	treeNode := initTree()
	binaryTreePost(treeNode)
}
