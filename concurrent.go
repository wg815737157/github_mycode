package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var channelA chan int
var muStruct *MuStruct

type MuStruct struct {
	mu     *sync.Mutex
	closed bool
}

func producerFunc(ctxProducer context.Context, ctxProducerCancel func()) {
	wg := &sync.WaitGroup{}
	go func() {
		<-time.After(5 * time.Second)
		fmt.Println("producer cancel")
		ctxProducerCancel()
	}()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctxProducer.Done():
					fmt.Println("Producer exit")
					return
				case channelA <- 1:
				}
			}
		}()
	}
	wg.Wait()

	muStruct.mu.Lock()
	if muStruct.closed == false {
		close(channelA)
		muStruct.closed = true
	}
	muStruct.mu.Unlock()

}
func consumerFunc(ctxProducerCancel func()) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range channelA {
				fmt.Println("consumer 消费", v)
			}
			return
		}()
	}
	<-time.After(6 * time.Second)
	fmt.Println("consumer cancel")
	ctxProducerCancel()
	wg.Wait()
}

func main() {

	channelA = make(chan int, 10)
	muStruct = &MuStruct{&sync.Mutex{}, false}
	ctxProducer, producerCancel := context.WithCancel(context.Background())
	go producerFunc(ctxProducer, producerCancel)
	go consumerFunc(producerCancel)

	time.Sleep(10 * time.Second)
}
