package binarytree

import "fmt"

func isBalanceTree(binaryTree *BinaryTree) bool {
	if binaryTree == nil {
		return true
	}
	slNodeLen := getDept(binaryTree.slnode)
	srNodeLen := getDept(binaryTree.srnode)
	if slNodeLen-srNodeLen > 1 || slNodeLen-srNodeLen < -1 {
		return false
	}
	result := isBalanceTree(binaryTree.slnode)
	if result == false {
		return false
	}
	result = isBalanceTree(binaryTree.srnode)
	if result == false {
		return false
	}
	return true
}

// 从下到上是否是平衡树算法优化
func isBalanceTree_2(binaryTree *BinaryTree) int {
	if binaryTree == nil {
		return 0
	}
	slNodeLen := isBalanceTree_2(binaryTree.slnode)
	srNodeLen := isBalanceTree_2(binaryTree.srnode)
	if slNodeLen < 0 {
		fmt.Println("返回左侧", binaryTree.slnode.value)
		return -1
	}
	if srNodeLen < 0 {
		fmt.Println("返回右侧", binaryTree.srnode.value)
		return -1
	}
	if slNodeLen-srNodeLen > 1 {
		fmt.Println("不平衡节点是", binaryTree.value)
		return -1
	}
	if srNodeLen-slNodeLen > 1 {
		fmt.Println("不平衡节点是", binaryTree.value)
	}
	if slNodeLen > srNodeLen {
		return slNodeLen + 1
	}
	return srNodeLen + 1
}
