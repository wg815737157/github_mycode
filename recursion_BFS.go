func BFS(binaryTreeSlice []*BinaryTree, routeArray []int, target int, resArray *[][]int) {
	if len(binaryTreeSlice) == 0 {
		return
	}
	tmp := make([]*BinaryTree, 0, 0)
	for _, binaryTree := range binaryTreeSlice {
		tmpValue := binaryTree.value
		routeArray = append(routeArray, tmpValue)
		if tmpValue == target {
			*resArray = append(*resArray, routeArray)
			continue
		}
		if binaryTree.slnode != nil {
			tmp = append(tmp, binaryTree.slnode)
		}
		if binaryTree.srnode != nil {
			tmp = append(tmp, binaryTree.srnode)
		}
	}
	BFS(tmp, routeArray, target, resArray)
	return
}
