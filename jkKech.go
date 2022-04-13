package main

import (
	"fmt"
	"runtime"
	"time"
)

func solution(s string) bool {

	//left := make(map[byte]bool)
	//left['('] = true
	//left['['] = true
	//left['{'] = true
	//right := make(map[byte]byte)
	//right[')'] = '('
	//right[']'] = ']'
	//right['}'] = '}'
	//a := []byte{}
	//for i := 0; i < len(s); i++ {
	//	if _, ok := left[s[i]]; ok {
	//		a = append(a, s[i])
	//		continue
	//	}
	//	if _, ok := right[s[i]]; ok {
	//		if a[len(a)-1] == right[s[i]] {
	//			a = a[:len(a)-1]
	//			continue
	//		}
	//		return false
	//	}
	//	return false
	//}
	//if len(a) != 0 {
	//	return false
	//}
	return true
}
func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		for {
			i := 0
			i++
		}
	}()
	runtime.Gosched()
	fmt.Println("DD")
	time.Sleep(10 * time.Second)
	//ch := make(chan int, 1)
	//
	//for i := 0; i < 10; i++ {
	//	select {
	//	case n := <-ch:
	//		fmt.Println(n)
	//	case ch <- i:
	//	}
	//}
	//a := int32(1)
	//ok := atomic.CompareAndSwapInt32(&a, 1, 3)
	//fmt.Println(a)
	//fmt.Println(ok)
	//v:=atomic.Value{}
	//v.Store(2)
	//i:=v.Load().(int)
	//fmt.Println(i)

	//var cond sync.Stdout

	//var er interface{}
	//var w io.WriteCloser
	//er = w

	//var syncMap sync.Map
	//syncMap.
	//
	//syncMap:=&sync.Map{}
	//syncMap.Store(1, "1")
	//if v, ok := syncMap.Load(1); ok {
	//	fmt.Println(v.(string))
	//}

	//node5 := &TreeNode{6, nil, nil}
	//node4 := &TreeNode{1, nil, nil}
	//node3 := &TreeNode{4, nil, node5}
	//node2 := &TreeNode{2, node4, nil}
	//node := &TreeNode{3, node2, node3}
	//
	//stack := []*TreeNode{}
	//for len(stack) != 0 || node != nil {
	//	for node != nil {
	//		stack = append(stack, node)
	//		node = node.Left
	//	}
	//	popNode := stack[len(stack)-1]
	//	stack = stack[:len(stack)-1]
	//	fmt.Println(popNode.Val)
	//	node = popNode.Right
	//}
}
