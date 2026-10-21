package uniquepaths

// 62. 不同路径
// https://leetcode.cn/problems/unique-paths/
// 一个机器人位于 m x n 网格的左上角，每次只能向下或向右移动一步，
// 问到达网格右下角总共有多少条不同的路径。
//
// 核心思路：dp[i][j] 表示从起点走到 (i, j) 的路径数。
// 只能从上方或左方到达，转移：dp[i][j] = dp[i-1][j] + dp[i][j-1]。
// 第一行、第一列都只有 1 条路径。按行从上到下、行内从左到右遍历。
// 每行只依赖上一行和本行左侧，可压缩为一维：dp[j] += dp[j-1]
// （更新前 dp[j] 是上一行的值，dp[j-1] 是本行已更新的值）。

// uniquePaths 解法一：一维滚动数组 DP ⭐ 最优方案
// 时间复杂度 O(m * n)，空间复杂度 O(n)。
func uniquePaths(m, n int) int {
	dp := make([]int, n)
	for j := range dp {
		dp[j] = 1 // 第一行全为 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1] // 上方 + 左方
		}
	}
	return dp[n-1]
}

// uniquePathsComb 解法二：组合数学（对照用）
// 总共要走 m-1 步向下、n-1 步向右，共 m+n-2 步，答案为 C(m+n-2, m-1)。
// 边乘边除：res = res * (n-1+i) / i，每一步结果都是整数（为 C(n-1+i, i)），不会出现小数。
// 时间复杂度 O(min(m, n))，空间复杂度 O(1)。
func uniquePathsComb(m, n int) int {
	if m > n {
		m, n = n, m // 取较小者循环
	}
	res := 1
	for i := 1; i < m; i++ {
		res = res * (n - 1 + i) / i
	}
	return res
}
