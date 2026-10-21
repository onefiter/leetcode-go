package getdecimalvalue

import "leetcode-go/template"

// 1290. 二进制链表转整数
// https://leetcode.cn/problems/convert-binary-number-in-a-linked-list-to-integer/
// 链表每个节点是 0 或 1，从头到尾表示一个二进制数（头是最高位），返回它的十进制值。
// 例如 1->0->1 表示 101，答案 5。
//
// 核心思路：和把字符串 "101" 转成数字一样，从高位往低位读，
// 每读一位就 ans = ans*2 + bit（即左移一位再或上新位）。

type ListNode = template.ListNode

// getDecimalValue 解法一：一次遍历，左移累加 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func getDecimalValue(head *ListNode) int {
	ans := 0
	for p := head; p != nil; p = p.Next {
		ans = ans<<1 | p.Val
	}
	return ans
}

// getDecimalValueEnumeration 解法二：先求长度，再按位权相加（对照用）
// 第 i 个节点（从 0 开始）的位权是 2^(n-1-i)。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func getDecimalValueEnumeration(head *ListNode) int {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	ans := 0
	i := 0
	for p := head; p != nil; p = p.Next {
		if p.Val == 1 {
			ans += 1 << (n - 1 - i)
		}
		i++
	}
	return ans
}
