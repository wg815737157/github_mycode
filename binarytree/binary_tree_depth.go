package binarytree

func getDept(binaryTree *BinaryTree) int {
	if binaryTree == nil {
		return 0
	}
	slNodeLen := getDept(binaryTree.slnode)
	srNodeLen := getDept(binaryTree.srnode)
	if slNodeLen > srNodeLen {
		return slNodeLen + 1
	}
	return srNodeLen + 1
}
