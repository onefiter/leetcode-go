package circularqueue

// 622. 设计循环队列
// https://leetcode.cn/problems/design-circular-queue/
// 设计一个容量为 k 的循环队列，支持 EnQueue、DeQueue、Front、Rear、IsEmpty、IsFull。
// 队列为空时 Front/Rear 返回 -1；入队满了、出队空了都返回 false。
//
// 核心思路：用一个长度为 k 的定长数组当「环」，
// 记录队头下标 head 和当前元素个数 size：
//   - 队尾下一个空位 = (head + size) % k；
//   - 出队只需 head = (head + 1) % k，size--；
//   - 用 size 判空/判满，避免「空和满都是 head==tail」的歧义。

// MyCircularQueue 解法一：定长数组 + 头指针 + 元素计数 ⭐ 最优方案
// 所有操作只做常数次下标运算。
// 时间复杂度 O(1)（每个操作），空间复杂度 O(k)。
type MyCircularQueue struct {
	data []int
	head int // 队头元素下标
	size int // 当前元素个数
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{data: make([]int, k)}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	tail := (this.head + this.size) % len(this.data) // 队尾后面的空位
	this.data[tail] = value
	this.size++
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % len(this.data)
	this.size--
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.data[(this.head+this.size-1)%len(this.data)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.size == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return this.size == len(this.data)
}
