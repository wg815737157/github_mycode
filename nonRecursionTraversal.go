package main

func preOrderNonRecursionTraversal(t *TreeNode) []int {
	res := []int{}
	s := &Stack{}
	for s.Len() > 0 || t != nil {
		for t != nil {
			s.Push(t)
			res = append(res, t.Val)
			t = t.Left
		}
		t = s.Pop().(*TreeNode)
		t = t.Right
	}
	return res
}
func inOrderNonRecursionTraversal(t *TreeNode) []int {
	res := []int{}
	s := &Stack{}
	node := t
	for s.Len() > 0 || node != nil {
		for node != nil {
			s.Push(node)
			node = node.Left
		}
		node = s.Pop().(*TreeNode)
		res = append(res, node.Val)
		node = node.Right
	}
	return res
}
