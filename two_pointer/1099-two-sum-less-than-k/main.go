package twosumlessk

import "slices"

// 1099. 小于 K 的两数之和
// https://leetcode.cn/problems/two-sum-less-than-k/
// 给定数组 nums 和整数 k，找两个不同下标 i<j，使 nums[i]+nums[j] < k 且和尽可能大，
// 返回这个最大和；不存在则返回 -1。例如 nums=[34,23,1,24,75,33,54,8], k=60 -> 58（34+24）。
//
// 核心思路：排序后左右双指针。和 < k 时，left 与 (left, right] 中任意数的和都不超过当前和，
// 所以记录答案后 left++；和 >= k 时只能让 right-- 把和变小。

// twoSumLessThanK 解法一：排序 + 对撞双指针 ⭐ 最优方案
// 排序前先 Clone，避免修改调用方的数组。
// 时间复杂度 O(n log n)，空间复杂度 O(n)（拷贝）。
func twoSumLessThanK(nums []int, k int) int {
	a := slices.Clone(nums)
	slices.Sort(a)
	res := -1
	l, r := 0, len(a)-1
	for l < r {
		s := a[l] + a[r]
		if s < k {
			res = max(res, s)
			l++
		} else {
			r--
		}
	}
	return res
}

// twoSumLessThanKEnumeration 解法二：暴力枚举所有数对（对照用）
// 时间复杂度 O(n²)，空间复杂度 O(1)。
func twoSumLessThanKEnumeration(nums []int, k int) int {
	res := -1
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if s := nums[i] + nums[j]; s < k {
				res = max(res, s)
			}
		}
	}
	return res
}
