package deleteduplicates

import "leetcode-go/template"

// 83. 删除排序链表中的重复元素
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list/
// 给一个已排序链表，删除重复元素，使每个元素只出现一次。
// 例如 [1,1,2,3,3] -> [1,2,3]。
//
// 核心思路：有序意味着相同的值一定挨在一起，
// 只要当前节点和下一个节点值相同，就把下一个节点跳过；否则往后走。

type ListNode = template.ListNode

// deleteDuplicates 解法一：一次遍历 ⭐ 最优方案
// 表头不会被删（保留每个值的第一个），所以不需要哨兵。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

// deleteDuplicatesRecursive 解法二：递归
// 先处理好 head.Next 之后的部分，再看 head 是否与新的下一个节点重复。
// 时间复杂度 O(n)，空间复杂度 O(n)（递归栈）。
func deleteDuplicatesRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head.Next = deleteDuplicatesRecursive(head.Next)
	if head.Val == head.Next.Val {
		return head.Next
	}
	return head
}

// deleteDuplicatesEnumeration 解法三：哈希集合去重（对照用）
// 用集合记录出现过的值，遇到已出现的值就跳过，没有利用「有序」条件。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func deleteDuplicatesEnumeration(head *ListNode) *ListNode {
	seen := map[int]bool{}
	dummy := &ListNode{}
	tail := dummy
	for cur := head; cur != nil; cur = cur.Next {
		if !seen[cur.Val] {
			seen[cur.Val] = true
			tail.Next = cur
			tail = cur
		}
	}
	tail.Next = nil
	return dummy.Next
}
