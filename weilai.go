package main

import (
	"fmt"
	"time"
)

// 在线面试平台。将链接分享给你的朋友以加入相同的房间。
// Author: tdzl2003<dengyun@meideng.net>
// QQ-Group: 839088532
// test

//是否完成：否

//请你设计并实现一个带超时的LRU (最近最少使用) 缓存 约束的数据结构。
//实现 LRUCache 类：
//LRUCache(int capacity, int timeoutSec) 以 正整数 作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
// 在线面试平台。将链接分享给你的朋友以加入相同的房间。
// Author: tdzl2003<dengyun@meideng.net>
// QQ-Group: 839088532
// test

//是否完成：否

//请你设计并实现一个带超时的LRU (最近最少使用) 缓存 约束的数据结构。
//实现 LRUCache 类：
//LRUCache(int capacity, int timeoutSec) 以 正整数 作为容量 capacity 初始化 LRU 缓存
//int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
//void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组 key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
type lRUCache struct {
	timeoutSec int64
	len        int
	capacity   int
	head       *linkedNode
	tail       *linkedNode
	cashMap    map[int]*linkedNode
}

type linkedNode struct {
	expireTime int64
	key        int
	value      int
	pre        *linkedNode
	next       *linkedNode
}

func (l *lRUCache) initLinkedNode(key, value int) *linkedNode {
	expireTime := time.Now().Unix() + l.timeoutSec
	node := &linkedNode{expireTime, key, value, nil, nil}
	return node
}
func LRUCache(capacity int, timeoutSec int64) *lRUCache {
	l := &lRUCache{timeoutSec, 0, capacity, nil, nil, make(map[int]*linkedNode)}
	head := &linkedNode{0, 0, 0, nil, nil}
	tail := &linkedNode{0, 0, 0, nil, nil}
	head.next = tail
	tail.pre = head
	l.head = head
	l.tail = tail
	return l
}

func (l *lRUCache) deleteNode(node *linkedNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	delete(l.cashMap, node.key)
	l.len--
}

func (l *lRUCache) moveToHead(node *linkedNode) {
	//更新时间
	node.expireTime = time.Now().Unix() + l.timeoutSec
	//删除节点
	node.pre.next = node.next
	node.next.pre = node.pre
	//移动到头
	node.pre = l.head
	node.next = l.head.next
	node.next.pre = node
	node.pre.next = node
	l.cashMap[node.key] = node
}

func (l *lRUCache) insertHead(node *linkedNode) {
	node.pre = l.head
	node.next = l.head.next
	node.next.pre = node
	node.pre.next = node
	l.cashMap[node.key] = node
	l.len++
}
func (l *lRUCache) deleteTail() {
	node := l.tail.pre
	node.pre.next = l.tail
	l.tail.pre = node.pre
	delete(l.cashMap, node.key)
	l.len--
}

func (l *lRUCache) Get(key int) int {
	if _, ok := l.cashMap[key]; !ok {
		return -1
	}
	//大于过期时间删除
	if l.cashMap[key].expireTime < time.Now().Unix() {
		l.deleteNode(l.cashMap[key])
		return -1
	}
	//更新过期时间 移到头
	l.moveToHead(l.cashMap[key])
	return l.cashMap[key].value
}

func (l *lRUCache) Put(key int, value int) {
	//存在
	if _, ok := l.cashMap[key]; ok {
		l.moveToHead(l.cashMap[key])
		return
	}
	//不存在
	node := l.initLinkedNode(key, value)
	l.insertHead(node)
	if l.len > l.capacity {
		for _, v := range l.cashMap {
			if v.expireTime >= time.Now().Unix() {
				l.deleteNode(v)
			}
		}
		if l.len>l.capacity{
			l.deleteTail()
		}
	}
}

func main() {
	l := LRUCache(2, 10)

	//  协程检查过期的key
	go func() {
		for {
			for _, v := range l.cashMap {
				if v.expireTime >= time.Now().Unix() {
					l.deleteNode(v)
				}
			}
			time.Sleep(time.Second)
		}
	}()
	l.Put(1, 1)
	//1
	fmt.Println(l.Get(1))
	//-1
	fmt.Println(l.Get(2))
	l.Put(2, 2)
	//2
	fmt.Println(l.Get(2))
	l.Put(3, 3)
	//1，1 退出 返回-1
	fmt.Println(l.Get(1))
}
