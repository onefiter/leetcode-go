package lcs

// 1143. 最长公共子序列
// https://leetcode.cn/problems/longest-common-subsequence/
// 给定两个字符串 text1 和 text2，返回这两个字符串的最长公共子序列的长度。
// 子序列是在不改变相对顺序的情况下删除某些字符（也可以不删除）得到的新字符串。
//
// 核心思路：dp[i][j] 表示 text1 前 i 个字符与 text2 前 j 个字符的 LCS 长度。
// 转移：text1[i-1] == text2[j-1] 时 dp[i][j] = dp[i-1][j-1] + 1，
// 否则 dp[i][j] = max(dp[i-1][j], dp[i][j-1])。边界 dp[0][*] = dp[*][0] = 0。
// i、j 都从小到大遍历。压缩为一维时，dp[j-1] 已是本行新值，
// 需用变量 prev 额外保存「左上角」dp[i-1][j-1]。

// longestCommonSubsequence 解法一：一维滚动数组 + 左上角变量 ⭐ 最优方案
// 时间复杂度 O(m * n)，空间复杂度 O(min(m, n))。
func longestCommonSubsequence(text1, text2 string) int {
	if len(text1) < len(text2) {
		text1, text2 = text2, text1 // 让 text2 较短，数组更小
	}
	n := len(text2)
	dp := make([]int, n+1)
	for i := 1; i <= len(text1); i++ {
		prev := 0 // dp[i-1][0]
		for j := 1; j <= n; j++ {
			tmp := dp[j] // 保存 dp[i-1][j]，作为下一个 j 的左上角
			if text1[i-1] == text2[j-1] {
				dp[j] = prev + 1
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = tmp
		}
	}
	return dp[n]
}

// longestCommonSubsequence2D 解法二：二维 dp 表（对照用）
// 直接按定义填表，便于回溯出具体的公共子序列。
// 时间复杂度 O(m * n)，空间复杂度 O(m * n)。
func longestCommonSubsequence2D(text1, text2 string) int {
	m, n := len(text1), len(text2)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[m][n]
}
