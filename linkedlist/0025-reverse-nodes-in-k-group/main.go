package reversekgroup

import "leetcode-go/template"

// 25. K 个一组翻转链表
// https://leetcode.cn/problems/reverse-nodes-in-k-group/
// 每 k 个节点一组进行翻转，最后不足 k 个的节点保持原顺序。只能改节点指针，不能改值。
// 例如 [1,2,3,4,5], k=2 -> [2,1,4,3,5]；k=3 -> [3,2,1,4,5]。
//
// 核心思路：用 dummy 哨兵，groupPrev 指向「当前组之前的节点」。
// 先从 groupPrev 往后数 k 个，不够 k 个就结束；够的话把这一段原地反转，
// 再把反转后的段接回原链表，groupPrev 移到这段的新尾巴（即原来的组头）。

type ListNode = template.ListNode

// reverseKGroup 解法一：迭代分组原地反转 ⭐ 最优方案
// 每组：找到组尾 tail，记下 next = tail.Next；把 [groupHead, tail] 反转，
// 然后 groupPrev.Next = tail，groupPrev = groupHead。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	groupPrev := dummy
	for {
		// 找本组的第 k 个节点，不足 k 个直接结束。
		tail := groupPrev
		for i := 0; i < k && tail != nil; i++ {
			tail = tail.Next
		}
		if tail == nil {
			break
		}
		groupHead, next := groupPrev.Next, tail.Next
		// 反转 [groupHead, tail]：prev 初始为 next，反转后 groupHead 自然接到 next。
		prev, cur := next, groupHead
		for cur != next {
			cur.Next, prev, cur = prev, cur, cur.Next
		}
		groupPrev.Next = tail
		groupPrev = groupHead
	}
	return dummy.Next
}

// reverseKGroupRecursive 解法二：递归
// 先检查前 k 个是否存在；存在就反转这 k 个，原头节点接上「剩余部分递归处理的结果」。
// 时间复杂度 O(n)，空间复杂度 O(n/k)（递归栈）。
func reverseKGroupRecursive(head *ListNode, k int) *ListNode {
	cur := head
	for i := 0; i < k; i++ {
		if cur == nil {
			return head // 不足 k 个，保持原样
		}
		cur = cur.Next
	}
	// 此时 cur 是下一组的头；反转 [head, cur)。
	var prev *ListNode
	node := head
	for node != cur {
		node.Next, prev, node = prev, node, node.Next
	}
	head.Next = reverseKGroupRecursive(cur, k)
	return prev
}

// reverseKGroupEnumeration 解法三：节点存切片后分段反转（对照用）
// 把节点放进切片，每满 k 个就在切片里反转这一段，最后按切片顺序重新串起来。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func reverseKGroupEnumeration(head *ListNode, k int) *ListNode {
	nodes := []*ListNode{}
	for cur := head; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
	}
	for i := 0; i+k <= len(nodes); i += k {
		for l, r := i, i+k-1; l < r; l, r = l+1, r-1 {
			nodes[l], nodes[r] = nodes[r], nodes[l]
		}
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
