package main

import (
	"container/list"
	"fmt"
)

func main() {
	cQueue := Constructor()
	cQueue.AppendTail(1)
	cQueue.AppendTail(2)
	cQueue.AppendTail(3)
	fmt.Println(cQueue.DeleteHead())
	fmt.Println(cQueue.DeleteHead())
	fmt.Println(cQueue.DeleteHead())

}

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}
func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		for this.stack1.Len() != 0 {
			e := this.stack1.Back()
			this.stack1.Remove(e)
			this.stack2.PushBack(e.Value.(int))
		}
	}
	if this.stack2.Len() == 0 {
		return -1
	}
	e := this.stack2.Back()
	this.stack2.Remove(e)
	return e.Value.(int)
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
