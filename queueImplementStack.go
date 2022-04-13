package main

type MyStack struct {
	Q1 []int
	Q2 []int
	M  map[string]string
}

func (mystack *MyStack) Push(x int) {
	mystack.Q2 = append(mystack.Q2, x)
	for i := 0; i < len(mystack.Q1); i++ {
		mystack.Q2 = append(mystack.Q2, mystack.Q1[i])
	}
	mystack.Q1 = make([]int, len(mystack.Q2))
	copy(mystack.Q1, mystack.Q2)
	mystack.Q2 = []int{}
}

func (this *MyStack) Pop() int {
	if len(this.Q1) != 0 {
		tmp := this.Q1[0]
		this.Q1 = this.Q1[1:]
		return tmp
	}
	return -1
}

func (this *MyStack) Top() int {
	if len(this.Q1) != 0 {
		return this.Q1[0]
	}
	return -1
}

func (this *MyStack) Empty() bool {
	return len(this.Q1) == 0
}

func main() {
}
