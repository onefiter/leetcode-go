package detectcycle

import "leetcode-go/template"

// 142. 环形链表 II
// https://leetcode.cn/problems/linked-list-cycle-ii/
// 给定一个链表的头节点 head，返回链表开始入环的第一个节点；无环则返回 nil。
// 例如 head=[3,2,0,-4]，尾节点连到下标 1 -> 返回值为 2 的那个节点。
//
// 一句话理解题意：不只是判断有没有环，还要找到「从哪个节点进入环」。
//
// 核心思路（Floyd）：设头到入口距离 a，入口到相遇点距离 b，环长 c。
// 相遇时 slow 走了 a+b，fast 走了 2(a+b)，且 fast 多走了整数圈：a+b = k*c，
// 所以 a = k*c - b，即「从相遇点再走 a 步」恰好回到入口。
// 于是相遇后让一个指针回到 head，两者每次各走一步，再次相遇处就是入口。

type ListNode = template.ListNode

// detectCycle 解法一：Floyd 快慢指针找入口 ⭐ 最优方案
// 第一阶段：快慢指针判断是否有环并找到相遇点；
// 第二阶段：ptr 从 head 出发、slow 从相遇点出发，同速前进，相遇点即入口。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}
	return nil
}

// detectCycleEnumeration 解法二：哈希集合（对照用）
// 沿链表遍历并记录访问过的节点，第一个被重复访问的节点就是入环点。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func detectCycleEnumeration(head *ListNode) *ListNode {
	seen := map[*ListNode]struct{}{}
	for cur := head; cur != nil; cur = cur.Next {
		if _, ok := seen[cur]; ok {
			return cur
		}
		seen[cur] = struct{}{}
	}
	return nil
}
