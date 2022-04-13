package main

import "fmt"

type CQueue struct {
	s1 []int
	s2 []int
}

func Constructor() CQueue {
	return CQueue{
		s1: []int{},
		s2: []int{},
	}
}

func (c *CQueue) AppendTail(x int) {
	c.s1 = append(c.s1, x)
}

func (c *CQueue) DeleteHead() int {
	res := -1
	if c.empty() {
		return res
	}
	if len(c.s2) == 0 {
		for i := len(c.s1) - 1; i >= 0; i-- {
			c.s2 = append(c.s2, c.s1[i])
			c.s1 = c.s1[:i]
		}
	}
	res = c.s2[len(c.s2)-1]
	c.s2 = c.s2[:len(c.s2)-1]
	return res

}

func (c *CQueue) empty() bool {
	if len(c.s1) == 0 && len(c.s2) == 0 {
		return true
	}
	return false
}

func main() {
	q := Constructor()
	q.AppendTail(5)
	q.AppendTail(2)

	fmt.Println(q.DeleteHead())
	fmt.Println(q.DeleteHead())
}
