package rotatelist

import "leetcode-go/template"

// 61. 旋转链表
// https://leetcode.cn/problems/rotate-list/
// 把链表每个节点向右移动 k 个位置。例如 [1,2,3,4,5], k=2 -> [4,5,1,2,3]。
//
// 核心思路：向右旋转 k 次 = 把最后 k%n 个节点整体搬到表头。
// 先求长度 n 并找到尾节点，把链表首尾相连成环，
// 再从头走 n-k%n-1 步找到新尾巴，在那里断开，新尾巴的下一个就是新头。

type ListNode = template.ListNode

// rotateRight 解法一：闭合成环再断开 ⭐ 最优方案
// k 可能远大于 n，先取模；k%n == 0 时链表不变，直接返回。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}
	n, tail := 1, head
	for tail.Next != nil {
		tail = tail.Next
		n++
	}
	k %= n
	if k == 0 {
		return head
	}
	tail.Next = head // 成环
	newTail := head
	for i := 0; i < n-k-1; i++ {
		newTail = newTail.Next
	}
	newHead := newTail.Next
	newTail.Next = nil // 断环
	return newHead
}

// rotateRightEnumeration 解法二：逐次右移一位（对照用）
// 每次把尾节点摘下来放到表头，重复 k%n 次。
// 每次都要走到倒数第二个节点，最坏 O(n²)。
// 时间复杂度 O(n·min(k, n))，空间复杂度 O(1)。
func rotateRightEnumeration(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := 0
	for cur := head; cur != nil; cur = cur.Next {
		n++
	}
	for range k % n {
		prev := head
		for prev.Next.Next != nil {
			prev = prev.Next
		}
		last := prev.Next
		prev.Next = nil
		last.Next = head
		head = last
	}
	return head
}
