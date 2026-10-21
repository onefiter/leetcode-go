package splitlist

import "leetcode-go/template"

// 725. 分隔链表
// https://leetcode.cn/problems/split-linked-list-in-parts/
// 给你链表头 head 和整数 k，把链表分成 k 个连续的部分：各部分长度尽可能相等，
// 任意两部分长度差不超过 1，且排在前面的部分长度不小于后面的。不足的部分用 nil 表示。
//
// 核心思路：先数出长度 n，则每部分基础长度 base = n / k，前 extra = n % k 个部分各多 1 个。
// 然后从头走，每段走够长度就把尾巴断开，下一段从断点继续。

type ListNode = template.ListNode

// splitListToParts 解法一：计算长度后原地切分 ⭐ 最优方案
// 第 i 段长度为 base + (i < extra ? 1 : 0)；走到段尾时记下下一段头并断开 Next。
// 时间复杂度 O(n + k)，空间复杂度 O(1)（不计返回的切片）。
func splitListToParts(head *ListNode, k int) []*ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	base, extra := n/k, n%k
	parts := make([]*ListNode, k)
	cur := head
	for i := 0; i < k && cur != nil; i++ {
		parts[i] = cur
		size := base
		if i < extra {
			size++
		}
		for j := 1; j < size; j++ { // 走到本段最后一个节点
			cur = cur.Next
		}
		next := cur.Next
		cur.Next = nil
		cur = next
	}
	return parts
}

// splitListToPartsEnumeration 解法二：转数组后按长度重建每段（对照用）
// 把值存进切片，按同样的长度规则切片后分别新建链表。
// 时间复杂度 O(n + k)，空间复杂度 O(n)。
func splitListToPartsEnumeration(head *ListNode, k int) []*ListNode {
	vals := template.ListToSlice(head)
	n := len(vals)
	parts := make([]*ListNode, k)
	start := 0
	for i := 0; i < k; i++ {
		size := n / k
		if i < n%k {
			size++
		}
		parts[i] = template.BuildList(vals[start : start+size])
		start += size
	}
	return parts
}
