package main

import (
	"fmt"
	"sort"
)

type interval [][]int

func (in interval) Len() int {
	return len(in)
}

func (in interval) Less(i, j int) bool {
	return in[i][0] < in[j][0]
}

func (in interval) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}
func main() {
	a :=[][]int{{1,3},{2,3},{4,5},{1,2},{1,6}}
	sort.Sort(interval(a))
	fmt.Println(a)
}
