package triangulation

// 1039. 多边形三角剖分的最低得分
// https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/
// 有一个凸的 n 边形，values[i] 是第 i 个顶点的值（顺时针顺序）。将多边形剖分为 n-2 个三角形，
// 每个三角形的值是其顶点值的乘积，剖分的分数是所有三角形值之和，返回可能的最低分数。
//
// 核心思路：区间 DP。dp[i][j] 表示顶点 i..j 构成的子多边形（以边 i-j 为底）的最低得分。
// 边 i-j 必然属于某个三角形 (i, k, j)，i < k < j，它把区间切成 [i, k] 和 [k, j] 两部分。
// 转移：dp[i][j] = min(dp[i][k] + dp[k][j] + v[i]*v[k]*v[j])，相邻顶点 dp[i][i+1] = 0。
// 遍历顺序：dp[i][j] 依赖更大的 i（dp[k][j]）和更小的 j（dp[i][k]），
// 所以 i 从大到小、j 从小到大遍历。答案为 dp[0][n-1]。

// minScoreTriangulation 解法一：自底向上区间 DP ⭐ 最优方案
// 时间复杂度 O(n^3)，空间复杂度 O(n^2)。
func minScoreTriangulation(values []int) int {
	n := len(values)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 3; i >= 0; i-- { // i 倒序
		for j := i + 2; j < n; j++ { // j 正序
			best := int(^uint(0) >> 1)
			for k := i + 1; k < j; k++ {
				best = min(best, dp[i][k]+dp[k][j]+values[i]*values[k]*values[j])
			}
			dp[i][j] = best
		}
	}
	return dp[0][n-1]
}

// minScoreTriangulationMemo 解法二：记忆化递归（对照用）
// dfs(i, j) 直接按定义枚举三角形顶点 k，j-i < 2 时无法构成三角形返回 0。
// 用 memo（-1 表示未计算）避免重复计算子区间。
// 时间复杂度 O(n^3)，空间复杂度 O(n^2)（memo + 递归栈 O(n)）。
func minScoreTriangulationMemo(values []int) int {
	n := len(values)
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
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
			best = min(best, dfs(i, k)+dfs(k, j)+values[i]*values[k]*values[j])
		}
		memo[i][j] = best
		return best
	}
	return dfs(0, n-1)
}
