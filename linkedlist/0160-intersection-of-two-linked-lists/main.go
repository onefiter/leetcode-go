package intersection

import "leetcode-go/template"

// 160. 相交链表
// https://leetcode.cn/problems/intersection-of-two-linked-lists/
// 给两个单链表的头节点 headA 和 headB，找出并返回两个链表相交的起始节点；不相交返回 nil。
// 注意「相交」指的是同一个节点（指针相同），而不是值相同。
//
// 核心思路：A 的长度 = a + c，B 的长度 = b + c（c 是公共尾部）。
// 让指针 pa 走完 A 再走 B，pb 走完 B 再走 A，两者都走 a + b + c 步后
// 恰好同时到达相交点；若不相交（c = 0），会同时走到 nil。

type ListNode = template.ListNode

// getIntersectionNode 解法一：双指针「走完自己走对方」 ⭐ 最优方案
// pa、pb 分别从 headA、headB 出发，走到末尾就跳到另一条链表的头，
// 相遇点即为交点（不相交时在 nil 处相遇）。
// 时间复杂度 O(m+n)，空间复杂度 O(1)。
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}

// getIntersectionNodeHash 解法二：哈希集合
// 先把 A 的所有节点放进集合，再遍历 B，第一个出现在集合里的节点就是交点。
// 时间复杂度 O(m+n)，空间复杂度 O(m)。
func getIntersectionNodeHash(headA, headB *ListNode) *ListNode {
	seen := map[*ListNode]bool{}
	for cur := headA; cur != nil; cur = cur.Next {
		seen[cur] = true
	}
	for cur := headB; cur != nil; cur = cur.Next {
		if seen[cur] {
			return cur
		}
	}
	return nil
}

// getIntersectionNodeEnumeration 解法三：双重循环暴力（对照用）
// 对 A 的每个节点，遍历整条 B 看是否有同一个节点。
// 时间复杂度 O(m*n)，空间复杂度 O(1)。
func getIntersectionNodeEnumeration(headA, headB *ListNode) *ListNode {
	for a := headA; a != nil; a = a.Next {
		for b := headB; b != nil; b = b.Next {
			if a == b {
				return a
			}
		}
	}
	return nil
}
