package triangle

// 120. 三角形最小路径和
// https://leetcode.cn/problems/triangle/
// 给定一个三角形 triangle，找出自顶向下的最小路径和。
// 每一步只能移动到下一行中相邻的结点上：位于第 i 行下标 j 时，下一步可移动到下一行的 j 或 j+1。
//
// 核心思路：自底向上 DP。dp[j] 表示从第 i 行第 j 个结点出发走到底部的最小路径和。
// 初始为最后一行本身，转移：dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]。
// 行从倒数第二行往上遍历，行内 j 从左到右（dp[j] 只依赖下一行的 dp[j]、dp[j+1]，
// 而 dp[j+1] 此时尚未被本行覆盖）。最终 dp[0] 即为答案，无需处理边界。

// minimumTotal 解法一：自底向上 + 一维数组 ⭐ 最优方案
// 时间复杂度 O(n^2)，空间复杂度 O(n)（n 为行数，不修改原数组）。
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)
	copy(dp, triangle[n-1])
	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}
	return dp[0]
}

// minimumTotalTopDown 解法二：自顶向下记忆化递归（对照用）
// dfs(i, j) 表示从 (i, j) 出发到底部的最小路径和：
// dfs(i, j) = triangle[i][j] + min(dfs(i+1, j), dfs(i+1, j+1))，最后一行直接返回自身。
// 用 memo 记录已算过的状态，避免指数级重复计算。
// 时间复杂度 O(n^2)，空间复杂度 O(n^2)（memo + 递归栈 O(n)）。
func minimumTotalTopDown(triangle [][]int) int {
	n := len(triangle)
	memo := make([][]int, n)
	done := make([][]bool, n)
	for i := range memo {
		memo[i] = make([]int, i+1)
		done[i] = make([]bool, i+1)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if i == n-1 {
			return triangle[i][j]
		}
		if done[i][j] {
			return memo[i][j]
		}
		res := triangle[i][j] + min(dfs(i+1, j), dfs(i+1, j+1))
		memo[i][j], done[i][j] = res, true
		return res
	}
	return dfs(0, 0)
}
