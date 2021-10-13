package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//前序遍历 preorder = [3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{}
	root.Val = preorder[0]
	lenInorder := len(inorder)

	for i := 0; i < lenInorder; i++ {
		if inorder[i] == preorder[0] {
			leftTreePreorder := preorder[1 : i+1]
			rightTreePreorder := preorder[i+1:]

			leftTreeInorder := inorder[:i]
			rightTreeInorder := inorder[i+1:]

			root.Left = buildTree(leftTreePreorder, leftTreeInorder)
			root.Right = buildTree(rightTreePreorder, rightTreeInorder)
		}
	}
	return root
}

func buildTree2(preorder, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVale := preorder[0]
	root := &TreeNode{rootVale, nil, nil}
	i := 0
	for ; i < len(inorder); i++ {
		if inorder[i] == rootVale {
			break
		}
	}
	left := inorder[:i]
	right := inorder[i+1:]
	root.Left = buildTree2(preorder[1:i+1], left)
	root.Right = buildTree2(preorder[i+1:], right)
	return root
}

func main() {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	treeNode := buildTree2(preorder, inorder)
	printTree(treeNode)
}
