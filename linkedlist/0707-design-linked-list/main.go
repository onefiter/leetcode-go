package mylinkedlist

import "leetcode-go/template"

// 707. 设计链表
// https://leetcode.cn/problems/design-linked-list/
// 实现 MyLinkedList：get(index)、addAtHead(val)、addAtTail(val)、
// addAtIndex(index, val)、deleteAtIndex(index)，下标从 0 开始。
//   - addAtIndex：index 等于长度时追加到末尾；index 大于长度时不插入；
//   - get / deleteAtIndex：下标无效时分别返回 -1 / 什么也不做。
//
// 核心思路：单链表 + 哑节点 + 记录长度 size。
// 有了哑节点，「在第 index 个位置插入/删除」统一变成：先走到第 index-1 个节点（即 dummy 走 index 步），
// 再改它的 Next；头部操作就是 index=0，尾部操作就是 index=size。

type ListNode = template.ListNode

// MyLinkedList 解法一：哑节点单链表 + size ⭐ 最优方案
// 所有操作都复用「找到第 index 个位置的前驱」这一步。
// 时间复杂度：AddAtHead O(1)，其余 O(index)（AddAtTail O(n)）；空间复杂度 O(n)。
type MyLinkedList struct {
	dummy *ListNode
	size  int
}

func Constructor() MyLinkedList {
	return MyLinkedList{dummy: &ListNode{}}
}

// prevOf 返回第 index 个节点的前驱（要求 0 <= index <= size）。
func (this *MyLinkedList) prevOf(index int) *ListNode {
	prev := this.dummy
	for i := 0; i < index; i++ {
		prev = prev.Next
	}
	return prev
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	return this.prevOf(index).Next.Val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index, val int) {
	if index > this.size {
		return
	}
	if index < 0 { // 兼容旧版题意：负数下标视为插到头部
		index = 0
	}
	prev := this.prevOf(index)
	prev.Next = &ListNode{Val: val, Next: prev.Next}
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	prev := this.prevOf(index)
	prev.Next = prev.Next.Next
	this.size--
}
