package lfucache

// 460. LFU 缓存
// https://leetcode.cn/problems/lfu-cache/
// 设计一个 LFU（最不经常使用）缓存：容量满时淘汰「使用次数最少」的 key；
// 若有多个 key 使用次数相同，淘汰其中「最久未使用」的那个。get 和 put 要求平均 O(1)。
//
// 核心思路：哈希表 + 按频次分组的双向链表 + minFreq。
//   - keyToNode：key -> 节点（节点里存 key、value、freq）；
//   - freqToList：freq -> 一条双向链表，同频次的节点按访问时间排列（头新尾旧）；
//   - minFreq：当前最小频次，淘汰时直接取 freqToList[minFreq] 的尾部节点。
// 每次访问把节点从 freq 链表摘下，freq+1 后插入新链表头部；
// 若旧链表空了且它正是 minFreq，minFreq 加 1。新插入的 key 频次为 1，minFreq 重置为 1。

type node struct {
	key, value, freq int
	prev, next       *node
}

// dlist 带哨兵的双向链表。
type dlist struct {
	head, tail *node
	size       int
}

func newDlist() *dlist {
	head, tail := &node{}, &node{}
	head.next = tail
	tail.prev = head
	return &dlist{head: head, tail: tail}
}

func (l *dlist) pushFront(n *node) {
	n.prev = l.head
	n.next = l.head.next
	l.head.next.prev = n
	l.head.next = n
	l.size++
}

func (l *dlist) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
	l.size--
}

// LFUCache 解法一：哈希表 + 频次双向链表 + minFreq ⭐ 最优方案
// Get：命中则频次 +1 并返回值，否则 -1。
// Put：已存在则更新值并频次 +1；不存在时若已满，先淘汰 minFreq 链表尾部，再插入频次为 1 的新节点。
// 时间复杂度 O(1)（每个操作），空间复杂度 O(capacity)。
type LFUCache struct {
	capacity   int
	minFreq    int
	keyToNode  map[int]*node
	freqToList map[int]*dlist
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		capacity:   capacity,
		keyToNode:  map[int]*node{},
		freqToList: map[int]*dlist{},
	}
}

func (this *LFUCache) Get(key int) int {
	n, ok := this.keyToNode[key]
	if !ok {
		return -1
	}
	this.touch(n)
	return n.value
}

func (this *LFUCache) Put(key, value int) {
	if this.capacity <= 0 {
		return
	}
	if n, ok := this.keyToNode[key]; ok {
		n.value = value
		this.touch(n)
		return
	}
	if len(this.keyToNode) == this.capacity {
		// 淘汰最小频次中最久未使用的节点。
		l := this.freqToList[this.minFreq]
		victim := l.tail.prev
		l.remove(victim)
		if l.size == 0 {
			delete(this.freqToList, this.minFreq)
		}
		delete(this.keyToNode, victim.key)
	}
	n := &node{key: key, value: value, freq: 1}
	this.keyToNode[key] = n
	this.listOf(1).pushFront(n)
	this.minFreq = 1
}

// touch 把节点频次 +1，移动到新频次链表的头部。
func (this *LFUCache) touch(n *node) {
	l := this.freqToList[n.freq]
	l.remove(n)
	if l.size == 0 {
		delete(this.freqToList, n.freq)
		if this.minFreq == n.freq {
			this.minFreq++
		}
	}
	n.freq++
	this.listOf(n.freq).pushFront(n)
}

func (this *LFUCache) listOf(freq int) *dlist {
	l, ok := this.freqToList[freq]
	if !ok {
		l = newDlist()
		this.freqToList[freq] = l
	}
	return l
}
