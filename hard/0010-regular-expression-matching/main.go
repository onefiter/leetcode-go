package regex

// 10. 正则表达式匹配
// https://leetcode.cn/problems/regular-expression-matching/
// 给字符串 s 和字符规律 p，实现支持 '.' 和 '*' 的正则匹配：
// '.' 匹配任意单个字符，'*' 匹配零个或多个前面那个元素，要求覆盖整个 s。
//
// 核心思路：二维 DP。dp[i][j] 表示 s 的前 i 个字符能否被 p 的前 j 个字符匹配。
//   - p[j-1] 是普通字符或 '.'：dp[i][j] = dp[i-1][j-1] && match(s[i-1], p[j-1])；
//   - p[j-1] 是 '*'：把「x*」看成整体，
//     匹配 0 次：dp[i][j-2]；
//     匹配 ≥1 次：match(s[i-1], p[j-2]) && dp[i-1][j]（s 吃掉一个字符，x* 还能继续用）。
// 初始 dp[0][0] = true，空串也能被 "a*b*" 这类模式匹配，所以 i 要从 0 开始算。

// match 判断单个字符 c 是否能被模式字符 pc 匹配。
func match(c, pc byte) bool {
	return pc == '.' || c == pc
}

// isMatch 解法一：自底向上二维 DP ⭐ 最优方案
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)。
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = j >= 2 && dp[i][j-2] // x* 匹配 0 次
				if !dp[i][j] && i > 0 && j >= 2 && match(s[i-1], p[j-2]) {
					dp[i][j] = dp[i-1][j] // x* 再多匹配一个字符
				}
			} else if i > 0 && match(s[i-1], p[j-1]) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[m][n]
}

// isMatchMemo 解法二：记忆化递归（对照用）
// dfs(i, j) 表示 s[i:] 能否被 p[j:] 匹配：
//   - 若 p[j+1] == '*'：跳过 x*（dfs(i, j+2)），或首字符匹配时 dfs(i+1, j)；
//   - 否则首字符匹配且 dfs(i+1, j+1)。
//
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)。
func isMatchMemo(s string, p string) bool {
	m, n := len(s), len(p)
	memo := make([][]int8, m+1) // 0 未计算，1 true，-1 false
	for i := range memo {
		memo[i] = make([]int8, n+1)
	}
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		if j == n {
			return i == m
		}
		if memo[i][j] != 0 {
			return memo[i][j] == 1
		}
		first := i < m && match(s[i], p[j])
		var res bool
		if j+1 < n && p[j+1] == '*' {
			res = dfs(i, j+2) || (first && dfs(i+1, j))
		} else {
			res = first && dfs(i+1, j+1)
		}
		if res {
			memo[i][j] = 1
		} else {
			memo[i][j] = -1
		}
		return res
	}
	return dfs(0, 0)
}
