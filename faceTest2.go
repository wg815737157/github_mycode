package main

import "fmt"

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}
type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkedNode
	head, tail *DLinkedNode
}

func constructor() *LRUCache {
	tmpMap := make(map[int]*DLinkedNode)
	head := &DLinkedNode{-1, -1, nil, nil}
	tail := &DLinkedNode{-1, -1, head, nil}
	head.next = tail
	lru := &LRUCache{0, 100, tmpMap, head, tail}
	return lru
}

func (lru *LRUCache) insert(node *DLinkedNode) {
	//存在移动到头
	if _, ok := lru.cache[node.key]; ok {
		lru.deleteNode(node)
		lru.insertNode(node)
	} else { //不存在插入
		lru.insertNode(node)
	}
}

func (lru *LRUCache) insertHead(node *DLinkedNode) {
	lru.head.next.prev = node
	node.next = lru.head.next
	lru.head.next = node
	node.prev = lru.head
	lru.cache[node.key] = node
	lru.size++
}

func (lru *LRUCache) insertNode(node *DLinkedNode) {
	//小于capacity 直接插入
	if lru.size < lru.capacity {
		lru.insertHead(node)
	} else { //大于等于capacity 删除尾 插入头
		lru.deleteTail()
		lru.insertHead(node)
	}
}

func (lru *LRUCache) deleteNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
	lru.size--
}

func (lru *LRUCache) deleteTail() {
	node := lru.tail.prev
	lru.deleteNode(node)
	delete(lru.cache, node.key)
}

func (lru *LRUCache) print() {
	node := lru.head.next
	for node != nil {
		if node.value != -1 {
			fmt.Println(node.value)
		}
		node = node.next
	}
}

func main() {
	lru := constructor()
	a := &DLinkedNode{1, 1, nil, nil}
	b := &DLinkedNode{2, 2, nil, nil}
	lru.insert(a)
	lru.insert(b)
	lru.print()
	lru.insert(a)
	lru.print()
}
