package circulardeque

// 641. 设计循环双端队列
// https://leetcode.cn/problems/design-circular-deque/
// 设计容量为 k 的循环双端队列，两端都能插入、删除、查看。
// 空时 GetFront/GetRear 返回 -1；满时插入、空时删除返回 false。
//
// 核心思路：和 622 一样用定长数组当环，记录队头下标 head 和元素个数 size：
//   - 头部插入：head 往左挪一格 (head-1+k)%k，再写入；
//   - 尾部插入：写到 (head+size)%k；
//   - 头部删除：head 往右挪一格；尾部删除：只需 size--。

// MyCircularDeque 解法一：定长数组 + 头指针 + 元素计数 ⭐ 最优方案
// 时间复杂度 O(1)（每个操作），空间复杂度 O(k)。
type MyCircularDeque struct {
	data []int
	head int
	size int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{data: make([]int, k)}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	k := len(this.data)
	this.head = (this.head - 1 + k) % k // 头指针左移，+k 防止负数
	this.data[this.head] = value
	this.size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data[(this.head+this.size)%len(this.data)] = value
	this.size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.data)
	this.size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.head+this.size-1)%len(this.data)]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.size == len(this.data)
}
