package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func proc() {
	panic("ok")
}
func logicFunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:" + strconv.Itoa(int(time.Now().Unix())))
		}
	}()
	proc()
}

func solution1() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			go func() {
				logicFunc()
			}()
		case <-time.After(10 * time.Second):
			return
		}
	}
}



func solution5(nums []int) {

}

func main() {
	solution1()
	select {}
}
