package main

import (
	"fmt"
	"sync"
	"time"
)

var ipControl *IpControl

type IpInfo struct {
	IP       string
	IpTicker *time.Ticker
	IpCount  int
}
type IpControl struct {
	rwMutex sync.RWMutex
	IpInfo  map[string]*IpInfo
}

func checkIp(ip string) {
	successCount := 0
	ipControl.rwMutex.Lock()
	defer ipControl.rwMutex.Unlock()

	//不存在插入并计数1，初始化时间3分钟计时器
	if _, ok := ipControl.IpInfo[ip]; !ok {
		ipInfo := &IpInfo{ip, time.NewTicker(3 * time.Minute), 0}
		ipInfo.IpCount++
		ipControl.IpInfo[ip] = ipInfo
		return
	}
	//存在
	for {
		select {
		//超过3分钟可以访问并且计数器+1
		case <-ipControl.IpInfo[ip].IpTicker.C:
			ipControl.IpInfo[ip].IpCount++
			if ipControl.IpInfo[ip].IpCount >= 100 {
				successCount++
				fmt.Println("success:", successCount)
			}
			return
		//还在3分钟之内
		default:
			fmt.Println("3分钟访问限制中")
			return
		}
	}
}

func main() {
	time.Sleep(time.Second)
}
