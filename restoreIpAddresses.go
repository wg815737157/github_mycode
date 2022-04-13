package main

import (
	"fmt"
	"strconv"
)

func recursion(s string, n int, element string, res *[]string) {
	if n == 0 {
		if len(s) == 0 {
			*res = append(*res, element[:len(element)-1])
		}
		return
	}
	if s == "" {
		return
	}
	if s[0] == '0' {
		element += string(s[0]) + "."
		recursion(s[1:], n-1, element, res)
		return
	}

	for i := 0; i < len(s); i++ {
		tmp, _ := strconv.Atoi(s[:i+1])
		if tmp <= 255 && tmp >= 0 {
			newElement := element + s[:i+1] + "."
			recursion(s[i+1:], n-1, newElement, res)
		}
	}
}

//["255.255.11.135","255.255.111.35"]
func restoreIpAddresses(s string) []string {
	var res []string
	recursion(s, 4, "", &res)
	return res
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
