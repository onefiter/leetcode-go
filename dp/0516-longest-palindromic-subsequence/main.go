package lps

// 516. 最长回文子序列
// https://leetcode.cn/problems/longest-palindromic-subsequence/
// 给一个字符串 s，找出其中最长的回文子序列，并返回该序列的长度。
// 子序列：不改变剩余字符顺序的情况下，删除某些字符或不删除任何字符形成的序列。
//
// 核心思路：区间 DP。dp[i][j] 表示 s[i..j] 内最长回文子序列的长度，dp[i][i] = 1。
// 转移：s[i] == s[j] 时 dp[i][j] = dp[i+1][j-1] + 2（两端一起加入）；
// 否则 dp[i][j] = max(dp[i+1][j], dp[i][j-1])（舍弃一端）。
// 遍历顺序：依赖下方、左方、左下方，因此 i 倒序、j 从 i+1 正序。答案为 dp[0][n-1]。

// longestPalindromeSubseq 解法一：区间 DP + 一维滚动数组 ⭐ 最优方案
// dp[j] 在第 i 轮更新前存的是 dp[i+1][j]，用 prev 暂存左下角 dp[i+1][j-1]。
// 时间复杂度 O(n^2)，空间复杂度 O(n)。
func longestPalindromeSubseq(s string) int {
	n := len(s)
	dp := make([]int, n)
	for i := n - 1; i >= 0; i-- {
		dp[i] = 1
		prev := 0 // dp[i+1][i]，空区间为 0
		for j := i + 1; j < n; j++ {
			tmp := dp[j] // 覆盖前是 dp[i+1][j]
			if s[i] == s[j] {
				dp[j] = prev + 2
			} else {
				dp[j] = max(dp[j], dp[j-1])
			}
			prev = tmp
		}
	}
	return dp[n-1]
}

// longestPalindromeSubseq2D 解法二：区间 DP 二维表（对照用）
// 时间复杂度 O(n^2)，空间复杂度 O(n^2)。
func longestPalindromeSubseq2D(s string) int {
	n := len(s)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < n; j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}
	return dp[0][n-1]
}
