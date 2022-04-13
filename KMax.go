package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 快排K大
func solution2(nums []int, k int) int {
	n := len(nums)
	if n <= 1 {
		return nums[0]
	}
	l, r := 0, 0
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Int31n(int32(n))
	nums[randomNum], nums[n-1] = nums[n-1], nums[randomNum]
	for r < n-1 {
		if nums[r] > nums[n-1] {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
		r++
	}
	nums[l], nums[n-1] = nums[n-1], nums[l]
	if l+1 == k {
		return nums[l]
	}
	if l+1 > k {
		return solution2(nums[:l], k)
	}

	return solution2(nums[l+1:], k-l-1)
}

//
//// 循环 实现方法 冒泡
//func bubbleSort(a []int) []int {
//	n := len(a)
//	for i := 0; i < n; i++ {
//		for j := 1; j < len(a)-i; j++ {
//			if a[j] < a[j-1] {
//				a[j], a[j-1] = a[j-1], a[j]
//			}
//		}
//	}
//	return a
//}
//
////递归实现方法-归并
//
//func (m *MergeSortStruct) Initial(inputArray []int) {
//	m.tmp = make([]int, len(inputArray))
//}
//



//
//func findKMax2_2(a []int) []int {
//	if len(a) <= 2 {
//		return a
//	}
//	quickSort(a)
//	return a[:2]
//}

//func solution(n int) {
//	if n%2 != 1 {
//		return
//	}
//	l, r := n/2, n/2
//	for l >= 0 {
//		for i := 0; i < n; i++ {
//			if i >= l && i <= r {
//				fmt.Printf("*")
//			} else {
//				fmt.Printf(" ")
//			}
//		}
//		fmt.Printf("\n")
//		l--
//		r++
//	}
//	l += 2
//	r -= 2
//	for l <= n/2 {
//		for i := 0; i < n; i++ {
//			if i >= l && i <= r {
//				fmt.Printf("*")
//			} else {
//				fmt.Printf(" ")
//			}
//		}
//		fmt.Printf("\n")
//		l++
//		r--
//	}
//}
func func1(ctx context.Context, f func()) {
	f()
	sonCtx, cancel := context.WithCancel(ctx)

	func1_1(sonCtx)

	cancel()

}

func func1_1(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(time.Second):
		}
	}
}

func main() {


	//lis, err := net.Listen("tcp", ":8080")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//s := grpc.NewServer()
	//if err := s.Serve(lis); err != nil {
	//	fmt.Println(err)
	//	return
	//}

}
