package main

import "fmt"

//给你一棵 完全二叉树 的根节点 root ，求出该树的节点个数。
//000000011111111
type treeNode struct {
	val       int
	leftNode  *treeNode
	rightNode *treeNode
}

func countTreeNode(t *treeNode) (sum int) {
	if t == nil {
		return 0
	}
	sum++
	return sum + countTreeNode(t.rightNode) + countTreeNode(t.leftNode)
}

//
func treeSumNode2(t *treeNode) (sum int) {
	return 0

}

func main() {

	t4 := &treeNode{val: 0}
	t2 := &treeNode{}
	t3 := &treeNode{0, t4, nil}
	t := &treeNode{0, t2, t3}

	fmt.Println(countTreeNode(t))
}
