package main

import "fmt"

func umuTest(a, b []byte) []byte {
	res := []byte{}
	n1 := len(a)
	n2 := len(b)
	if n1 >= n2 {
		for i := 0; i < n2; i++ {
			for j := 0; j < n1; j++ {
				if b[j] != a[i] {

				}
			}
		}
	} else {
		for i := 0; i < n1; i++ {

		}
	}

	return res

}

func main() {

	a := "abcdefg"
	b := "DDDdef"
	fmt.Println(string(umuTest([]byte(a), []byte(b))))

}
