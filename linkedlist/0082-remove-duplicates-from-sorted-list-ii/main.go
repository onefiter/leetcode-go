package deleteduplicatesii

import "leetcode-go/template"

// 82. 删除排序链表中的重复元素 II
// https://leetcode.cn/problems/remove-duplicates-from-sorted-list-ii/
// 给一个已排序链表，删除所有「出现过重复」的数字，只保留原链表中只出现一次的数字。
// 例如 [1,2,3,3,4,4,5] -> [1,2,5]；[1,1,1,2,3] -> [2,3]。
//
// 核心思路：表头也可能被删，所以用 dummy 哨兵。
// prev 指向「已确认保留」的最后一个节点，看 prev 后面两个节点值是否相同：
// 相同就记下这个值，把所有等于它的节点全部跳过；不同就说明 prev.Next 是唯一的，prev 前进。

type ListNode = template.ListNode

// deleteDuplicates 解法一：哨兵 + 一次遍历 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func deleteDuplicates(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		if prev.Next.Val == prev.Next.Next.Val {
			dup := prev.Next.Val
			for prev.Next != nil && prev.Next.Val == dup {
				prev.Next = prev.Next.Next
			}
		} else {
			prev = prev.Next
		}
	}
	return dummy.Next
}

// deleteDuplicatesRecursive 解法二：递归
// 若头节点与下一个值不同，头保留，递归处理后面；
// 若相同，跳过所有等于该值的节点，对剩余部分递归（头不保留）。
// 时间复杂度 O(n)，空间复杂度 O(n)（递归栈）。
func deleteDuplicatesRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Val != head.Next.Val {
		head.Next = deleteDuplicatesRecursive(head.Next)
		return head
	}
	cur := head
	for cur != nil && cur.Val == head.Val {
		cur = cur.Next
	}
	return deleteDuplicatesRecursive(cur)
}

// deleteDuplicatesEnumeration 解法三：哈希计数（对照用）
// 先统计每个值出现次数，再遍历一遍只保留次数为 1 的节点。
// 没有利用「有序」这一条件，对无序链表也适用。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func deleteDuplicatesEnumeration(head *ListNode) *ListNode {
	cnt := map[int]int{}
	for cur := head; cur != nil; cur = cur.Next {
		cnt[cur.Val]++
	}
	dummy := &ListNode{}
	tail := dummy
	for cur := head; cur != nil; cur = cur.Next {
		if cnt[cur.Val] == 1 {
			tail.Next = cur
			tail = cur
		}
	}
	tail.Next = nil
	return dummy.Next
}
