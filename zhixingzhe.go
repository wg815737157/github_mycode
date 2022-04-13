package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	endCh1 := make(chan int)
	endCh2 := make(chan int)
	i := 1
	go func() {
		for {
			select {
			case <-ch1:
				fmt.Println(0)
				ch2 <- 0
			case <-endCh2:
				fmt.Println("0 协程退出")
				return
			}

		}

	}()
	go func() {
		for {
			<-ch2
			fmt.Println(i)
			i++
			if i > 100 {
				endCh1 <- 0
				endCh2 <- 0
				return
			}
			ch1 <- 0
		}
	}()

	ch1 <- 0
	<-endCh1
	time.Sleep(time.Second)
}
