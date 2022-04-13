package main

func lowestCommonAncestor(node, p, q *TreeNode) *TreeNode {
	if node == nil {
		return nil
	}
	if node.Val == p.Val || node.Val == q.Val {
		return node
	}
	left := lowestCommonAncestor(node.Left, p, q)
	right := lowestCommonAncestor(node.Right, p, q)
	if left != nil && right != nil {
		return node
	}
	if left != nil {
		return left
	}
	return right
}
func main() {
	root := &TreeNode{}
	lowestCommonAncestor(root, root, root)

}
