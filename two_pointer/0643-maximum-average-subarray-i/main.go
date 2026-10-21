package maxaverage

// 643. 子数组最大平均数 I
// https://leetcode.cn/problems/maximum-average-subarray-i/
// 给你一个由 n 个元素组成的整数数组 nums 和一个整数 k（1 <= k <= n），
// 找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。
// 例如 nums=[1,12,-5,-6,50,3], k=4 -> 12.75（(12-5-6+50)/4）。
//
// 一句话理解题意：长度固定为 k，平均数最大等价于「和最大」，最后除以 k 即可。
//
// 核心思路：定长滑动窗口。窗口右移一格时，加上新进来的元素、减去离开的元素，
// 每次 O(1) 得到新窗口的和。

// findMaxAverage 解法一：定长滑动窗口 ⭐ 最优方案
// 先算前 k 个元素的和，之后每次 sum += nums[i] - nums[i-k]，记录最大和。
// 全程用整数累加避免浮点误差，最后再转 float64 除以 k。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	best := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		best = max(best, sum)
	}
	return float64(best) / float64(k)
}

// findMaxAveragePrefixSum 解法二：前缀和
// pre[i] 表示前 i 个元素之和，窗口 [i-k, i) 的和为 pre[i]-pre[i-k]。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func findMaxAveragePrefixSum(nums []int, k int) float64 {
	n := len(nums)
	pre := make([]int, n+1)
	for i, x := range nums {
		pre[i+1] = pre[i] + x
	}
	best := pre[k]
	for i := k + 1; i <= n; i++ {
		best = max(best, pre[i]-pre[i-k])
	}
	return float64(best) / float64(k)
}

// findMaxAverageEnumeration 解法三：枚举每个起点重新求和（对照用）
// 对每个起点 i，暴力累加 nums[i..i+k-1]。
// 时间复杂度 O(n·k)，空间复杂度 O(1)。
func findMaxAverageEnumeration(nums []int, k int) float64 {
	best := 0
	for i := 0; i+k <= len(nums); i++ {
		sum := 0
		for j := i; j < i+k; j++ {
			sum += nums[j]
		}
		if i == 0 || sum > best {
			best = sum
		}
	}
	return float64(best) / float64(k)
}
