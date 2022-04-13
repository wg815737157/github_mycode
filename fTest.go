package main

import (
	"fmt"
	"os"
)

func main() {
	//c := make(chan int)

	f, err := os.OpenFile("jd.go", os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("DD")
		fmt.Println(err)
		return
	}

	defer f.Close()
	n, err := f.Write([]byte("wg"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)


}
