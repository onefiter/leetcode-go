package climbstairs

// 70. 爬楼梯
// https://leetcode.cn/problems/climbing-stairs/
// 需要 n 阶才能到达楼顶，每次可以爬 1 或 2 个台阶，求有多少种不同的方法爬到楼顶（1 <= n <= 45）。
//
// 核心思路：dp[i] 表示爬到第 i 阶的方法数。
// 最后一步要么从 i-1 跨 1 阶上来，要么从 i-2 跨 2 阶上来，所以 dp[i] = dp[i-1] + dp[i-2]，
// 初始 dp[0] = 1, dp[1] = 1。dp[i] 只依赖前两项，用两个变量滚动即可。

// climbStairs 解法一：滚动 DP ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func climbStairs(n int) int {
	prev, cur := 1, 1 // dp[i-1], dp[i]
	for i := 2; i <= n; i++ {
		prev, cur = cur, prev+cur
	}
	return cur
}

// climbStairsTable 解法二：完整 dp 表（对照用）
// 显式开 n+1 长度的数组逐项填表，状态含义一目了然。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func climbStairsTable(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1]
		if i >= 2 {
			dp[i] += dp[i-2]
		}
	}
	return dp[n]
}
