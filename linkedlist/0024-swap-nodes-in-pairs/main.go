package swappairs

import "leetcode-go/template"

// 24. 两两交换链表中的节点
// https://leetcode.cn/problems/swap-nodes-in-pairs/
// 两两交换相邻节点并返回新表头，只能交换节点本身，不能只改节点值。
// 例如 [1,2,3,4] -> [2,1,4,3]；奇数个时最后一个不动：[1,2,3] -> [2,1,3]。
//
// 核心思路：用哨兵 dummy 指向表头，prev 指向「待交换的一对」之前的节点，
// 每次把 prev -> a -> b -> next 改成 prev -> b -> a -> next，然后 prev 前进到 a。

type ListNode = template.ListNode

// swapPairs 解法一：迭代 + 哨兵节点 ⭐ 最优方案
// 每轮只改三根指针：prev.Next = b，a.Next = b.Next，b.Next = a。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	for prev.Next != nil && prev.Next.Next != nil {
		a, b := prev.Next, prev.Next.Next
		prev.Next = b
		a.Next = b.Next
		b.Next = a
		prev = a
	}
	return dummy.Next
}

// swapPairsRecursive 解法二：递归
// 交换前两个节点，第二个节点变成新头，原头节点接上「剩余部分递归交换后的结果」。
// 时间复杂度 O(n)，空间复杂度 O(n)（递归栈）。
func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	head.Next = swapPairsRecursive(newHead.Next)
	newHead.Next = head
	return newHead
}

// swapPairsEnumeration 解法三：存进切片后重新连接（对照用）
// 把所有节点放进切片，在切片里两两交换位置，再按新顺序把节点串起来。
// 交换的仍然是节点本身，满足题意，但需要 O(n) 额外空间。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func swapPairsEnumeration(head *ListNode) *ListNode {
	nodes := []*ListNode{}
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}
	for i := 0; i+1 < len(nodes); i += 2 {
		nodes[i], nodes[i+1] = nodes[i+1], nodes[i]
	}
	dummy := &ListNode{}
	cur := dummy
	for _, node := range nodes {
		cur.Next = node
		cur = node
	}
	cur.Next = nil
	return dummy.Next
}
