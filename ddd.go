package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime/trace"
	"syscall"
	"time"
)

var ch chan int

func a() {
	time.Sleep(10 * time.Millisecond)
	ch <- 1
}

func b() {
	<-ch
}

func main() {
	err := trace.Start(os.Stderr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()
	fd, err := syscall.Open("1.txt", 0666, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	data, err := syscall.Mmap(fd, 0, 10, syscall.PROT_WRITE|syscall.PROT_READ, syscall.MAP_SHARED)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range data {
		fmt.Println(string(v))
	}
	data[0] = '0' + 9
	for _, v := range data {
		fmt.Println(string(v))
	}
	err = syscall.Munmap(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	ch = make(chan int)
	go a()

	go b()
	l, err := net.Listen("tcp", "127.0.0.1")
	if err != nil {
		fmt.Println(err)
		return
	}
	funcA := http.Handle("/", ServeHTTP)
	err = http.Serve(l, funcA)
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(time.Second)
}
