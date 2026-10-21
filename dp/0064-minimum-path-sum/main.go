package minpathsum

// 64. 最小路径和
// https://leetcode.cn/problems/minimum-path-sum/
// 给一个包含非负整数的 m x n 网格 grid，找一条从左上角到右下角的路径，
// 使路径上的数字总和最小。每次只能向下或向右移动一步。
//
// 核心思路：dp[i][j] 表示从起点走到 (i, j) 的最小路径和。
// 只能从上方或左方到达，转移：dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]。
// 第一行只能从左边来，第一列只能从上边来。按行从上到下、行内从左到右遍历。
// 每行只依赖上一行与本行左侧，可压缩为一维数组 dp[j] = min(dp[j], dp[j-1]) + grid[i][j]。

// minPathSum 解法一：一维滚动数组 DP ⭐ 最优方案
// 时间复杂度 O(m * n)，空间复杂度 O(n)（不修改原数组）。
func minPathSum(grid [][]int) int {
	n := len(grid[0])
	dp := make([]int, n)
	dp[0] = grid[0][0]
	for j := 1; j < n; j++ {
		dp[j] = dp[j-1] + grid[0][j] // 第一行只能从左边来
	}
	for i := 1; i < len(grid); i++ {
		dp[0] += grid[i][0] // 第一列只能从上边来
		for j := 1; j < n; j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j] // dp[j] 为上方，dp[j-1] 为左方
		}
	}
	return dp[n-1]
}

// minPathSum2D 解法二：二维 dp 表（对照用）
// 完整保存每个格子的最小路径和，便于回溯具体路径。
// 时间复杂度 O(m * n)，空间复杂度 O(m * n)。
func minPathSum2D(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch {
			case i == 0 && j == 0:
				dp[i][j] = grid[0][0]
			case i == 0:
				dp[i][j] = dp[i][j-1] + grid[i][j]
			case j == 0:
				dp[i][j] = dp[i-1][j] + grid[i][j]
			default:
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}
