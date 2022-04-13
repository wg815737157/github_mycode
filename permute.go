package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func permutation(num []int, n int, curIndex int, res *[][]int) {
	if curIndex == n-1 {
		fmt.Println(num)
		tmp := make([]int, len(num))
		copy(tmp, num)
		*res = append(*res, tmp)
		return
	}
	for i := curIndex; i < n; i++ {
		swap(num, curIndex, i)
		permutation(num, n, curIndex+1, res)
		swap(num, i, curIndex)
	}
}

type P struct {
}

func (p *P) sayA() {
	fmt.Println("P.sayA")
	p.sayB()
}
func (p *P) sayB() {
	fmt.Println("P.sayB")
}

type S struct {
	P
}

func (s *S) sayB() {
	fmt.Println("S.sayB")
}

func testMap(m map[string]string) {
	m["a"] = "a"
	return
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer trace.Stop()

	ch := make(chan int, 100)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("func 1")
			}
		}()
		close(ch)
	}()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("func 2")
			}
		}()
		for {
			i := 0
			i++
		}
	}()
	time.Sleep(time.Second)

	return

	//a := []int{1, 2, 3, 4, 5}
	//b := []int{6, 7}
	//copy(a, b)
	//fmt.Println(a)
	//time.Sleep(time.Second)
	//a := []int{1, 2, 3, 4}
	//n := len(a)
	//res := [][]int{}
	//permutation(a, n, 0, &res)
	//fmt.Println(len(res), res)

}
