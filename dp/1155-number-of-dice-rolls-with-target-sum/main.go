package dicerolls

// 1155. 掷骰子等于目标和的方法数
// https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/
// 有 n 个一样的骰子，每个有 k 个面（点数 1 到 k），给定整数 target，
// 返回掷出的点数之和等于 target 的方法数（每个骰子按顺序区分），结果对 1e9+7 取模。
//
// 核心思路：分组背包。每个骰子是一组物品（点数 1..k 中恰好选一个），容量是 target。
// dp[j] 表示前 i 个骰子点数和为 j 的方法数，dp[0] = 1，
// 转移：dp[j] = Σ dp_prev[j-x]（1 <= x <= min(k, j)）。
// 每组「恰好选一个」，因此内层容量「倒序」遍历（类似 01 背包），并且先把 dp[j] 清零，
// 倒序时 j 越界到 0 也要清零，表示用了 i 个骰子后和不可能为 0。

const mod = 1_000_000_007

// numRollsToTarget 解法一：分组背包 + 一维滚动数组 ⭐ 最优方案
// 时间复杂度 O(n * target * k)，空间复杂度 O(target)。
func numRollsToTarget(n int, k int, target int) int {
	if target < n || target > n*k {
		return 0
	}
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= n; i++ {
		for j := target; j >= 0; j-- { // 倒序：dp[j-x] 仍是上一个骰子的结果
			dp[j] = 0
			for x := 1; x <= k && x <= j; x++ {
				dp[j] = (dp[j] + dp[j-x]) % mod
			}
		}
	}
	return dp[target]
}

// numRollsToTargetMemo 解法二：记忆化 DFS（对照用）
// dfs(i, remain) 表示还剩 i 个骰子、需要凑出 remain 的方法数，
// 枚举当前骰子点数 x，递归 dfs(i-1, remain-x)，并用剪枝排除 remain 超出 [i, i*k] 的状态。
// 时间复杂度 O(n * target * k)，空间复杂度 O(n * target)。
func numRollsToTargetMemo(n int, k int, target int) int {
	memo := make([][]int, n+1)
	for i := range memo {
		memo[i] = make([]int, target+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}
	var dfs func(i, remain int) int
	dfs = func(i, remain int) int {
		if remain < i || remain > i*k {
			return 0 // 剩余点数不可能凑出
		}
		if i == 0 {
			return 1 // 此时 remain 必为 0
		}
		if memo[i][remain] >= 0 {
			return memo[i][remain]
		}
		res := 0
		for x := 1; x <= k && x <= remain; x++ {
			res = (res + dfs(i-1, remain-x)) % mod
		}
		memo[i][remain] = res
		return res
	}
	return dfs(n, target)
}
