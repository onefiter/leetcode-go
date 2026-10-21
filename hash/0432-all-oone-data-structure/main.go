package allone

// 432. 全 O(1) 的数据结构
// https://leetcode.cn/problems/all-oone-data-structure/
// 设计 AllOne 存储字符串计数：Inc(key) 计数加 1（不存在则插入计数 1）；
// Dec(key) 计数减 1，减到 0 时删除（保证调用时 key 存在）；
// GetMaxKey / GetMinKey 返回任意一个计数最大 / 最小的 key，为空时返回 ""。所有操作 O(1)。
//
// 核心思路：哈希表 + 「计数桶」双向链表。
//   - 链表每个节点是一个桶，存同一计数 count 的所有 key，链表按 count 从小到大严格递增；
//   - 哈希表 key -> 所在桶，O(1) 定位；
//   - Inc/Dec 时 key 只会移动到相邻计数的桶：相邻桶计数恰好是 count±1 就直接放入，否则就地新建；
//     原桶空了立即删掉，保证 head.next 是最小桶、tail.prev 是最大桶。

// bucket 计数桶：同一计数的 key 集合。
type bucket struct {
	count      int
	keys       map[string]struct{}
	prev, next *bucket
}

// AllOne 解法一：哈希表 + 计数桶双向链表 ⭐ 最优方案
// 时间复杂度：Inc / Dec / GetMaxKey / GetMinKey 均为 O(1)；空间复杂度 O(n)。
type AllOne struct {
	nodes      map[string]*bucket
	head, tail *bucket // 哨兵：head.next 计数最小，tail.prev 计数最大
}

func Constructor() AllOne {
	head, tail := &bucket{}, &bucket{}
	head.next = tail
	tail.prev = head
	return AllOne{nodes: map[string]*bucket{}, head: head, tail: tail}
}

func (this *AllOne) Inc(key string) {
	cur, ok := this.nodes[key]
	if !ok {
		cur = this.head // 视为从计数 0 的 head 哨兵移出
	}
	nxt := cur.next
	if nxt == this.tail || nxt.count != cur.count+1 {
		nxt = this.insertAfter(cur, cur.count+1)
	}
	nxt.keys[key] = struct{}{}
	this.nodes[key] = nxt
	if ok {
		this.removeKey(cur, key)
	}
}

func (this *AllOne) Dec(key string) {
	cur, ok := this.nodes[key]
	if !ok {
		return
	}
	if cur.count == 1 {
		delete(this.nodes, key)
	} else {
		prev := cur.prev
		if prev == this.head || prev.count != cur.count-1 {
			prev = this.insertAfter(prev, cur.count-1)
		}
		prev.keys[key] = struct{}{}
		this.nodes[key] = prev
	}
	this.removeKey(cur, key)
}

func (this *AllOne) GetMaxKey() string {
	return anyKey(this.tail.prev)
}

func (this *AllOne) GetMinKey() string {
	return anyKey(this.head.next)
}

// insertAfter 在 b 之后插入一个计数为 count 的新桶并返回。
func (this *AllOne) insertAfter(b *bucket, count int) *bucket {
	nb := &bucket{count: count, keys: map[string]struct{}{}, prev: b, next: b.next}
	b.next.prev = nb
	b.next = nb
	return nb
}

// removeKey 从桶 b 中删除 key，桶空了就把桶从链表摘掉。
func (this *AllOne) removeKey(b *bucket, key string) {
	delete(b.keys, key)
	if len(b.keys) == 0 {
		b.prev.next = b.next
		b.next.prev = b.prev
	}
}

// anyKey 返回桶中任意一个 key；哨兵桶（keys 为 nil）返回 ""。
// 桶内 map 只取第一个元素即返回，是 O(1)。
func anyKey(b *bucket) string {
	for k := range b.keys {
		return k
	}
	return ""
}

// AllOneMap 解法二：只用哈希表，查询时扫描找最值（对照用）
// 时间复杂度：Inc / Dec O(1)，GetMaxKey / GetMinKey O(n)；空间复杂度 O(n)。
type AllOneMap struct {
	cnt map[string]int
}

func NewAllOneMap() AllOneMap {
	return AllOneMap{cnt: map[string]int{}}
}

func (this *AllOneMap) Inc(key string) {
	this.cnt[key]++
}

func (this *AllOneMap) Dec(key string) {
	if c, ok := this.cnt[key]; ok {
		if c == 1 {
			delete(this.cnt, key)
		} else {
			this.cnt[key] = c - 1
		}
	}
}

func (this *AllOneMap) GetMaxKey() string {
	best, bestCnt := "", 0
	for k, c := range this.cnt {
		if c > bestCnt {
			best, bestCnt = k, c
		}
	}
	return best
}

func (this *AllOneMap) GetMinKey() string {
	best, bestCnt := "", 0
	for k, c := range this.cnt {
		if bestCnt == 0 || c < bestCnt {
			best, bestCnt = k, c
		}
	}
	return best
}
