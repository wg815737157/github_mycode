package binarytree

import "fmt"

// 深度优先搜索
func rightSideView2(root *TreeNode) []int {
	res := make([]int, 0, 0)
	if root == nil {
		fmt.Println("root is null")
		return nil
	}
	itemStack := &Stack{}
	itemStack.Push(root)
	depthStack := &Stack{}
	depthStack.Push(0)

	mapDepthVal := make(map[int]int)
	for itemStack.Len() > 0 {
		item := itemStack.Pop().(*TreeNode)
		depth := depthStack.Pop().(int)
		if _, ok := mapDepthVal[depth]; !ok {
			mapDepthVal[depth] = item.Val
			res = append(res, item.Val)
		}
		if item.Left != nil {
			itemStack.Push(item.Left)
			depthStack.Push(depth + 1)
		}
		if item.Right != nil {
			itemStack.Push(item.Right)
			depthStack.Push(depth + 1)
		}
	}
	return res
}
