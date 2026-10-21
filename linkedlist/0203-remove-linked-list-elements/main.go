package removeelements

import "leetcode-go/template"

// 203. 移除链表元素
// https://leetcode.cn/problems/remove-linked-list-elements/
// 给你链表头节点 head 和整数 val，删除所有满足 Node.Val == val 的节点，返回新的头节点。
//
// 核心思路：删除节点要改「前一个节点」的 Next，而头节点没有前驱，
// 所以在头前面放一个哑节点 dummy，这样头节点和普通节点处理方式完全一样。

type ListNode = template.ListNode

// removeElements 解法一：哑节点 + 迭代 ⭐ 最优方案
// cur 从 dummy 出发，若 cur.Next 要删就跳过它（cur 不动，继续检查新的 cur.Next），否则 cur 后移。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{Next: head}
	for cur := dummy; cur.Next != nil; {
		if cur.Next.Val == val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return dummy.Next
}

// removeElementsRecursive 解法二：递归
// 先递归处理 head.Next 之后的部分，再决定 head 本身是否保留。
// 时间复杂度 O(n)，空间复杂度 O(n)（递归栈）。
func removeElementsRecursive(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	head.Next = removeElementsRecursive(head.Next, val)
	if head.Val == val {
		return head.Next
	}
	return head
}

// removeElementsEnumeration 解法三：收集保留的值再重建（对照用）
// 遍历一遍把不等于 val 的值放进切片，再用这些值新建一条链表。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func removeElementsEnumeration(head *ListNode, val int) *ListNode {
	keep := []int{}
	for cur := head; cur != nil; cur = cur.Next {
		if cur.Val != val {
			keep = append(keep, cur.Val)
		}
	}
	return template.BuildList(keep)
}
