package addtwonumbers

import "leetcode-go/template"

// 2. 两数相加
// https://leetcode.cn/problems/add-two-numbers/
// 两个非空链表按「逆序」存储两个非负整数（个位在表头），每个节点存一位数字，
// 求两数之和，同样以逆序链表返回。例如 [2,4,3] + [5,6,4] = [7,0,8]（342+465=807）。
//
// 核心思路：逆序存储正好对应竖式加法「从个位开始」，
// 两条链表同步往后走，逐位相加并维护进位 carry，任意一条没走完或还有进位就继续。

type ListNode = template.ListNode

// addTwoNumbers 解法一：模拟竖式加法 ⭐ 最优方案
// 用哨兵节点 dummy 简化头节点处理；每一位 sum = a + b + carry，
// 新节点存 sum%10，进位 carry = sum/10。循环条件包含 carry != 0，处理最高位进位。
// 时间复杂度 O(max(m, n))，空间复杂度 O(1)（不计结果链表）。
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	carry := 0
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		carry = sum / 10
	}
	return dummy.Next
}

// addTwoNumbersRecursive 解法二：递归
// 把「当前位相加 + 进位」交给本层，剩下的位递归处理，进位通过参数传下去。
// 时间复杂度 O(max(m, n))，空间复杂度 O(max(m, n))（递归栈）。
func addTwoNumbersRecursive(l1 *ListNode, l2 *ListNode) *ListNode {
	return addWithCarry(l1, l2, 0)
}

func addWithCarry(l1, l2 *ListNode, carry int) *ListNode {
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	sum := carry
	var n1, n2 *ListNode
	if l1 != nil {
		sum += l1.Val
		n1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		n2 = l2.Next
	}
	return &ListNode{Val: sum % 10, Next: addWithCarry(n1, n2, sum/10)}
}

// addTwoNumbersEnumeration 解法三：转成数字切片逐位相加（对照用）
// 先把两条链表各自读成切片，再在切片上做竖式加法，最后建新链表。
// 思路最朴素，但多用了 O(m+n) 的额外空间。
// 时间复杂度 O(max(m, n))，空间复杂度 O(m + n)。
func addTwoNumbersEnumeration(l1 *ListNode, l2 *ListNode) *ListNode {
	a, b := template.ListToSlice(l1), template.ListToSlice(l2)
	res := []int{}
	carry := 0
	for i := 0; i < len(a) || i < len(b) || carry != 0; i++ {
		sum := carry
		if i < len(a) {
			sum += a[i]
		}
		if i < len(b) {
			sum += b[i]
		}
		res = append(res, sum%10)
		carry = sum / 10
	}
	return template.BuildList(res)
}
