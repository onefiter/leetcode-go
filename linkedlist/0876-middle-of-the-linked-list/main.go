package middlenode

import "leetcode-go/template"

// 876. 链表的中间结点
// https://leetcode.cn/problems/middle-of-the-linked-list/
// 给你单链表的头结点 head，请你找出并返回链表的中间结点。
// 如果有两个中间结点，则返回第二个中间结点。
// 例如 head=[1,2,3,4,5] -> 从 3 开始；head=[1,2,3,4,5,6] -> 从 4 开始。
//
// 一句话理解题意：找链表正中间那个节点，偶数长度取靠后的那个。
//
// 核心思路：快指针每次走两步、慢指针每次走一步，
// 快指针走到末尾时，慢指针恰好走了一半。

type ListNode = template.ListNode

// middleNode 解法一：快慢指针 ⭐ 最优方案
// 循环条件 fast != nil && fast.Next != nil 保证偶数长度时 slow 停在第二个中间节点。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// middleNodeEnumeration 解法二：两次遍历（对照用）
// 第一次遍历求长度 n，第二次从头走 n/2 步即为（第二个）中间节点。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func middleNodeEnumeration(head *ListNode) *ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	cur := head
	for i := 0; i < n/2; i++ {
		cur = cur.Next
	}
	return cur
}
