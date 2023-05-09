package binarytree

import "fmt"

func LevelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0, 0)
	if root == nil {
		fmt.Println("root is null")
		return nil
	}
	queue := &Queue{}
	queue.Push(root)
	for queue.Len() > 0 {
		item := make([]int, 0, 0)
		size := queue.Len()
		for i := 0; i < size; i++ {

			v := queue.Peak()
			if v == nil {
				fmt.Println("queue Peak is null")
				return nil
			}
			item = append(item, v.(*TreeNode).Val)
			if v.(*TreeNode).Left != nil {
				queue.Push(v.(*TreeNode).Left)
			}
			if v.(*TreeNode).Right != nil {
				queue.Push(v.(*TreeNode).Right)
			}
		}
		res = append(res, item)
	}
	return res
}
