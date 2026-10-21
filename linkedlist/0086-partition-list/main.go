package partitionlist

import "leetcode-go/template"

// 86. 分隔链表
// https://leetcode.cn/problems/partition-list/
// 给链表和值 x，把所有小于 x 的节点放到大于或等于 x 的节点之前，
// 两部分内部都要保持原来的相对顺序。例如 [1,4,3,2,5,2], x=3 -> [1,2,2,4,3,5]。
//
// 核心思路：准备两条新链表 small 和 large（各带一个哨兵），
// 遍历原链表，按 Val < x 分别接到两条链表尾部，最后 small 的尾巴接上 large 的头。

type ListNode = template.ListNode

// partition 解法一：双哨兵拆分再拼接 ⭐ 最优方案
// 注意 large 尾部要置 nil，否则可能指向原链表中排在它后面的 small 节点，形成环。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func partition(head *ListNode, x int) *ListNode {
	small, large := &ListNode{}, &ListNode{}
	s, l := small, large
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			s.Next = cur
			s = cur
		} else {
			l.Next = cur
			l = cur
		}
	}
	l.Next = nil
	s.Next = large.Next
	return small.Next
}

// partitionEnumeration 解法二：收集值后重建链表（对照用）
// 先把 < x 的值和 >= x 的值分别收集到两个切片，再拼成新链表。
// 思路直白，但创建了新节点且用了 O(n) 额外空间。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func partitionEnumeration(head *ListNode, x int) *ListNode {
	less, more := []int{}, []int{}
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val < x {
			less = append(less, cur.Val)
		} else {
			more = append(more, cur.Val)
		}
	}
	return template.BuildList(append(less, more...))
}
