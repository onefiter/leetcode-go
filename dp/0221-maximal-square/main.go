package maxsquare

// 221. 最大正方形
// https://leetcode.cn/problems/maximal-square/
// 在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。
//
// 核心思路：dp[i][j] 表示以 (i-1, j-1) 为右下角、只含 '1' 的最大正方形边长。
// 转移：若 matrix[i-1][j-1] == '1'，dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1，
// 即受上、左、左上三个方向中最短的那条边限制；否则为 0。
// 遍历顺序：i、j 均正序，依赖上、左、左上；答案为最大边长的平方。

// maximalSquare 解法一：DP + 一维滚动数组 ⭐ 最优方案
// 用 prev 暂存左上角 dp[i-1][j-1]。
// 时间复杂度 O(m * n)，空间复杂度 O(n)。
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	n := len(matrix[0])
	dp := make([]int, n+1)
	maxSide := 0
	for i := range matrix {
		prev := 0 // 左上角 dp[i-1][0]
		for j := 1; j <= n; j++ {
			tmp := dp[j] // 覆盖前是 dp[i-1][j]
			if matrix[i][j-1] == '1' {
				dp[j] = min(dp[j], dp[j-1], prev) + 1
				maxSide = max(maxSide, dp[j])
			} else {
				dp[j] = 0
			}
			prev = tmp
		}
	}
	return maxSide * maxSide
}

// maximalSquareBrute 解法二：暴力枚举（对照用）
// 以每个 '1' 为左上角，逐步扩大边长 k，每次只需检查新增的最下一行和最右一列。
// 时间复杂度 O(m * n * min(m, n)^2)，空间复杂度 O(1)。
func maximalSquareBrute(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	maxSide := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] != '1' {
				continue
			}
			side := 1
			for ok := true; ok && i+side < m && j+side < n; {
				// 检查新增的第 side 行和第 side 列
				for k := 0; k <= side; k++ {
					if matrix[i+side][j+k] != '1' || matrix[i+k][j+side] != '1' {
						ok = false
						break
					}
				}
				if ok {
					side++
				}
			}
			maxSide = max(maxSide, side)
		}
	}
	return maxSide * maxSide
}
