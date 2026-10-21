package sortlist

import (
	"leetcode-go/template"
	"sort"
)

// 148. 排序链表
// https://leetcode.cn/problems/sort-list/
// 给你链表的头节点 head，请将其按升序排列并返回排序后的链表。
// 进阶：在 O(n log n) 时间、O(1) 额外空间下完成。
//
// 核心思路：链表天然适合归并排序——合并两个有序链表只需改指针，不需要额外数组。
//   - 自顶向下：快慢指针找中点切两半，递归排序后合并（递归栈 O(log n)）；
//   - 自底向上：先把长度为 1 的段两两合并，再合并长度为 2 的段……直到覆盖整条链表，
//     全程迭代，真正做到 O(1) 额外空间。

type ListNode = template.ListNode

// sortList 解法一：自底向上归并 ⭐ 最优方案
// step 从 1 开始每轮翻倍：每轮从头扫描，切下两段长度为 step 的子链表合并，
// 再接到已处理部分的尾巴上。step >= n 时整条链表有序。
// 时间复杂度 O(n log n)，空间复杂度 O(1)。
func sortList(head *ListNode) *ListNode {
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	dummy := &ListNode{Next: head}
	for step := 1; step < n; step <<= 1 {
		prev, cur := dummy, dummy.Next
		for cur != nil {
			left := cur
			right := split(left, step) // 切出第一段，right 是第二段开头
			cur = split(right, step)   // 切出第二段，cur 是剩余部分开头
			prev.Next = merge(left, right)
			for prev.Next != nil { // prev 走到合并结果的尾部
				prev = prev.Next
			}
		}
	}
	return dummy.Next
}

// split 从 head 起保留 step 个节点并断开，返回后半部分的头（可能为 nil）。
func split(head *ListNode, step int) *ListNode {
	for i := 1; head != nil && i < step; i++ {
		head = head.Next
	}
	if head == nil {
		return nil
	}
	rest := head.Next
	head.Next = nil
	return rest
}

// merge 合并两个有序链表。
func merge(a, b *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for a != nil && b != nil {
		if a.Val <= b.Val {
			tail.Next, a = a, a.Next
		} else {
			tail.Next, b = b, b.Next
		}
		tail = tail.Next
	}
	if a != nil {
		tail.Next = a
	} else {
		tail.Next = b
	}
	return dummy.Next
}

// sortListTopDown 解法二：自顶向下归并
// 快慢指针找中点（slow 停在前半段末尾），断开后分别递归排序再合并。
// 时间复杂度 O(n log n)，空间复杂度 O(log n)（递归栈）。
func sortListTopDown(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	return merge(sortListTopDown(head), sortListTopDown(mid))
}

// sortListEnumeration 解法三：转数组排序（对照用）
// 把值拷到切片里排序，再按顺序写回原节点。
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func sortListEnumeration(head *ListNode) *ListNode {
	vals := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	sort.Ints(vals)
	i := 0
	for cur := head; cur != nil; cur = cur.Next {
		cur.Val = vals[i]
		i++
	}
	return head
}
