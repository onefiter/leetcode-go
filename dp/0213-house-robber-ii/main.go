package rob2

// 213. 打家劫舍 II
// https://leetcode.cn/problems/house-robber-ii/
// 房屋围成一圈（第一间和最后一间相邻），相邻两间不能同时偷，求能偷到的最高金额。
//
// 核心思路：环的约束只是「首尾不能同时偷」，拆成两个线性的 198 题：
// 不偷最后一间（区间 [0, n-2]）和不偷第一间（区间 [1, n-1]），取两者较大值。
// 线性部分 dp[i] = max(dp[i-1], dp[i-2]+nums[i])，用两个变量滚动；n == 1 时需特判。

// rob 解法一：拆成两段线性 DP + 滚动变量 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	return max(robRange(nums, 0, n-2), robRange(nums, 1, n-1))
}

// robRange 对闭区间 [lo, hi] 做线性打家劫舍。
func robRange(nums []int, lo, hi int) int {
	prev, cur := 0, 0
	for i := lo; i <= hi; i++ {
		prev, cur = cur, max(cur, prev+nums[i])
	}
	return cur
}

// robTable 解法二：显式 dp 表 + 「是否偷第一间」二维状态（对照用）
// dp[k][i] 表示在第一间偷与否为 k（0 不偷、1 偷）的前提下，前 i+1 间的最大金额；
// 最后一间只在 k == 0 时允许偷。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func robTable(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	res := 0
	for k := 0; k < 2; k++ {
		dp := make([]int, n)
		dp[0] = k * nums[0] // k==1 强制偷第一间，k==0 强制不偷
		for i := 1; i < n; i++ {
			take := nums[i]
			if i >= 2 {
				take += dp[i-2]
			} else if k == 1 {
				take = 0 // 偷了第一间则第二间不能偷，退化为不偷
			}
			if i == n-1 && k == 1 {
				take = 0 // 偷了第一间则最后一间不能偷
			}
			dp[i] = max(dp[i-1], take)
		}
		res = max(res, dp[n-1])
	}
	return res
}
