package main

func postOrderRecursionTraversal(t *TreeNode) {
	res := []int{}

	if t.Left != nil {
		postOrderRecursionTraversal(t.Left)
	}
	if t.Right != nil {
		postOrderRecursionTraversal(t.Left)
	}
	res = append(res, t.Val)
}
