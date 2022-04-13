package main

import (
	"fmt"
	"math"
	"strings"
)

func containStr(s string, sub string) {
	mapA := make(map[byte]int)
	for i := 0; i < len(sub); i++ {
		if _, ok := mapA[sub[i]]; ok {
			mapA[sub[i]]++
		} else {
			mapA[sub[i]] = 1
		}
	}
	for i := 0; i < len(s); i++ {
		if _, ok := mapA[s[i]]; ok {
			mapA[s[i]]--
		}
	}
}

func minWindow(s []byte, t []byte) int {
	l, r := 0, 1
	minLen := math.MaxInt32
	for ; r < len(s); r++ {
		if !strings.Contains(string(s[l:r]), string(t)) {
			continue
		}
		minLen = minInt(r-l, minLen)
		for ; l < r; l++ {
			if strings.Contains(string(s[l:r]), string(t)) {
				continue
			}
			minLen = minInt(r-l+1, minLen)
			break
		}
	}
	return minLen
}
func main() {
	s := "abccadde"
	t := "add"
	fmt.Println(minWindow([]byte(s), []byte (t)))
}
