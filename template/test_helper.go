package template

import (
	"fmt"
	"strings"
)

// ListNode LeetCode 单链表节点。
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode LeetCode 二叉树节点。
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildList 由切片构建链表。
func BuildList(vals []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// ListToSlice 将链表转为切片，便于测试比较。
func ListToSlice(head *ListNode) []int {
	res := []int{}
	for ; head != nil; head = head.Next {
		res = append(res, head.Val)
	}
	return res
}

// ListString 链表转字符串，如 1->2->3。
func ListString(head *ListNode) string {
	var sb strings.Builder
	for cur := head; cur != nil; cur = cur.Next {
		if cur != head {
			sb.WriteString("->")
		}
		fmt.Fprint(&sb, cur.Val)
	}
	return sb.String()
}

// PrintList 打印链表。
func PrintList(head *ListNode) {
	fmt.Println(ListString(head))
}

// PrintArray 打印一维数组。
func PrintArray[T any](a []T) {
	fmt.Println(a)
}

// PrintMatrix 逐行打印二维数组。
func PrintMatrix[T any](m [][]T) {
	for _, row := range m {
		fmt.Println(row)
	}
}
