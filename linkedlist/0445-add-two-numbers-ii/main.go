package addtwonumbers

import "leetcode-go/template"

// 445. 两数相加 II
// https://leetcode.cn/problems/add-two-numbers-ii/
// 两个非空链表表示两个非负整数，数字最高位在链表开头，每个节点存一位。
// 将两数相加，返回同样「高位在前」的结果链表。除 0 本身外不会有前导 0。
//
// 核心思路：加法要从低位（链表尾部）开始算并处理进位，但单链表只能从头走。
//   - 栈：把两条链表的数字分别压栈，弹栈就是从低位到高位；结果用「头插法」组装，天然高位在前；
//   - 反转：先把两条链表反转成低位在前，按 2. 两数相加的方法逐位相加，再把结果反转回来。

type ListNode = template.ListNode

// addTwoNumbers 解法一：双栈 + 头插法 ⭐ 最优方案（不修改输入链表）
// 每次从两个栈各弹一位（空栈记 0），加上进位得到当前位，新节点插到结果头部。
// 时间复杂度 O(m+n)，空间复杂度 O(m+n)（两个栈）。
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	s1, s2 := toStack(l1), toStack(l2)
	var head *ListNode
	carry := 0
	for len(s1) > 0 || len(s2) > 0 || carry > 0 {
		sum := carry
		if len(s1) > 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		head = &ListNode{Val: sum % 10, Next: head} // 头插
		carry = sum / 10
	}
	return head
}

func toStack(head *ListNode) []int {
	s := []int{}
	for ; head != nil; head = head.Next {
		s = append(s, head.Val)
	}
	return s
}

// addTwoNumbersReverse 解法二：反转链表后逐位相加
// 反转两条输入链表（低位在前），同时遍历相加并处理进位，最后反转结果。
// 会原地修改输入链表。
// 时间复杂度 O(m+n)，空间复杂度 O(1)（不计结果链表）。
func addTwoNumbersReverse(l1, l2 *ListNode) *ListNode {
	l1, l2 = reverse(l1), reverse(l2)
	dummy := &ListNode{}
	tail := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		tail.Next = &ListNode{Val: sum % 10}
		tail = tail.Next
		carry = sum / 10
	}
	return reverse(dummy.Next)
}

func reverse(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}

// addTwoNumbersEnumeration 解法三：转成数字切片后按竖式相加（对照用）
// 把两条链表转成切片，从切片末尾逐位相加得到结果切片（低位在前），倒序后建链表。
// 不能直接转成 int 相加，因为链表最长 100 位会溢出。
// 时间复杂度 O(m+n)，空间复杂度 O(m+n)。
func addTwoNumbersEnumeration(l1, l2 *ListNode) *ListNode {
	a, b := template.ListToSlice(l1), template.ListToSlice(l2)
	res := []int{}
	carry := 0
	for i, j := len(a)-1, len(b)-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		sum := carry
		if i >= 0 {
			sum += a[i]
		}
		if j >= 0 {
			sum += b[j]
		}
		res = append(res, sum%10)
		carry = sum / 10
	}
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}
	return template.BuildList(res)
}
