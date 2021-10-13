package main

import (
	"fmt"
	"strconv"
)

func numDecodings(runeArray []rune) int {
	n := len(runeArray)
	if n == 0 {
		return 0
	}
	if string(runeArray[0]) == "0" {
		return 0
	}

	f := make([]int, n)
	f[0] = 1

	for i := 1; i < n; i++ {
		//10012
		if string(runeArray[i]) == "0" {
			if string(runeArray[i-1]) == "1" || string(runeArray[i-1]) == "2" {
				if i-2 < 0 {
					f[i] = 1
				} else {
					f[i] = f[i-2]
				}
			} else {
				f[i] = 0
			}
			continue
		}
		if string(runeArray[i-1]) == "0" {
			f[i] = f[i-1]
			continue
		}
		v, _ := strconv.Atoi(string(runeArray[i-1 : i+1]))
		if v <= 26 {
			if i-2 < 0 {
				f[i] = f[i-1] + 1
			} else {
				f[i] = f[i-1] + f[i-2]
			}
		} else {
			f[i] = f[i-1]
		}
		fmt.Println(i, f[i])
	}
	return f[n-1]
}

func numDecodingsStandard(s string) int {
	n := len(s)
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; i <= n; i++ {
		if s[i-1] != '0' {
			f[i] += f[i-1]
		}
		if i > 1 && s[i-2] != '0' && ((s[i-2]-'0')*10+(s[i-1]-'0') <= 26) {
			f[i] += f[i-2]
		}
	}
	return f[n]
}

func main() {
	s := "2101"
	runeArray := []rune(s)
	fmt.Println(numDecodings(runeArray))
	fmt.Println(numDecodingsStandard(s))
}
