package main

import (
	"fmt"
	"math"
)

func solution1Bak(n int, a [][]int) int {
	inDegree, outDegree := make([]int, n+1), make([]int, n+1)
	for i := 0; i < len(a); i++ {
		inDegree[a[i][1]]++
		outDegree[a[i][0]]++
	}
	for i := 1; i <= n; i++ {
		if inDegree[i] == n-1 && outDegree[i] == 0 {
			return i
		}
	}
	return -1
}

func solution(n int, a [][]int) int {
	trustMap := make(map[int]struct{})
	beTrustedMap := make(map[int]int)
	aLen := len(a)
	for i := 0; i < aLen; i++ {
		trustMap[a[i][0]] = struct{}{}
		if _, ok := beTrustedMap[a[i][1]]; ok {
			beTrustedMap[a[i][1]]++
		} else {
			beTrustedMap[a[i][1]] = 1
		}
	}
	if len(trustMap) != n-1 {
		return -1
	}
	for i := 1; i <= n; i++ {
		if _, ok := trustMap[i]; ok {
			continue
		}
		if beTrustedMap[i] == n-1 {
			return i
		}
	}
	return -1
}

func solution2(s []string) []byte {
	initSlice := make([]map[byte]int, len(s))

	minColumn := math.MaxInt32
	column := 0
	for i := 0; i < len(s); i++ {
		if len(initSlice[i]) == 0 {
			initSlice[i] = make(map[byte]int)
		}
		n := len(s[i])
		if minColumn > n {
			minColumn = n
			column = i
		}
		for j := 0; j < n; j++ {
			if _, ok := initSlice[i][s[i][j]]; ok {
				initSlice[i][s[i][j]]++
			} else {
				initSlice[i][s[i][j]] = 1
			}
		}
	}
	res := []byte{}

	for i := 0; i < len(s[column]); i++ {
		k := s[column][i]
		ifErr := false
		for i := 0; i < len(initSlice); i++ {
			if _, ok := initSlice[i][k]; !ok {
				ifErr = true
				break
			}
			initSlice[i][k]--
			if initSlice[i][k] == 0 {
				delete(initSlice[i], k)
			}
		}
		if ifErr {
			continue
		}
		res = append(res, k)
	}
	return res
}

func solution3(a []int) int {
	maxNum := make([]int, len(a)+1)
	continueSum := 0
	continueNum := 0
	for i := 1; i < len(maxNum); i++ {
		if a[i-1] == 0 {
			continueSum = 0
			continueNum = 0
			maxNum[i] = maxNum[i-1]
			continue
		}
		if a[i-1] >= 8 {
			continueSum++
			continueNum++
		} else {
			continueSum--
			continueNum++
		}
		if continueSum > 0 {
			if continueNum > maxNum[i-1] {
				maxNum[i] = continueNum
			}
			continue
		}
		maxNum[i] = maxNum[i-1]
	}
	return maxNum[len(a)]
}

func solution3Bak(a []int) int {
	maxNum := make([]int, len(a)+1)
	continueSum := 0
	continueNum := 0
	for i := 1; i < len(maxNum); i++ {

	}
	return maxNum[len(a)]
}
func main() {
	hours := []int{9, 9, 6, 0, 6, 6, 9}
	fmt.Println(solution3(hours))
}
