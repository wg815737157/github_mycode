package main

import (
	"fmt"
	"strconv"
)

func findPath(root *TreeNode, path string, res *[]string) {
	path += strconv.Itoa(root.Val)
	if root.Left != nil {
		newPath := path + "->"
		findPath(root.Left, newPath, res)
	}
	if root.Right != nil {
		newPath := path + "->"
		findPath(root.Right, newPath, res)
	}
	if root.Right == nil && root.Left == nil {
		*res = append(*res, path)
	}
	return
}

func binaryTreePaths(root *TreeNode) []string {
	res := &[]string{}
	var path string
	if root == nil {
		return nil
	}
	findPath(root, path, res)
	return *res
}

func main() {
	node3 := &TreeNode{3, nil, nil}
	node2 := &TreeNode{2, nil, nil}
	node1 := &TreeNode{1, node3, nil}
	root := &TreeNode{0, node1, node2}
	fmt.Println(binaryTreePaths(root))
}
