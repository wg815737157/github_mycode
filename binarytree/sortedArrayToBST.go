func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	treeNode := &TreeNode{}
	mid := int(math.Floor(float64(len(nums) / 2)))
	treeNode.Val = nums[mid]

	treeNode.Left = sortedArrayToBST(nums[:mid])
	treeNode.Right = sortedArrayToBST(nums[mid+1:])
	return treeNode
}
