package lis

import "sort"

// 300. 最长递增子序列（DP 版）
// https://leetcode.cn/problems/longest-increasing-subsequence/
// 给一个整数数组 nums，返回其中最长严格递增子序列的长度（子序列可以不连续）。
//
// 核心思路：dp[i] 表示「以 nums[i] 结尾」的最长递增子序列长度。
// 枚举 i 前面的每个 j，若 nums[j] < nums[i]，则 nums[i] 可以接在以 nums[j] 结尾的子序列后面，
// 转移方程 dp[i] = max(dp[j] + 1)（0 <= j < i 且 nums[j] < nums[i]），初始 dp[i] = 1（只含自己）。
// 答案是所有 dp[i] 的最大值（最长子序列不一定以最后一个元素结尾），时间 O(n²)。
// 进阶：维护 tails[k] =「长度为 k+1 的递增子序列的最小结尾」，tails 单调递增，
// 每个数用二分找到第一个 >= 它的位置替换（或追加），可优化到 O(n log n)。

// lengthOfLIS 解法一：贪心 + 二分（tails 数组） ⭐ 最优方案
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func lengthOfLIS(nums []int) int {
	tails := []int{}
	for _, x := range nums {
		i := sort.SearchInts(tails, x) // 第一个 >= x 的位置，保证严格递增
		if i == len(tails) {
			tails = append(tails, x)
		} else {
			tails[i] = x // 同样长度下结尾更小，后续更容易接
		}
	}
	return len(tails)
}

// lengthOfLISDP 解法二：经典 O(n²) 动态规划
// dp[i] 为以 nums[i] 结尾的 LIS 长度，枚举前驱 j 转移。
// 时间复杂度 O(n²)，空间复杂度 O(n)。
func lengthOfLISDP(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	res := 0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		res = max(res, dp[i])
	}
	return res
}
