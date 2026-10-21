package reversebetween

import "leetcode-go/template"

// 92. 反转链表 II
// https://leetcode.cn/problems/reverse-linked-list-ii/
// 反转链表中从第 left 个到第 right 个（1 开始计数）的节点，返回反转后的链表。
// 例如 [1,2,3,4,5], left=2, right=4 -> [1,4,3,2,5]。
//
// 核心思路：用 dummy 哨兵（left 可能是 1），先走到第 left-1 个节点 prev。
// 然后「头插法」：把 prev 后面那段的第 2、3、…个节点依次摘下来插到 prev 后面，
// 共做 right-left 次，这一段就反转好了，且只遍历一遍。

type ListNode = template.ListNode

// reverseBetween 解法一：一次遍历 + 头插法 ⭐ 最优方案
// cur 始终是这段的原第一个节点（反转后会成为段尾），每次把 cur.Next 移到 prev 后面。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 0; i < left-1; i++ {
		prev = prev.Next
	}
	cur := prev.Next
	for i := 0; i < right-left; i++ {
		next := cur.Next
		cur.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dummy.Next
}

// reverseBetweenEnumeration 解法二：收集值、反转区间、写回（对照用）
// 把所有值读进切片，反转 [left-1, right-1] 区间后按顺序写回原节点。
// 简单但改的是节点值，并且需要 O(n) 额外空间。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func reverseBetweenEnumeration(head *ListNode, left int, right int) *ListNode {
	vals := template.ListToSlice(head)
	for l, r := left-1, right-1; l < r; l, r = l+1, r-1 {
		vals[l], vals[r] = vals[r], vals[l]
	}
	i := 0
	for cur := head; cur != nil; cur = cur.Next {
		cur.Val = vals[i]
		i++
	}
	return head
}
