package main

func postTraversal(r *TreeNode) []int {
	var res []int
	s := Stack{}

	var tmpRight *TreeNode
	for r != nil || s.Len() > 0 {
		for r != nil {
			s.Push(r)
			r = r.Left
		}
		node := s.Pop().(*TreeNode)
		if node.Right == nil || node.Right == tmpRight {
			tmpRight = node
			res = append(res, node.Val)
		} else {
			s.Push(node)
			r = node.Right
		}
	}
	return res
}
func main() {

}
