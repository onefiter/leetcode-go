package frontmiddleback

import "container/list"

// 1670. 设计前中后队列
// https://leetcode.cn/problems/design-front-middle-back-queue/
// 实现一个队列，支持在前、中、后三个位置 push 和 pop。
// 「中间」的约定：长度为偶数时有两个中间位置，选靠前的那个。
// 例如 [1,2,3,4,5] PushMiddle(6) -> [1,2,6,3,4,5]；[1,2,3,4,5,6] PopMiddle() 弹出 3。
// 队列为空时 Pop 返回 -1。
//
// 核心思路：把队列拆成左右两半 left、right（各自是双端队列），保持
// len(right) - len(left) ∈ {0, 1}，那么「中间」永远在两半的交界处：
//   - PushMiddle：长度相等时放到 right 头部，否则放到 left 尾部；
//   - PopMiddle：长度相等时弹 left 尾部，否则弹 right 头部；
//   - 其余操作改完一端后，用 balance 挪一个元素恢复不变式。

// FrontMiddleBackQueue 解法一：两个双端队列保持平衡 ⭐ 最优方案
// 双端队列用标准库 container/list（双向链表），两端增删都是 O(1)。
// 时间复杂度 O(1)（每个操作），空间复杂度 O(n)。
type FrontMiddleBackQueue struct {
	left, right *list.List
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{left: list.New(), right: list.New()}
}

// balance 恢复不变式 len(right)-len(left) ∈ {0, 1}。
func (this *FrontMiddleBackQueue) balance() {
	if this.left.Len() > this.right.Len() {
		this.right.PushFront(this.left.Remove(this.left.Back()))
	} else if this.right.Len() > this.left.Len()+1 {
		this.left.PushBack(this.right.Remove(this.right.Front()))
	}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.left.PushFront(val)
	this.balance()
}

func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	if this.left.Len() == this.right.Len() {
		this.right.PushFront(val)
	} else {
		this.left.PushBack(val)
	}
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.right.PushBack(val)
	this.balance()
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if this.right.Len() == 0 { // right 为空说明整个队列为空
		return -1
	}
	var v int
	if this.left.Len() > 0 {
		v = this.left.Remove(this.left.Front()).(int)
	} else {
		v = this.right.Remove(this.right.Front()).(int) // 只剩 1 个元素，在 right
	}
	this.balance()
	return v
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if this.right.Len() == 0 {
		return -1
	}
	if this.left.Len() == this.right.Len() {
		return this.left.Remove(this.left.Back()).(int)
	}
	return this.right.Remove(this.right.Front()).(int)
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if this.right.Len() == 0 {
		return -1
	}
	v := this.right.Remove(this.right.Back()).(int)
	this.balance()
	return v
}
