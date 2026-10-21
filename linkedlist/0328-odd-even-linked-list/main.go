package oddevenlist

import "leetcode-go/template"

// 328. 奇偶链表
// https://leetcode.cn/problems/odd-even-linked-list/
// 给定单链表头节点 head，把所有「索引为奇数」的节点放在一起，再接上所有「索引为偶数」的节点
// （第一个节点索引为 1）。要求 O(1) 额外空间、O(n) 时间，且两组内部保持原有相对顺序。
//
// 核心思路：一边遍历一边把链表拆成「奇链」和「偶链」两条，最后把偶链接到奇链尾部。
// odd 和 even 交替前进：odd.Next = even.Next，even.Next = odd.Next，像拉拉链一样拆开。

type ListNode = template.ListNode

// oddEvenList 解法一：原地拆分奇偶两条链 ⭐ 最优方案
// evenHead 记住偶链的头；循环条件 even != nil && even.Next != nil 保证下一个奇节点存在。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	odd, even := head, head.Next
	evenHead := even
	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}
	odd.Next = evenHead
	return head
}

// oddEvenListEnumeration 解法二：按位置收集值再重建（对照用）
// 先收集奇数位的值，再收集偶数位的值，拼接后新建一条链表。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func oddEvenListEnumeration(head *ListNode) *ListNode {
	odds, evens := []int{}, []int{}
	i := 1
	for cur := head; cur != nil; cur = cur.Next {
		if i%2 == 1 {
			odds = append(odds, cur.Val)
		} else {
			evens = append(evens, cur.Val)
		}
		i++
	}
	return template.BuildList(append(odds, evens...))
}
