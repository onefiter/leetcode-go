package maxproduct

// 152. 乘积最大子数组
// https://leetcode.cn/problems/maximum-product-subarray/
// 给一个整数数组 nums，找出乘积最大的非空连续子数组，返回该乘积（结果保证在 32 位整数范围内）。
//
// 核心思路：与最大子数组和不同，负数会让最大值和最小值互换，所以要同时维护两个状态：
// maxDp[i] / minDp[i] 表示「以 nums[i] 结尾」的子数组的最大 / 最小乘积。
// 转移：maxDp[i] = max(nums[i], maxDp[i-1]*nums[i], minDp[i-1]*nums[i])，
//       minDp[i] = min(nums[i], maxDp[i-1]*nums[i], minDp[i-1]*nums[i])。
// 答案为所有 maxDp[i] 的最大值；只依赖上一项，可压缩成两个变量。

// maxProduct 解法一：同时维护最大/最小乘积的滚动 DP ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func maxProduct(nums []int) int {
	curMax, curMin, best := nums[0], nums[0], nums[0]
	for _, x := range nums[1:] {
		// 用旧值同时计算，负数会把最小值翻成最大值
		a, b := curMax*x, curMin*x
		curMax = max(x, a, b)
		curMin = min(x, a, b)
		best = max(best, curMax)
	}
	return best
}

// maxProductBrute 解法二：暴力枚举所有子数组（对照用）
// 固定左端点 i，向右扩展右端点 j 时累乘，逐个更新最大值。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func maxProductBrute(nums []int) int {
	best := nums[0]
	for i := 0; i < len(nums); i++ {
		prod := 1
		for j := i; j < len(nums); j++ {
			prod *= nums[j]
			best = max(best, prod)
		}
	}
	return best
}
