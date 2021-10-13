func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	treeNode := &TreeNode{}
	root := preorder[0]
	treeNode.Val = root

	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	if i == len(inorder) {
		fmt.Println("参数错 未找到节点")
		return nil
	}
	treeNode.Left = buildTree(preorder[1:i+1], inorder[:i])
	treeNode.Right = buildTree(preorder[i+1:], inorder[i+1:])

	return treeNode
}
