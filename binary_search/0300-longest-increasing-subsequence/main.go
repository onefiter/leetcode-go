package lis

import "sort"

// 300. 最长递增子序列
// https://leetcode.cn/problems/longest-increasing-subsequence/
// 给定整数数组 nums，求其中最长「严格递增」子序列的长度（子序列可以不连续）。
// 例如 [10,9,2,5,3,7,101,18] -> 4（如 [2,3,7,101]）。
//
// 一句话理解题意：从数组里按原顺序挑出尽量多的数，使它们严格递增。
//
// 核心思路：
//   - DP：dp[i] 表示以 nums[i] 结尾的最长递增子序列长度，枚举前面比它小的元素转移。
//   - 贪心+二分：维护 tails[k] = 长度为 k+1 的递增子序列的最小结尾，
//     结尾越小越容易接上后面的数；tails 严格递增，可以二分。

// lengthOfLIS 解法一：贪心 + 二分 ⭐ 最优方案
// 遍历每个 x，在 tails 中找第一个 >= x 的位置 i：
//   - i == len(tails)：x 比所有结尾都大，接在最长的后面，长度 +1；
//   - 否则用 x 替换 tails[i]，让长度为 i+1 的子序列结尾更小。
//
// 最终 tails 的长度就是答案（tails 本身不一定是真实的子序列）。
// 时间复杂度 O(n log n)，空间复杂度 O(n)。
func lengthOfLIS(nums []int) int {
	tails := make([]int, 0, len(nums))
	for _, x := range nums {
		i := sort.SearchInts(tails, x) // 第一个 >= x 的下标，保证严格递增
		if i == len(tails) {
			tails = append(tails, x)
		} else {
			tails[i] = x
		}
	}
	return len(tails)
}

// lengthOfLISDP 解法二：动态规划（对照用）
// dp[i] = 1 + max(dp[j])，其中 j < i 且 nums[j] < nums[i]；没有这样的 j 时 dp[i] = 1。
// 答案是所有 dp[i] 的最大值。
// 时间复杂度 O(n^2)，空间复杂度 O(n)。
func lengthOfLISDP(nums []int) int {
	dp := make([]int, len(nums))
	ans := 0
	for i := range nums {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

// lengthOfLISEnumeration 解法三：枚举所有子序列（对照用）
// 用二进制掩码枚举每个子序列，检查是否严格递增，取最长。仅用于小规模校验。
// 时间复杂度 O(2^n · n)，空间复杂度 O(1)。
func lengthOfLISEnumeration(nums []int) int {
	n := len(nums)
	ans := 0
	for mask := 0; mask < 1<<n; mask++ {
		cnt, prev, ok := 0, 0, true
		for i := 0; i < n && ok; i++ {
			if mask>>i&1 == 0 {
				continue
			}
			if cnt > 0 && nums[i] <= prev {
				ok = false
			}
			prev = nums[i]
			cnt++
		}
		if ok {
			ans = max(ans, cnt)
		}
	}
	return ans
}
