package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	var s []int
	for i := 0; i < len(tokens); i++ {
		switch tokens[i] {
		case "+":
			s[len(s)-2] += s[len(s)-1]
			s = s[:len(s)-1]
		case "-":
			s[len(s)-2] -= s[len(s)-1]
			s = s[:len(s)-1]
		case "*":
			s[len(s)-2] *= s[len(s)-1]
			s = s[:len(s)-1]
		case "/":
			s[len(s)-2] /= s[len(s)-1]
			s = s[:len(s)-1]
		default:
			v, _ := strconv.Atoi(tokens[i])
			s = append(s, v)
		}
	}
	return s[0]
}

func main() {
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))

}
