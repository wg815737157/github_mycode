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
