package mergetwolists

import (
	"leetcode-go/template"
	"slices"
)

// 21. 合并两个有序链表
// https://leetcode.cn/problems/merge-two-sorted-lists/
// 将两个升序链表合并为一个新的升序链表并返回。
// 例如 l1=[1,2,4], l2=[1,3,4] -> [1,1,2,3,4,4]。
//
// 一句话理解题意：像归并排序的「合并」步骤一样，
// 每次比较两个链表的头，把较小的那个接到结果链表的尾巴上。
//
// 核心思路：用哑节点 dummy 作为结果链表的头，tail 指向结果末尾；
// 某个链表先走完后，把另一个链表剩余部分整体接上即可。

type ListNode = template.ListNode

// mergeTwoLists 解法一：迭代 + 哑节点 ⭐ 最优方案
// 每轮取 list1、list2 中较小的头节点接到 tail 后面，并让对应指针后移。
// 相等时优先取 list1，保证稳定。
// 时间复杂度 O(m+n)，空间复杂度 O(1)。
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}
		tail = tail.Next
	}
	// 剩余部分本身有序，直接接上。
	if list1 != nil {
		tail.Next = list1
	} else {
		tail.Next = list2
	}
	return dummy.Next
}

// mergeTwoListsRecursive 解法二：递归
// 较小的头节点作为结果的头，它的 Next 等于「剩余部分的合并结果」。
// 任一链表为空时直接返回另一个。
// 时间复杂度 O(m+n)，空间复杂度 O(m+n)（递归栈）。
func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val <= list2.Val {
		list1.Next = mergeTwoListsRecursive(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoListsRecursive(list1, list2.Next)
	return list2
}

// mergeTwoListsEnumeration 解法三：收集所有值后排序再重建（对照用）
// 不利用「两个链表已有序」的性质，把所有值放进切片排序，再新建链表。
// 时间复杂度 O((m+n) log(m+n))，空间复杂度 O(m+n)。
func mergeTwoListsEnumeration(list1 *ListNode, list2 *ListNode) *ListNode {
	vals := []int{}
	for cur := list1; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	for cur := list2; cur != nil; cur = cur.Next {
		vals = append(vals, cur.Val)
	}
	slices.Sort(vals)
	dummy := &ListNode{}
	tail := dummy
	for _, v := range vals {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return dummy.Next
}
