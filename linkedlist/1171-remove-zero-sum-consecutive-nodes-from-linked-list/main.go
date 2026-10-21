package removezerosum

import "leetcode-go/template"

// 1171. 从链表中删去总和值为零的连续节点
// https://leetcode.cn/problems/remove-zero-sum-consecutive-nodes-from-linked-list/
// 反复删去链表中总和为 0 的连续节点组成的序列，直到不存在这样的序列，返回头节点。
// 例如 [1,2,-3,3,1] -> [3,1]（[1,2,1] 也是正确答案）。
//
// 核心思路：前缀和。如果两个位置的前缀和相同，说明它们之间那段的和为 0，可以整段删掉。
// 第一遍：记录每个前缀和「最后一次」出现的节点（后出现的覆盖前面的）；
// 第二遍：对每个节点，直接把 Next 指向「同一前缀和最后出现节点」的 Next，
// 中间所有和为 0 的片段一次性被跳过。

type ListNode = template.ListNode

// removeZeroSumSublists 解法一：前缀和 + 哈希表，两次遍历 ⭐ 最优方案
// 用哑节点代表前缀和 0，这样从头开始的零和段也能被删除。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func removeZeroSumSublists(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	last := map[int]*ListNode{}
	sum := 0
	for p := dummy; p != nil; p = p.Next {
		sum += p.Val
		last[sum] = p
	}
	sum = 0
	for p := dummy; p != nil; p = p.Next {
		sum += p.Val
		p.Next = last[sum].Next
	}
	return dummy.Next
}

// removeZeroSumSublistsEnumeration 解法二：暴力反复删除（对照用）
// 枚举起点，从起点往后累加，一旦和为 0 就把这一段删掉，然后从头重新开始，
// 直到一整轮都找不到零和段为止。
// 时间复杂度 O(n³)，空间复杂度 O(1)。
func removeZeroSumSublistsEnumeration(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	for changed := true; changed; {
		changed = false
	scan:
		for prev := dummy; prev.Next != nil; prev = prev.Next {
			sum := 0
			for q := prev.Next; q != nil; q = q.Next {
				sum += q.Val
				if sum == 0 {
					prev.Next = q.Next // 删除 prev.Next..q
					changed = true
					break scan // 直接跳出，避免 prev 前进到已删除区段之后的 nil
				}
			}
		}
	}
	return dummy.Next
}
