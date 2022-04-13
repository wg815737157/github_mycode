package main

import "fmt"

var res []string

func swapByte(s []byte, i int, j int) {
	s[i], s[j] = s[j], s[i]
	return
}
func quan(s []byte, index int) {
	if index == len(s) {
		element := make([]byte, len(s))
		copy(element, s)
		res = append(res, string(element))
		return
	}
	for i := index; i < len(s); i++ {
		swapByte(s, index, i)
		quan(s, index+1)
		swapByte(s, i, index)
	}
}
func main() {
	s := "abc"
	quan([]byte(s), 0)
	fmt.Println(res)
}
