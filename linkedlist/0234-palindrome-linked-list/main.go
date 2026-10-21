package palindromelist

import "leetcode-go/template"

// 234. 回文链表
// https://leetcode.cn/problems/palindrome-linked-list/
// 给你一个单链表的头节点 head，判断该链表是否为回文链表。
// 例如 head=[1,2,2,1] -> true；head=[1,2] -> false。
//
// 一句话理解题意：从前往后读和从后往前读一样。
//
// 核心思路：单链表没法倒着走，那就把后半段反转，
// 然后前半段和反转后的后半段同时从头比较；比较完再把后半段反转回去，恢复原链表。

type ListNode = template.ListNode

// isPalindrome 解法一：快慢指针找中点 + 反转后半段 ⭐ 最优方案
// 1. 快慢指针找到前半段末尾 slow（奇数长度时中间节点归前半段）；
// 2. 反转 slow.Next 开始的后半段；
// 3. 两个指针分别从 head 和后半段头开始逐个比较，后半段走完即结束；
// 4. 再次反转后半段并接回，保持输入不被破坏。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondHead := reverse(slow.Next)
	ok := true
	for p1, p2 := head, secondHead; p2 != nil; p1, p2 = p1.Next, p2.Next {
		if p1.Val != p2.Val {
			ok = false
			break
		}
	}
	slow.Next = reverse(secondHead) // 恢复原链表
	return ok
}

// reverse 原地反转链表并返回新头节点。
func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}

// isPalindromeEnumeration 解法二：转数组 + 双指针（对照用）
// 把链表值复制到切片，再用首尾双指针向中间比较。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func isPalindromeEnumeration(head *ListNode) bool {
	vals := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		if vals[i] != vals[j] {
			return false
		}
	}
	return true
}
