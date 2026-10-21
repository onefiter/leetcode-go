// Package lrucache 用标准库实现一个线程安全的 LRU 缓存。
// LRU = Least Recently Used，最近最少使用淘汰。核心数据结构：
//   - map[key]*list.Element：O(1) 定位到链表节点；
//   - container/list 双向链表：越靠近队首越「新」，队尾是最久未使用的。
//
// Get 命中时把节点移到队首；Put 时若 key 已存在则更新并移到队首，
// 若不存在则插入队首，超过容量时淘汰队尾。
// 时间 O(1)，空间 O(capacity)。
package lrucache

import (
	"container/list"
	"sync"
)

type entry struct {
	key   string
	value int
}

type LRUCache struct {
	capacity int
	mu       sync.Mutex
	ll       *list.List // 双向链表，front 最新，back 最旧
	items    map[string]*list.Element
}

func New(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		ll:       list.New(),
		items:    make(map[string]*list.Element, capacity),
	}
}

// Get 返回 key 对应的值；不存在返回 -1（沿用 LeetCode 约定）。
// 命中视为一次「使用」，把节点移到队首。
func (c *LRUCache) Get(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.items[key]; ok {
		c.ll.MoveToFront(el)
		return el.Value.(*entry).value
	}
	return -1
}

// Put 写入或更新 key，超出容量时淘汰最久未使用的项。
func (c *LRUCache) Put(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.items[key]; ok { // 已存在：更新 + 移到队首
		el.Value.(*entry).value = value
		c.ll.MoveToFront(el)
		return
	}
	el := c.ll.PushFront(&entry{key: key, value: value})
	c.items[key] = el
	if c.ll.Len() > c.capacity { // 超容量：淘汰队尾（最久未使用）
		oldest := c.ll.Back()
		c.ll.Remove(oldest)
		delete(c.items, oldest.Value.(*entry).key)
	}
}
