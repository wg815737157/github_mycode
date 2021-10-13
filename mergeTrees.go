package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 != nil && root2 != nil {
		root1.Val = root1.Val + root2.Val
		root1.Left = mergeTrees(root1.Left, root2.Left)
		root1.Right = mergeTrees(root1.Right, root2.Right)
	} else if root1 == nil {
		return root2
	}
	return root1
}
func main() {
	root1 := initTreeNodeBySortedArray([]int{1, 2, 3})
	root2 := initTreeNodeBySortedArray([]int{1, 2, 3})
	root := mergeTrees(root1, root2)
	printTree(root)
}
