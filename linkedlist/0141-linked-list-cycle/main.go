package hascycle

import "leetcode-go/template"

// 141. 环形链表
// https://leetcode.cn/problems/linked-list-cycle/
// 给你一个链表的头节点 head，判断链表中是否有环。
// 例如 head=[3,2,0,-4]，尾节点连到下标 1 -> true。
//
// 一句话理解题意：沿着 Next 一直走，如果永远走不到 nil，就说明有环。
//
// 核心思路：龟兔赛跑。慢指针一次走一步，快指针一次走两步；
// 无环时快指针会先走到 nil；有环时两者都会进入环，
// 快指针每轮比慢指针多走一步，相当于在环上「追」慢指针，最终一定相遇。

type ListNode = template.ListNode

// hasCycle 解法一：快慢指针（Floyd 判圈）⭐ 最优方案
// fast 每次走两步、slow 每次走一步，相遇即有环，fast 走到 nil 即无环。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// hasCycleEnumeration 解法二：哈希集合（对照用）
// 遍历时把访问过的节点指针放进集合，再次遇到同一个节点说明有环。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func hasCycleEnumeration(head *ListNode) bool {
	seen := map[*ListNode]struct{}{}
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := seen[cur]; ok {
			return true
		}
		seen[cur] = struct{}{}
	}
	return false
}
