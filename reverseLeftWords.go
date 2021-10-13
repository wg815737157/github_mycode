package main

import "fmt"

func reverseLeftWords(s string, n int) string {
	sRune := []rune(s)
	resRune := append(sRune[n:], sRune[:n]...)
	return string(resRune)
}
func main() {

	s := "abkddessv"
	fmt.Println(string(reverseLeftWords(s, 2)))
}
