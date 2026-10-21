package editdistance

// 72. 编辑距离
// https://leetcode.cn/problems/edit-distance/
// 给两个单词 word1 和 word2，返回将 word1 转换成 word2 所使用的最少操作数。
// 可以对一个单词进行三种操作：插入一个字符、删除一个字符、替换一个字符。
//
// 核心思路：dp[i][j] 表示 word1 前 i 个字符转换成 word2 前 j 个字符的最少操作数。
// 边界：dp[i][0] = i（全删），dp[0][j] = j（全插）。
// 转移：若 word1[i-1] == word2[j-1]，dp[i][j] = dp[i-1][j-1]（无需操作）；
// 否则 dp[i][j] = 1 + min(dp[i-1][j]（删）, dp[i][j-1]（插）, dp[i-1][j-1]（替换）)。
// 遍历顺序：i、j 均正序，依赖左、上、左上三个格子。

// minDistance 解法一：二维 DP 压缩为一维滚动数组 ⭐ 最优方案
// 用 prev 暂存左上角 dp[i-1][j-1]，避免被覆盖。
// 时间复杂度 O(m * n)，空间复杂度 O(n)。
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	dp := make([]int, n+1)
	for j := 0; j <= n; j++ {
		dp[j] = j
	}
	for i := 1; i <= m; i++ {
		prev := dp[0] // 左上角 dp[i-1][0]
		dp[0] = i
		for j := 1; j <= n; j++ {
			tmp := dp[j] // 本轮覆盖前是 dp[i-1][j]，下一列的左上角
			if word1[i-1] == word2[j-1] {
				dp[j] = prev
			} else {
				dp[j] = 1 + min(dp[j], dp[j-1], prev)
			}
			prev = tmp
		}
	}
	return dp[n]
}

// minDistanceMemo 解法二：记忆化递归（对照用）
// dfs(i, j) 表示 word1[i:] 转换成 word2[j:] 的最少操作数，自顶向下求解。
// 时间复杂度 O(m * n)，空间复杂度 O(m * n)。
func minDistanceMemo(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	memo := make([][]int, m+1)
	for i := range memo {
		memo[i] = make([]int, n+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == m {
			return n - j // 剩余全部插入
		}
		if j == n {
			return m - i // 剩余全部删除
		}
		if memo[i][j] != -1 {
			return memo[i][j]
		}
		var res int
		if word1[i] == word2[j] {
			res = dfs(i+1, j+1)
		} else {
			res = 1 + min(dfs(i+1, j), dfs(i, j+1), dfs(i+1, j+1))
		}
		memo[i][j] = res
		return res
	}
	return dfs(0, 0)
}
