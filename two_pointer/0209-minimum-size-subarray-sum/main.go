package minsubarraylen

import "sort"

// 209. 长度最小的子数组
// https://leetcode.cn/problems/minimum-size-subarray-sum/
// 给定正整数数组 nums 和正整数 target，找出和 >= target 的长度最小的连续子数组，
// 返回其长度；不存在则返回 0。例如 target=7, nums=[2,3,1,2,4,3] -> 2（[4,3]）。
//
// 核心思路：元素都是正数，窗口右端扩张时和单调增加、左端收缩时和单调减少，
// 因此可以用滑动窗口：右指针不断加入元素，一旦和 >= target 就尽量收缩左指针并更新答案。

// minSubArrayLen 解法一：滑动窗口 ⭐ 最优方案
// 每个元素最多进窗口一次、出窗口一次。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	best := n + 1
	sum, left := 0, 0
	for right, x := range nums {
		sum += x
		for sum >= target {
			best = min(best, right-left+1)
			sum -= nums[left]
			left++
		}
	}
	if best == n+1 {
		return 0
	}
	return best
}

// minSubArrayLenBinarySearch 解法二：前缀和 + 二分
// pre[i] 为前 i 个元素之和，正数保证 pre 严格递增。
// 对每个起点 i，二分找最小的 j 使 pre[j] >= pre[i]+target，长度为 j-i。
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func minSubArrayLenBinarySearch(target int, nums []int) int {
	n := len(nums)
	pre := make([]int, n+1)
	for i, x := range nums {
		pre[i+1] = pre[i] + x
	}
	best := n + 1
	for i := range n {
		j := sort.SearchInts(pre, pre[i]+target)
		if j <= n {
			best = min(best, j-i)
		}
	}
	if best == n+1 {
		return 0
	}
	return best
}

// minSubArrayLenEnumeration 解法三：暴力枚举（对照用）
// 枚举每个起点，向右累加直到和 >= target。
// 时间复杂度 O(n²)，空间复杂度 O(1)。
func minSubArrayLenEnumeration(target int, nums []int) int {
	best := 0
	for i := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum >= target {
				if best == 0 || j-i+1 < best {
					best = j - i + 1
				}
				break
			}
		}
	}
	return best
}
