package palindrome

import "slices"

// 131. 分割回文串
// https://leetcode.cn/problems/palindrome-partitioning/
// 给一个字符串 s，把 s 分割成若干子串，使每个子串都是回文串，
// 返回所有可能的分割方案，顺序任意。
//
// 核心思路：回溯枚举每一刀切在哪里，切出的前缀必须是回文才继续。
// 先用 DP 预处理 dp[i][j] 表示 s[i..j] 是否回文（dp[i][j] = s[i]==s[j] && dp[i+1][j-1]），
// 回溯时 O(1) 判断，避免重复检查。

// partition 解法一：DP 预处理回文 + 回溯 ⭐ 最优方案
// 时间复杂度 O(n*2^n)，空间复杂度 O(n^2)。
func partition(s string) [][]string {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j < n; j++ {
			dp[i][j] = s[i] == s[j] && (j-i < 2 || dp[i+1][j-1])
		}
	}
	res := [][]string{}
	path := []string{}
	var dfs func(start int)
	dfs = func(start int) {
		if start == n {
			res = append(res, slices.Clone(path))
			return
		}
		for end := start; end < n; end++ {
			if dp[start][end] {
				path = append(path, s[start:end+1])
				dfs(end + 1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs(0)
	return res
}

// isPalindrome 双指针判断 s[l..r] 是否回文。
func isPalindrome(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

// partitionEnumeration 解法二：位掩码枚举所有切分方式（对照用）
// 长度为 n 的串有 n-1 个缝隙，mask 第 i 位为 1 表示在 s[i] 后切一刀，
// 枚举 2^(n-1) 种切法，逐段双指针校验回文。
// 时间复杂度 O(n*2^n)，空间复杂度 O(n)（不计结果）。
func partitionEnumeration(s string) [][]string {
	n := len(s)
	res := [][]string{}
	if n == 0 {
		return res
	}
	for mask := 0; mask < 1<<(n-1); mask++ {
		parts := []string{}
		start, ok := 0, true
		for i := 0; i < n && ok; i++ {
			if i == n-1 || mask>>i&1 == 1 {
				if !isPalindrome(s, start, i) {
					ok = false
				}
				parts = append(parts, s[start:i+1])
				start = i + 1
			}
		}
		if ok {
			res = append(res, parts)
		}
	}
	return res
}
