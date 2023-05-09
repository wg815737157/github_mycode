package binarytree

import "fmt"
import "algorithm/util"

// BFS 广度优先搜索
func rightSideView(root *util.TreeNode) []int {
	if root == nil {
		fmt.Println("root is null")
		return nil
	}
	returnArray := make([]int, 0, 0)
	queue := &util.Queue{}
	queue.Add(root)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			v := queue.Peak()
			if v == nil {
				fmt.Println("queue peak is null")
				return nil
			}
			if v.(*util.TreeNode).Left != nil {
				queue.Add(v.(*util.TreeNode).Left)
			}
			if v.(*util.TreeNode).Right != nil {
				queue.Add(v.(*util.TreeNode).Right)
			}
			if i == size-1 {
				returnArray = append(returnArray, v.(*util.TreeNode).Val)
			}
		}
	}
	return returnArray
}
