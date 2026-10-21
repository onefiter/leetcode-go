package deletemiddle

import "leetcode-go/template"

// 2095. 删除链表的中间节点
// https://leetcode.cn/problems/delete-the-middle-node-of-a-linked-list/
// 长度为 n 的链表，中间节点是下标 ⌊n/2⌋（从 0 开始）的节点，删除它并返回头节点。
// 例如 [1,3,4,7,1,2,6] -> [1,3,4,1,2,6]；只有一个节点时返回空链表。
//
// 核心思路：要删除一个节点，需要拿到它的前驱。
// 用快慢指针：慢指针从哑节点出发，快指针从 head 出发，快指针每次走两步，
// 快指针走到尽头时，慢指针恰好停在中间节点的前一个位置。

type ListNode = template.ListNode

// deleteMiddle 解法一：快慢指针 ⭐ 最优方案
// slow 从 dummy 出发，fast 从 head 出发；当 fast 或 fast.Next 为空时停止，
// 此时 slow.Next 就是下标 ⌊n/2⌋ 的节点。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func deleteMiddle(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// deleteMiddleEnumeration 解法二：两次遍历（对照用）
// 第一次遍历求长度 n，第二次走到下标 ⌊n/2⌋ 的前驱处删除。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func deleteMiddleEnumeration(head *ListNode) *ListNode {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	dummy := &ListNode{Next: head}
	prev := dummy
	for i := 0; i < n/2; i++ {
		prev = prev.Next
	}
	prev.Next = prev.Next.Next
	return dummy.Next
}
