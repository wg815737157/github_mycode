package main

import "fmt"

func validPalindromeTmp(s []byte) bool {
	n := len(s)
	i := 0
	j := n - 1
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func validPalindrome(s string) bool {
	b := []byte(s)
	n := len(b)
	i := 0
	j := n - 1

	for i < j {
		if b[i] != b[j] {
			fmt.Println(string(b[i+1:j+1]), string(b[i:j]))
			return validPalindromeTmp(b[i+1:j+1]) || validPalindromeTmp(b[i:j])
		}
		i++
		j--
	}
	return true
}

func main() {

	s := "abc"
	fmt.Println(validPalindrome(s))

}
