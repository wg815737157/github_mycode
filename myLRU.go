package main

import "fmt"

type DLinkNode struct {
	key, value int
	prev, next *DLinkNode
}
type LRU struct {
	size, capacity int
	cache          map[int]*DLinkNode
	head, tail     *DLinkNode
}

func construct() *LRU {
	l := LRU{0, 1, make(map[int]*DLinkNode), &DLinkNode{}, &DLinkNode{}}
	l.head.next = l.tail
	l.tail.prev = l.head
	return &l
}
func (lru *LRU) removeNode(node *DLinkNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}
func (lru *LRU) moveToHead(node *DLinkNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}
func (lru *LRU) get(key int) int {
	if _, ok := lru.cache[key]; !ok {
		return -1
	}
	lru.moveToHead(lru.cache[key])
	return lru.cache[key].value
}
func (lru *LRU) removeTail() *DLinkNode {
	node := lru.tail.prev
	lru.removeNode(node)
	return node
}
func initNode(key, value int) *DLinkNode {
	return &DLinkNode{key, value, nil, nil}
}
func (lru *LRU) addToHead(node *DLinkNode) {
	node.prev = lru.head
	node.next = lru.head.next
	lru.head.next.prev = node
	lru.head.next = node
}

func (lru *LRU) put(key int, value int) {
	//存在
	if _, ok := lru.cache[key]; ok {
		lru.cache[key].value = value
		lru.moveToHead(lru.cache[key])
		return
	}
	//不存在
	newNode := initNode(key, value)
	lru.cache[key] = newNode
	lru.addToHead(newNode)
	lru.size++

	if lru.size > lru.capacity {
		node := lru.removeTail()
		lru.size--
		delete(lru.cache, node.key)
		return
	}
}

func main() {
	lru := construct()
	lru.put(1, 1)
	lru.put(2, 1)
	v := lru.get(1)
	fmt.Println(v)
	v = lru.get(2)
	fmt.Println(v)
	lru.put(2, 2)
	v = lru.get(2)
	fmt.Println(v)
	lru.put(0, 1)
	v = lru.get(0)
	fmt.Println(v)
}
