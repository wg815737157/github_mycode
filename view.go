package main

import (
	"fmt"
	"sync"
)

func Init() *TreeNode {
	treeNode := &TreeNode{}
	treeNode.Val = 1
	treeNodeTmp := &TreeNode{}
	treeNodeTmp.Val = 5
	treeNode.Left = &TreeNode{2, nil, treeNodeTmp}
	treeNodeTmp2 := &TreeNode{}
	treeNodeTmp2.Val = 4
	treeNode.Right = &TreeNode{3, nil, treeNodeTmp2}
	return treeNode
}

func climbStairs(n int) int {
	sliceN := make([]int, n+1, n+1)
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}
	sliceN[2] = 2
	sliceN[1] = 1
	for i := 3; i < n+1; i++ {
		sliceN[i] = sliceN[i-1] + sliceN[i-2]
	}
	return sliceN[n]
}

//func levelOrderBottom(root *TreeNode) [][]int {
//
//}

// 层级遍历
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0, 0)
	if root == nil {
		fmt.Println("root is null")
		return nil
	}
	queue := &Queue{}
	queue.Add(root)
	for queue.Len() > 0 {
		size := queue.Len()
		item := make([]int, size)
		for i := 0; i < size; i++ {
			v := queue.Remove()
			item[i] = v.(*TreeNode).Val
			if v.(*TreeNode).Left != nil {
				queue.Add(v.(*TreeNode).Left)
			}
			if v.(*TreeNode).Right != nil {
				queue.Add(v.(*TreeNode).Right)
			}
		}
		res = append(res, item)
	}
	return res
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		fmt.Println("head is nil")
		return nil
	}
	// 头结点是要的值
	if head.Val == val {
		return head.Next
	}

	headTmp := head

	for head.Next != nil && head.Next.Val != val {
		head = head.Next
	}

	if head.Next == nil {
		fmt.Println("未找到")
		return nil
	}
	head.Next = head.Next.Next

	return headTmp
}

func DFSTmp(root *TreeNode) {
	if root == nil {
		return
	}
	sT := &Stack{}
	sT.Push(root)
	for sT.Len() > 0 {
		item := sT.Pop()
		if item != nil {
			fmt.Println("item is null")
			return
		}
		if _, ok := item.(*TreeNode); !ok {
			fmt.Println("type err")
			return
		}
		fmt.Println(item.(*TreeNode).Val)
		if item.(*TreeNode).Left != nil {
			sT.Push(item.(*TreeNode).Left)
		}
		if item.(*TreeNode).Right != nil {
			sT.Push(item.(*TreeNode).Right)
		}
	}
}

func preOrder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	preOrder(root.Left)
	preOrder(root.Right)
}

var syncMap sync.Map

//右视图
func rightSideView(root *TreeNode) []int {
	if root == nil {
		fmt.Println("root is null")
		return nil
	}
	returnArray := make([]int, 0, 0)
	queue := &Queue{}
	queue.Add(root)
	for queue.Len() > 0 {
		size := queue.Len()
		for i := 0; i < size; i++ {
			v := queue.Peak()
			if v == nil {
				fmt.Println("queue peak is null")
				return nil
			}
			if v.(*TreeNode).Left != nil {
				queue.Add(v.(*TreeNode).Left)
			}
			if v.(*TreeNode).Right != nil {
				queue.Add(v.(*TreeNode).Right)
			}
			if i == size-1 {
				returnArray = append(returnArray, v.(*TreeNode).Val)
			}
		}
	}
	return returnArray
}

//右视图
func rightSideView2(root *TreeNode) []int {
	res := make([]int, 0, 0)
	if root == nil {
		fmt.Println("root is null")
		return nil
	}
	itemStack := &Stack{}
	itemStack.Push(root)
	depthStack := &Stack{}
	depthStack.Push(0)

	mapDepthVal := make(map[int]int)
	for itemStack.Len() > 0 {
		item := itemStack.Pop().(*TreeNode)
		depth := depthStack.Pop().(int)
		if _, ok := mapDepthVal[depth]; !ok {
			mapDepthVal[depth] = item.Val
			res = append(res, item.Val)
		}
		if item.Left != nil {
			itemStack.Push(item.Left)
			depthStack.Push(depth + 1)
		}
		if item.Right != nil {
			itemStack.Push(item.Right)
			depthStack.Push(depth + 1)
		}
	}
	return res
}

// 开根号
func squareTwo() {
	l := 1.4
	h := 1.5

	precision := 0.001

	for h-l > precision {
		mid := (h + l) / 2
		if mid*mid > 2 {
			h = mid
		} else {
			l = mid
		}
	}
	fmt.Println(fmt.Sprintf("%.3f", l))
}

func fib2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib2(n-1) + fib2(n-2)
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	sliceA := make([]int, n+1, n+1)

	sliceA[0] = 0
	sliceA[1] = 1

	for i := 2; i <= n; i++ {
		sliceA[i] = sliceA[i-1] + sliceA[i-2]
	}
	return sliceA[n]
}

//最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight, rightHeight := 0, 0
	if root.Left != nil {
		leftHeight = maxDepth(root.Left)
	}
	if root.Right != nil {
		rightHeight = maxDepth(root.Right)
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}

//func buildTree(preorder []int, inorder []int) *TreeNode {
//	if len(preorder) == 0 {
//		return nil
//	}
//	treeNode := &TreeNode{}
//	root := preorder[0]
//	treeNode.Val = root
//
//	i := 0
//	for ; i < len(inorder); i++ {
//		if preorder[0] == inorder[i] {
//			break
//		}
//	}
//	if i == len(inorder) {
//		fmt.Println("参数错 未找到节点")
//		return nil
//	}
//	treeNode.Left = buildTree(preorder[1:i+1], inorder[:i])
//	treeNode.Right = buildTree(preorder[i+1:], inorder[i+1:])
//
//	return treeNode
//}

func sortedArrayToBST(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}
	treeNode := &TreeNode{}
	mid := n / 2
	treeNode.Val = nums[mid]

	treeNode.Left = sortedArrayToBST(nums[:mid])
	treeNode.Right = sortedArrayToBST(nums[mid+1:])
	return treeNode
}

func revertNum() {
	var input1 int
	i, err := fmt.Scan(&input1)

	fmt.Println("输入的值为：", input1)
	if err != nil {
		fmt.Println(err)
		return
	}
	if i != 1 {
		fmt.Println(i, ":argument err")
		return
	}
	ret := 0
	x := input1

	for x != 0 {
		pop := x % 10
		x /= 10
		ret = ret*10 + pop
	}
	fmt.Println(ret)

}

func main() {

	preorder := []int{3, 9, 8, 5, 4, 10, 20, 15, 7}
	inorder := []int{4, 5, 8, 10, 9, 3, 15, 20, 7}

	treeNode := buildTree(preorder, inorder)

	fmt.Println(levelOrder(treeNode))

	//fmt.Println(fib2(45))
	//
	//sliceA := make([]int, 10, 10)
	//sliceB := sliceA[1:1]
	//fmt.Println(sliceB)
	//
	//if sliceA[0] == 1 {
	//	return
	//}

	//squareTwo()

	//var inputS string
	//_, err := fmt.Scan(&inputS)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//inputSArray := []rune(inputS)
	//lenInputS := len(inputSArray)
	//var res []rune
	//for i := 0; i < lenInputS; {
	//	if i+1 < lenInputS && inputSArray[i] == inputSArray[i+1] {
	//		i += 2
	//		continue
	//	}
	//	res = append(res, inputSArray[i])
	//	i++
	//}
	//fmt.Println(string(res))

	//fmt.Println(climbStairs(44))

	//root := Init()
	//fmt.Println(rightSideView2(root))
	//preOrder(root)
	//fmt.Println("\n")
	//BFSTmp(root)
}
