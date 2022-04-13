package main

import (
	"fmt"
)

func main() {

	fmt.Println(solution(101, 18))

}
func solution(m int, n int) string {
	res := ""
	for m != 0 {
		ch := m % n
		if ch >= 10 {
			ch = 'A' + ch - 10
		} else {
			ch = '0' + ch
		}
		res = string(byte(ch)) + res
		m = m / n
	}
	return res
}

func solution2(m int) int {
	n := 0
	for m != 0 {
		n = 10*n + m%10
		m = m / 10
	}
	return n
}
