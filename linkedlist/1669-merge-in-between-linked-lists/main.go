package mergeinbetween

import "leetcode-go/template"

// 1669. 合并两个链表
// https://leetcode.cn/problems/merge-in-between-linked-lists/
// 给两个链表 list1 和 list2，把 list1 中下标从 a 到 b 的节点删掉，
// 并把 list2 接到被删除的位置。题目保证 1 <= a <= b < len(list1) - 1。
//
// 核心思路：找到「第 a 个节点的前一个节点」prev 和「第 b 个节点的后一个节点」next，
// 然后 prev.Next = list2，list2 的尾巴.Next = next，就完成了拼接。

type ListNode = template.ListNode

// mergeInBetween 解法一：找两个断点直接拼接 ⭐ 最优方案
// 从头走 a-1 步得到 prev，再从 prev 继续走 b-a+2 步得到 next；
// 遍历 list2 找到尾节点，把三段接起来。
// 时间复杂度 O(n + m)，空间复杂度 O(1)。
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	prev := list1
	for i := 0; i < a-1; i++ {
		prev = prev.Next
	}
	next := prev
	for i := 0; i < b-a+2; i++ {
		next = next.Next
	}
	prev.Next = list2
	tail := list2
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = next
	return list1
}

// mergeInBetweenEnumeration 解法二：转成数组再重建（对照用）
// 把 list1 的值拆成 [0,a) 和 (b,n) 两段，中间塞入 list2 的值，最后重新建链表。
// 时间复杂度 O(n + m)，空间复杂度 O(n + m)。
func mergeInBetweenEnumeration(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var v1, v2 []int
	for p := list1; p != nil; p = p.Next {
		v1 = append(v1, p.Val)
	}
	for p := list2; p != nil; p = p.Next {
		v2 = append(v2, p.Val)
	}
	vals := make([]int, 0, len(v1)-(b-a+1)+len(v2))
	vals = append(vals, v1[:a]...)
	vals = append(vals, v2...)
	vals = append(vals, v1[b+1:]...)
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}
