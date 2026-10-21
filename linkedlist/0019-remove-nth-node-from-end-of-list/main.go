package removenthfromend

import "leetcode-go/template"

// 19. 删除链表的倒数第 N 个结点
// https://leetcode.cn/problems/remove-nth-node-from-end-of-list/
// 给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
// 例如 head=[1,2,3,4,5], n=2 -> [1,2,3,5]。
//
// 一句话理解题意：要删掉一个节点，必须先找到它的「前一个节点」。
//
// 核心思路：用哑节点 dummy 统一处理「删除头节点」的情况；
// 让快指针先走 n 步，再让快慢指针一起走，快指针到达末尾时，
// 慢指针正好停在倒数第 n 个节点的前一个节点上。

type ListNode = template.ListNode

// removeNthFromEnd 解法一：快慢指针（一次遍历）⭐ 最优方案
// fast 从 dummy 先走 n 步，然后 fast、slow 同步前进直到 fast.Next 为 nil，
// 此时 slow 与 fast 相距 n，slow.Next 就是倒数第 n 个节点。
// 时间复杂度 O(L)，空间复杂度 O(1)。
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	fast, slow := dummy, dummy
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next // 跳过倒数第 n 个节点
	return dummy.Next
}

// removeNthFromEndEnumeration 解法二：两次遍历（对照用）
// 第一次遍历求链表长度 L，倒数第 n 个就是正数第 L-n+1 个；
// 第二次从 dummy 走 L-n 步到达它的前一个节点，然后删除。
// 时间复杂度 O(L)，空间复杂度 O(1)。
func removeNthFromEndEnumeration(head *ListNode, n int) *ListNode {
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}
	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 0; i < length-n; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	return dummy.Next
}
