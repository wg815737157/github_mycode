package main

import (
	"fmt"
	"reflect"
)

type T struct{
	A int
}

type T2 struct{
	T
	A int
}


func replaceSpace(s string) string {
	n := len(s)
	res := make([]byte, 3*n)
	size := 0
	for i := 0; i < n; i++ {
		if s[i] == ' ' {
			res[size] = '%'
			size++
			res[size] = '2'
			size++
			res[size] = '0'
			size++
		} else {
			res[size] = s[i]
			size++
		}
	}
	return string(res[:size])
}
func main() {

	test:=&T2{}

	fmt.Println(reflect.TypeOf(test).String())
	return


	s := "a +ddd ddw"
	fmt.Println(replaceSpace(s))
}
