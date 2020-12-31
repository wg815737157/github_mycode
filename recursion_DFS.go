func DFS(binaryTree *BinaryTree, routeArray []int, target int, resArray *[][]int) {
	if binaryTree == nil {
		return
	}
	tmp := binaryTree.value
	routeArray = append(routeArray, tmp)
	if tmp == target {
		*resArray = append(*resArray, routeArray)
		return
	}
	if binaryTree.slnode != nil {
		DFS(binaryTree.slnode, routeArray, target, resArray)
	}
	if binaryTree.srnode != nil {
		DFS(binaryTree.srnode, routeArray, target, resArray)
	}
	return
}
