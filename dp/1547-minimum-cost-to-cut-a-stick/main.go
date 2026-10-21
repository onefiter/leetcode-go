package mincutcost

import "slices"

// 1547. 切棍子的最小成本
// https://leetcode.cn/problems/minimum-cost-to-cut-a-stick/
// 有一根长度为 n 的木棍，刻度从 0 到 n。数组 cuts[i] 表示需要切开的位置，
// 每次切割的成本是当前所切木棍的长度，切割顺序可以任意，返回切完所有位置的最小总成本。
//
// 核心思路：区间 DP。把 cuts 排序并在两端补上 0 和 n，得到刻度数组 c（长度 m）。
// dp[i][j] 表示把 c[i]..c[j] 这一段内部所有切点切完的最小成本。
// 枚举这段中第一刀切在 c[k]（i < k < j），成本为本段长度 c[j]-c[i]，之后左右两段独立。
// 转移：dp[i][j] = min(dp[i][k] + dp[k][j]) + c[j] - c[i]，相邻刻度 dp[i][i+1] = 0。
// 遍历顺序：i 从大到小、j 从小到大（与多边形三角剖分相同）。答案为 dp[0][m-1]。

// minCost 解法一：排序 + 自底向上区间 DP ⭐ 最优方案
// 时间复杂度 O(k^3)，空间复杂度 O(k^2)，k 为 cuts 长度（与 n 无关）。
func minCost(n int, cuts []int) int {
	c := append([]int{0, n}, cuts...)
	slices.Sort(c)
	m := len(c)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := m - 3; i >= 0; i-- {
		for j := i + 2; j < m; j++ {
			best := int(^uint(0) >> 1)
			for k := i + 1; k < j; k++ {
				best = min(best, dp[i][k]+dp[k][j])
			}
			dp[i][j] = best + c[j] - c[i]
		}
	}
	return dp[0][m-1]
}

// minCostMemo 解法二：记忆化递归（对照用）
// dfs(i, j) 按定义枚举第一刀位置，j-i < 2 时中间没有切点返回 0。
// 注意：这里直接对 cuts 原地排序（会修改入参），测试中传入副本。
// 时间复杂度 O(k^3 + k log k)，空间复杂度 O(k^2)。
func minCostMemo(n int, cuts []int) int {
	slices.Sort(cuts)
	c := make([]int, 0, len(cuts)+2)
	c = append(c, 0)
	c = append(c, cuts...)
	c = append(c, n)
	m := len(c)
	memo := make([][]int, m)
	for i := range memo {
		memo[i] = make([]int, m)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if j-i < 2 {
			return 0
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		best := int(^uint(0) >> 1)
		for k := i + 1; k < j; k++ {
			best = min(best, dfs(i, k)+dfs(k, j))
		}
		memo[i][j] = best + c[j] - c[i]
		return memo[i][j]
	}
	return dfs(0, m-1)
}
