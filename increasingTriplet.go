package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	first := math.MaxInt32
	second := math.MaxInt32

	for i := 0; i < len(nums); i++ {
		if nums[i] == first || nums[i] == second {
			continue
		}
		if first > nums[i] {
			first = nums[i]
		} else if second < nums[i] {
			return true
		} else {
			second = nums[i]
		}
	}
	return false
}

func main() {

	r:=[]rune{}
	r=append(r,rune(49),rune(50),rune(51))
	fmt.Println(string(r))
	l:=len(r)/2
	for i:=0;i<l;i++{
		r[i],r[len(r)-1-i]=r[len(r)-i-1],r[i]
	}

	b:=[]byte{}
	b=append(b,'1','2','3')
	l=len(b)/2
	for i:=0;i<l;i++{
		b[i],b[len(b)-1-i]=b[len(b)-i-1],b[i]
	}
	fmt.Println(b)
	fmt.Println(r)
	return

	//fmt.Println(increasingTriplet([]int{1, 1, -2, 6}))

}
