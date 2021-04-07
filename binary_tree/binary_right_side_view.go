//BFS 广度优先搜索
func rightSideView(root *TreeNode) []int {
	if root == nil {
		fmt.Println("root is null")
		return nil
	}
	returnArray := make([]int, 0, 0)
	queue := &Queue{}
	queue.Push(root)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			v := queue.Peak()
			if v == nil {
				fmt.Println("queue peak is null")
				return nil
			}
			if v.(*TreeNode).Left != nil {
				queue.Push(v.(*TreeNode).Left)
			}
			if v.(*TreeNode).Right != nil {
				queue.Push(v.(*TreeNode).Right)
			}
			if i == size-1 {
				returnArray = append(returnArray, v.(*TreeNode).Val)
			}
		}
	}
	return returnArray
}

