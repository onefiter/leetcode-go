package reverselist

import "leetcode-go/template"

// 206. 反转链表
// https://leetcode.cn/problems/reverse-linked-list/
// 给你单链表的头节点 head，请你反转链表，并返回反转后的链表头。
//
// 核心思路：把每个节点的 Next 指针掉个头，指向它原来的前一个节点。
// 迭代时用 prev 记录「已反转部分的头」，逐个把 cur 摘下来接到 prev 前面；
// 递归时先反转后面的部分，再把当前节点挂到后半段的尾巴上。

type ListNode = template.ListNode

// reverseList 解法一：迭代（双指针） ⭐ 最优方案
// 每一步：先记住 cur.Next，再让 cur.Next 指向 prev，然后 prev、cur 各前进一步。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}

// reverseListRecursive 解法二：递归
// reverseListRecursive(head.Next) 返回反转后的新头，此时 head.Next 是后半段的尾巴，
// 让 head.Next.Next = head 把 head 接上，再断开 head.Next 防止成环。
// 时间复杂度 O(n)，空间复杂度 O(n)（递归栈）。
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return newHead
}

// reverseListEnumeration 解法三：存到切片再倒序写回（对照用）
// 先把所有值存下来，再从头到尾把值按倒序写回各节点。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func reverseListEnumeration(head *ListNode) *ListNode {
	vals := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	i := len(vals) - 1
	for cur := head; cur != nil; cur = cur.Next {
		cur.Val = vals[i]
		i--
	}
	return head
}
