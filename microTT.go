package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan int, 100)
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
		close(ch)
	}()
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			time.Sleep(time.Second)
			v, ok := <-ch
			if !ok {
				fmt.Println("closed:", v)
				return
			}
			fmt.Println(v)
		}()
	}
	wg.Wait()
}
