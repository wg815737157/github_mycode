package main

import "fmt"

func checkStackQueue(a []int, b []int) bool {
	lenA := len(a)
	s := &Stack{}
	j := 0
	for i := 0; i < lenA; i++ {
		s.Push(a[i])
		for ; s.Len() > 0 && s.Top().(int) == b[j]; {
			s.Pop()
			j++
		}
	}
	if s.Len() > 0 {
		return false
	}
	return true

}
func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 5, 4, 1, 3}
	fmt.Println(checkStackQueue(a, b))
}
