package main

import (
	"fmt"
)

func mq() {

}
func main() {
	ch := make(chan int, 100)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- 1
		}
		close(ch)
	}()
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()

	//

}
