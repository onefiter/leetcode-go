package maxsubarray

// 53. 最大子数组和
// https://leetcode.cn/problems/maximum-subarray/
// 给一个整数数组 nums，找出一个具有最大和的连续子数组（至少包含一个元素），返回其最大和。
//
// 核心思路：dp[i] 表示「以 nums[i] 结尾」的最大子数组和。
// 转移：dp[i] = max(dp[i-1]+nums[i], nums[i])——前面的和为负就丢掉，从自己重新开始；
// 答案是所有 dp[i] 的最大值。dp[i] 只依赖 dp[i-1]，可压缩成一个变量（即 Kadane 算法）。

// maxSubArray 解法一：滚动 DP（Kadane） ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func maxSubArray(nums []int) int {
	cur, best := nums[0], nums[0]
	for _, x := range nums[1:] {
		cur = max(cur+x, x) // 接在前面 or 自己另起炉灶
		best = max(best, cur)
	}
	return best
}

// maxSubArrayDivide 解法二：分治（对照用）
// 区间 [l, r] 的最大子数组要么完全在左半、要么完全在右半、要么跨过中点；
// 跨中点的部分 = 从 mid 向左的最大后缀和 + 从 mid+1 向右的最大前缀和。
// 时间复杂度 O(n log n)，空间复杂度 O(log n)（递归深度）。
func maxSubArrayDivide(nums []int) int {
	var solve func(l, r int) int
	solve = func(l, r int) int {
		if l == r {
			return nums[l]
		}
		mid := l + (r-l)/2
		leftBest, sum := nums[mid], 0
		for i := mid; i >= l; i-- {
			sum += nums[i]
			leftBest = max(leftBest, sum)
		}
		rightBest := nums[mid+1]
		sum = 0
		for i := mid + 1; i <= r; i++ {
			sum += nums[i]
			rightBest = max(rightBest, sum)
		}
		return max(solve(l, mid), solve(mid+1, r), leftBest+rightBest)
	}
	return solve(0, len(nums)-1)
}
