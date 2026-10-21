package sortedsquares

import "slices"

// 977. 有序数组的平方
// https://leetcode.cn/problems/squares-of-a-sorted-array/
// 给定按非递减顺序排列的整数数组 nums（可含负数），返回每个数平方组成的新数组，也按非递减排序。
// 例如 [-4,-1,0,3,10] -> [0,1,9,16,100]。
//
// 核心思路：平方后最大值一定出现在数组两端（绝对值最大的负数或正数），
// 所以用左右双指针比较两端平方，把较大者从结果数组末尾往前填。

// sortedSquares 解法一：对撞双指针，从后往前填 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)（不计返回数组）。
func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	l, r := 0, n-1
	for k := n - 1; k >= 0; k-- {
		a, b := nums[l]*nums[l], nums[r]*nums[r]
		if a > b {
			res[k] = a
			l++
		} else {
			res[k] = b
			r--
		}
	}
	return res
}

// sortedSquaresEnumeration 解法二：直接平方后排序（对照用）
// 时间复杂度 O(n log n)，空间复杂度 O(log n)（排序栈）。
func sortedSquaresEnumeration(nums []int) []int {
	res := make([]int, len(nums))
	for i, x := range nums {
		res[i] = x * x
	}
	slices.Sort(res)
	return res
}
