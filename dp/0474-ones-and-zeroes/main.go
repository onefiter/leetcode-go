package onesandzeroes

import "strings"

// 474. 一和零
// https://leetcode.cn/problems/ones-and-zeroes/
// 给一个二进制字符串数组 strs 和两个整数 m、n，找出 strs 的最大子集的长度，
// 该子集中最多有 m 个 0 和 n 个 1。
//
// 核心思路：二维费用的 01 背包。物品是每个字符串（只能选一次），价值为 1，
// 有两种容量：0 的个数上限 m、1 的个数上限 n。
// dp[i][j] 表示最多 i 个 0、j 个 1 时能选的最多字符串数，
// 转移：dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones] + 1)。
// 两个容量维度都要「倒序」遍历，保证每个字符串只被选一次。

// findMaxForm 解法一：二维费用 01 背包 + 滚动数组 ⭐ 最优方案
// 时间复杂度 O(L*m*n + 总字符数)，空间复杂度 O(m*n)，L 为字符串个数。
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for _, s := range strs {
		zeros := strings.Count(s, "0")
		ones := len(s) - zeros
		for i := m; i >= zeros; i-- { // 倒序：每个字符串只用一次
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], dp[i-zeros][j-ones]+1)
			}
		}
	}
	return dp[m][n]
}

// findMaxFormMemo 解法二：记忆化 DFS（对照用）
// dfs(k, i, j) 表示从第 k 个字符串开始、剩余 i 个 0 和 j 个 1 时能选的最多个数，
// 对每个字符串做「选 / 不选」，结果缓存在三维数组里。
// 时间复杂度 O(L*m*n)，空间复杂度 O(L*m*n)。
func findMaxFormMemo(strs []string, m int, n int) int {
	L := len(strs)
	zeros := make([]int, L)
	for k, s := range strs {
		zeros[k] = strings.Count(s, "0")
	}
	memo := make([][][]int, L)
	for k := range memo {
		memo[k] = make([][]int, m+1)
		for i := range memo[k] {
			memo[k][i] = make([]int, n+1)
			for j := range memo[k][i] {
				memo[k][i][j] = -1
			}
		}
	}
	var dfs func(k, i, j int) int
	dfs = func(k, i, j int) int {
		if k == L {
			return 0
		}
		if memo[k][i][j] >= 0 {
			return memo[k][i][j]
		}
		res := dfs(k+1, i, j) // 不选
		z, o := zeros[k], len(strs[k])-zeros[k]
		if i >= z && j >= o {
			res = max(res, dfs(k+1, i-z, j-o)+1) // 选
		}
		memo[k][i][j] = res
		return res
	}
	return dfs(0, m, n)
}
