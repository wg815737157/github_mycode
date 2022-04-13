package main

import "fmt"

//给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
//000000011111111

func getTreeHeight(root *TreeNode) int {
	height := 0
	for root != nil {
		root = root.Left
		height++
	}
	return height
}
func countNodes(root *TreeNode) int {
	h := getTreeHeight(root)

	if h == 0 {
		return 0
	}
	if h == 1 {
		return 1
	}
	// h<=31
	l := 1 << (h - 1)
	r := 1<<h - 1
	for l <= r {
		mid := l + (r-l)>>1
		tmp := root
		for i := h - 2; i >= 0; i-- {
			flag := 1 << i
			if mid&flag == flag {
				tmp = tmp.Right
			} else {
				tmp = tmp.Left
			}
		}
		if tmp == nil {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}

func main() {
	t6 := &TreeNode{Val: 6}
	t5 := &TreeNode{Val: 5}
	t4 := &TreeNode{Val: 4}
	t3 := &TreeNode{3, t6, nil}
	t2 := &TreeNode{2, t4, t5}
	t1 := &TreeNode{1, t2, t3}
	fmt.Println(countNodes(t1))
}
