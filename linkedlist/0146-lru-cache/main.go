package lrucache

// 146. LRU 缓存
// https://leetcode.cn/problems/lru-cache/
// 设计一个容量为 capacity 的 LRU（最近最少使用）缓存，get 和 put 都要求平均 O(1)。
// 容量满时再插入新 key，要先淘汰「最久没被访问」的那个 key。
//
// 核心思路：哈希表 + 双向链表。
//   - 哈希表 key -> 链表节点，O(1) 定位；
//   - 双向链表按「访问时间」排序：越靠近头部越新，尾部是最久未使用的；
//   - 每次 get/put 命中就把节点移到头部；容量超了就删尾部节点。
// 用两个哨兵节点 head/tail，避免处理空链表和边界的特判。

// dlinkedNode 双向链表节点，同时存 key，淘汰尾部时才能从哈希表里删掉对应 key。
type dlinkedNode struct {
	key, value int
	prev, next *dlinkedNode
}

// LRUCache 解法一：哈希表 + 手写双向链表 ⭐ 最优方案
// Get：查哈希表，命中则把节点移到头部并返回值，否则返回 -1。
// Put：已存在则更新值并移到头部；不存在则新建节点插到头部，超容量就删掉尾部节点。
// 时间复杂度 O(1)（每个操作），空间复杂度 O(capacity)。
type LRUCache struct {
	capacity   int
	cache      map[int]*dlinkedNode
	head, tail *dlinkedNode // 哨兵：head.next 是最新的，tail.prev 是最旧的
}

func Constructor(capacity int) LRUCache {
	head, tail := &dlinkedNode{}, &dlinkedNode{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*dlinkedNode, capacity),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.moveToHead(node)
	return node.value
}

func (this *LRUCache) Put(key, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
		return
	}
	node := &dlinkedNode{key: key, value: value}
	this.cache[key] = node
	this.addToHead(node)
	if len(this.cache) > this.capacity {
		// 淘汰最久未使用的尾部节点。
		last := this.tail.prev
		this.removeNode(last)
		delete(this.cache, last.key)
	}
}

// addToHead 把节点插到 head 哨兵之后。
func (this *LRUCache) addToHead(node *dlinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

// removeNode 把节点从链表中摘下。
func (this *LRUCache) removeNode(node *dlinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *dlinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}
