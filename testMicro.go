package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type node struct {
	Value int
	Left  *node
	Right *node
}

func solution2(sArr []string) string {
	if len(sArr) == 0 {
		return ""
	}
	res := []byte{}
	for j := 0; j < len(sArr[0]); j++ {
		ch := sArr[0][j]
		for i := 1; i < len(sArr); i++ {
			if len(sArr[i]) == j || sArr[i-1][j] != sArr[i][j] {
				return string(res)
			}
		}
		res = append(res, ch)
	}
	return string(res)
}

var M *MutexString

func _init() {
	M = &MutexString{nodeNameMap: make(map[string]*microNode)}
}

type microNode struct {
	mutex sync.Mutex
}
type MutexString struct {
	nodeNameMap map[string]*microNode
}

func Lock(nodeName string) {
	M.nodeNameMap[nodeName] = &microNode{}
	M.nodeNameMap[nodeName].mutex.Lock()
	return
}
func UnLock(nodeName string) error {
	if _, ok := M.nodeNameMap[nodeName]; !ok {
		return errors.New("node 不存在")
	}
	M.nodeNameMap[nodeName].mutex.Unlock()
	delete(M.nodeNameMap, nodeName)
	return nil
}

func solutions3(root *ListNode, k int) *ListNode {
	if k <= 0 {
		return root
	}
	dummy := &ListNode{-1, root}
	l, r := dummy, dummy
	i := 0
	for r != nil {
		if i > k {
			l = l.Next
		}
		r = r.Next
		i++
	}
	if i >= k {
		l.Next = l.Next.Next
	}
	return dummy.Next
}
func solution(root *node, k int) int {
	arr := []int{}
	s := &Stack{}
	i := 0
	for root != nil || s.Len() > 0 {
		for root != nil {
			s.Push(root)
			root = root.Left
		}
		tmp := s.Pop().(*node)
		if i == k-1 {
			return tmp.Value
		}
		arr = append(arr, tmp.Value)
		root = tmp.Right
		i++
	}
	return -1
}
func solution4() {
	rand.Seed(time.Now().Unix())
	resA := 0
	resB := 0
	isA := false
	turn := rand.Intn(2)
	if turn%1 == 0 {
		isA = true
	}
	for i := 0; i < 100000; i++ {
		r := rand.Intn(6)
		if isA {
			resA += 2
			isA = false
			continue
		}
		if r%6 == 0 {
			resA += 2
			isA = false
			continue
		}
		resB += 3
		isA = true
	}
	fmt.Println(resA, resB)
}

func dfs(matrix [][]int, i, j int) int {
	row := len(matrix)
	column := len(matrix[0])
	elementLen := 0
	// 向右
	for right := j + 1; right < column && matrix[i][right] > 0; right++ {
		elementLen += dfs(matrix, i, right)
	}
	// 向下
	for down := i + 1; down < row && matrix[i][down] > 0; down++ {
		elementLen += dfs(matrix, i, down)
		matrix[i][down] = -1
	}
	return elementLen
}

func solution5(matrix [][]int) int {
	row := len(matrix)
	if row == 0 {
		return -1
	}
	column := len(matrix[0])
	countN := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if matrix[i][j] <= 0 {
				continue
			}
			elementLen := 0
			if left := j - 1; 0 <= left && matrix[i][left] > 0 {
				elementLen++
			}
			if right := j + 1; right < column && matrix[i][right] > 0 {
				elementLen++
			}
			if up := i - 1; 0 <= up && matrix[i][up] > 0 {
				elementLen++
			}
			if down := i + 1; down < row && matrix[i][down] > 0 {
				elementLen++
			}
			countN += elementLen
		}
	}
	return countN / 2
}

func main() {
	// 题目一
	//node4 := &node{1, nil, nil}
	//ndoe3 := &node{4, nil, nil}
	//node2 := &node{2, node4, nil}
	//node1 := &node{3, node2, ndoe3}
	//fmt.Println(solution(node1, 3))
	// 题目二
	//fmt.Println(solution2([]string{"ab", "a"}))
	// 题目三
	//listNode3 := &ListNode{3, nil}
	//listNode2 := &ListNode{2, listNode3}
	//root := &ListNode{1, listNode2}
	//newRoot := solutions3(root, 3)
	//for newRoot != nil {
	//	fmt.Println(newRoot.Val)
	//	newRoot = newRoot.Next
	//}

	// 题目四
	//solution4()
	matrix := [][]int{{0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}, {0, 1, 1, 1}}
	fmt.Println(solution5(matrix))
}
