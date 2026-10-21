package removenodes

import "leetcode-go/template"

// 2487. 从链表中移除节点
// https://leetcode.cn/problems/remove-nodes-from-linked-list/
// 如果某个节点右侧存在一个严格更大的值，就删除这个节点，返回修改后的链表头。
// 例如 [5,2,13,3,8] -> [13,8]。
//
// 核心思路：保留下来的节点从左到右是「非递增」的——
// 每个保留节点都 >= 它右边所有节点。
// 可以用单调栈从左往右维护；也可以反转链表后从左往右扫，记录当前最大值；
// 或者用递归：先处理好后面的链表，再决定当前节点要不要保留。

type ListNode = template.ListNode

// removeNodes 解法一：单调栈
// 从左往右遍历，栈里维护一个非递增序列；遇到更大的值就把栈顶比它小的全部弹出
// （它们右边出现了更大的值，必须删除）。最后把栈里的节点依次串起来。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func removeNodes(head *ListNode) *ListNode {
	var st []*ListNode
	for p := head; p != nil; p = p.Next {
		for len(st) > 0 && st[len(st)-1].Val < p.Val {
			st = st[:len(st)-1]
		}
		st = append(st, p)
	}
	for i := 0; i+1 < len(st); i++ {
		st[i].Next = st[i+1]
	}
	st[len(st)-1].Next = nil
	return st[0]
}

// removeNodesReverse 解法二：反转 + 一次遍历 + 再反转 ⭐ 最优方案（空间 O(1)）
// 反转后问题变成「删除左侧存在更大值的节点」：从头扫，cur 始终是已保留节点中的最大值，
// cur.Next 比它小就跳过。最后再反转回来。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func removeNodesReverse(head *ListNode) *ListNode {
	head = reverse(head)
	cur := head
	for cur.Next != nil {
		if cur.Next.Val < cur.Val {
			cur.Next = cur.Next.Next // 左边有更大的，删除
		} else {
			cur = cur.Next
		}
	}
	return reverse(head)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

// removeNodesRecursive 解法三：递归
// 先递归处理 head.Next，得到的新头就是右侧的最大值；
// 若它比 head 大，则 head 被删，直接返回它；否则把它接在 head 后面。
// 时间复杂度 O(n)，空间复杂度 O(n)（递归栈）。
func removeNodesRecursive(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	next := removeNodesRecursive(head.Next)
	if next.Val > head.Val {
		return next
	}
	head.Next = next
	return head
}

// removeNodesEnumeration 解法四：暴力（对照用）
// 对每个节点向右扫描一遍，看是否存在严格更大的值，不存在才保留。
// 时间复杂度 O(n²)，空间复杂度 O(1)。
func removeNodesEnumeration(head *ListNode) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for p := head; p != nil; {
		next := p.Next
		keep := true
		for q := next; q != nil; q = q.Next {
			if q.Val > p.Val {
				keep = false
				break
			}
		}
		if keep {
			tail.Next = p
			tail = p
		}
		p = next
	}
	tail.Next = nil
	return dummy.Next
}
