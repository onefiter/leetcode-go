package reorderlist

import "leetcode-go/template"

// 143. 重排链表
// https://leetcode.cn/problems/reorder-list/
// 给定单链表 L0→L1→…→Ln-1→Ln，将其重新排列为 L0→Ln→L1→Ln-1→L2→Ln-2→…
// 不能只改节点值，需要实际交换节点。
// 例如 head=[1,2,3,4,5] -> [1,5,2,4,3]。
//
// 一句话理解题意：从两头交替取节点，前一个、后一个、前一个、后一个……
//
// 核心思路：单链表不能从后往前走，那就把后半段「反转」过来，
// 变成两条都从头开始走的链表，再像拉拉链一样交替合并即可。

type ListNode = template.ListNode

// reorderList 解法一：找中点 + 反转后半 + 交替合并 ⭐ 最优方案
// 1. 快慢指针找中点（偶数长度时取前半的最后一个），断开成两段；
// 2. 原地反转后半段；
// 3. 前半段和反转后的后半段交替连接（前半段长度 >= 后半段）。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	// 1. 找中点：slow 停在前半段末尾。
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	second := slow.Next
	slow.Next = nil // 断开
	// 2. 反转后半段。
	var prev *ListNode
	for second != nil {
		next := second.Next
		second.Next = prev
		prev, second = second, next
	}
	// 3. 交替合并。
	first, second := head, prev
	for second != nil {
		n1, n2 := first.Next, second.Next
		first.Next = second
		second.Next = n1
		first, second = n1, n2
	}
}

// reorderListEnumeration 解法二：数组 + 双指针（对照用）
// 把所有节点存进切片，用 i 从头、j 从尾交替连接，最后把末尾节点的 Next 置空。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func reorderListEnumeration(head *ListNode) {
	if head == nil {
		return
	}
	nodes := []*ListNode{}
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}
	i, j := 0, len(nodes)-1
	for i < j {
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}
	nodes[i].Next = nil // i 是新链表的最后一个节点
}
