package coinchange2

// 518. 零钱兑换 II
// https://leetcode.cn/problems/coin-change-ii/
// 给一个整数数组 coins 表示不同面额的硬币，和一个整数 amount 表示总金额，
// 返回凑成总金额的硬币组合数（每种硬币数量无限），凑不出返回 0。
//
// 核心思路：完全背包求组合数。物品是每种硬币（可无限次使用），容量是 amount。
// dp[j] 表示凑出金额 j 的组合数，dp[0] = 1，转移：dp[j] += dp[j-coin]。
// 循环顺序很关键：外层遍历物品、内层遍历容量 => 组合数（每种组合只按硬币顺序计一次）；
// 若外层容量、内层物品则得到排列数（如 377 题）。
// 完全背包内层容量「正序」遍历，使同一硬币可以被重复选择。

// change 解法一：完全背包 + 一维滚动数组 ⭐ 最优方案
// 时间复杂度 O(n * amount)，空间复杂度 O(amount)。
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1
	for _, coin := range coins { // 外层物品：求组合数
		for j := coin; j <= amount; j++ { // 正序：硬币可重复使用
			dp[j] += dp[j-coin]
		}
	}
	return dp[amount]
}

// change2D 解法二：完全背包 + 二维 dp 表（对照用）
// dp[i][j] 表示只用前 i 种硬币凑出 j 的组合数：
// dp[i][j] = dp[i-1][j]（不用第 i 种）+ dp[i][j-coin]（至少再用一枚第 i 种，需 j >= coin）。
// 注意「选」的那一项来自同一行 dp[i]，这正是完全背包与 01 背包的区别。
// 时间复杂度 O(n * amount)，空间复杂度 O(n * amount)。
func change2D(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, amount+1)
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		coin := coins[i-1]
		for j := 1; j <= amount; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= coin {
				dp[i][j] += dp[i][j-coin]
			}
		}
	}
	return dp[n][amount]
}
