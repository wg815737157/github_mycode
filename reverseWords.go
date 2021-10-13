package main

import (
	"fmt"
)

func reverseWords(s string) string {
	tmpS := []byte(s)
	l := len(tmpS)
	for i := 0; i < l; i++ {
		start := i
		for ; i < l && tmpS[i] != ' '; i++ {

		}
		for n := start; n < start+(i-start)/2; n++ {
			tmpS[n], tmpS[i-1+start-n] = tmpS[i-1+start-n], tmpS[n]
		}
	}
	return string(tmpS)
}
func main() {
	s := "Let's take LeetCode contest"
	fmt.Println(reverseWords(s))

}
