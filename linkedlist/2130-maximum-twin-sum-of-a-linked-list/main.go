package pairsum

import "leetcode-go/template"

// 2130. 链表最大孪生和
// https://leetcode.cn/problems/maximum-twin-sum-of-a-linked-list/
// 长度为偶数 n 的链表，第 i 个节点和第 n-1-i 个节点互为「孪生节点」，
// 孪生和 = 两者值之和，求最大孪生和。
// 例如 [5,4,2,1]：5+1=6，4+2=6，答案 6。
//
// 核心思路：孪生对就是「前半段从左往右」和「后半段从右往左」一一配对。
// 用快慢指针找到后半段起点，把后半段反转，然后两个指针同时走，逐对求和取最大。

type ListNode = template.ListNode

// pairSum 解法一：快慢指针 + 反转后半段 ⭐ 最优方案
// slow 每次走一步、fast 每次走两步，fast 到尾时 slow 正好是后半段第一个节点；
// 反转后半段后，前半段头和反转后的头同时向后走，逐对相加。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func pairSum(head *ListNode) int {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 反转后半段
	var prev *ListNode
	for slow != nil {
		next := slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}
	ans := 0
	for p, q := head, prev; q != nil; p, q = p.Next, q.Next {
		ans = max(ans, p.Val+q.Val)
	}
	return ans
}

// pairSumEnumeration 解法二：转数组 + 双指针（对照用）
// 把值存进数组，然后 i 从左、n-1-i 从右配对求和。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func pairSumEnumeration(head *ListNode) int {
	var vals []int
	for p := head; p != nil; p = p.Next {
		vals = append(vals, p.Val)
	}
	ans := 0
	for i, j := 0, len(vals)-1; i < j; i, j = i+1, j-1 {
		ans = max(ans, vals[i]+vals[j])
	}
	return ans
}
