package maxprofit2

// 122. 买卖股票的最佳时机 II
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
// 给一个数组 prices，prices[i] 是股票第 i 天的价格。每天可以买入或卖出，但任何时候最多持有一股，
// 可以进行任意多次交易（也可同一天卖出再买入），返回能获得的最大利润。
//
// 核心思路：dp[i][0] 表示第 i 天结束时「不持股」的最大利润，dp[i][1] 表示「持股」的最大利润。
// 转移：dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])（今天卖出）；
//       dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])（今天买入，交易次数不限，可以基于之前的利润）。
// 初始 dp[0][0] = 0, dp[0][1] = -prices[0]，答案 dp[n-1][0]。每天只依赖前一天，可压缩成两个变量。

// maxProfit 解法一：滚动状态机 DP ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func maxProfit(prices []int) int {
	hold, free := -prices[0], 0
	for _, p := range prices[1:] {
		// 用旧值同时更新，避免同一天内状态互相污染
		hold, free = max(hold, free-p), max(free, hold+p)
	}
	return free
}

// maxProfitTable 解法二：完整 dp 表（对照用）
// 显式维护 n×2 的状态表，与状态定义一一对应，便于理解后再做空间压缩。
// （另一种思路是贪心：累加所有 prices[i]-prices[i-1] > 0 的上涨段，结果相同。）
// 时间复杂度 O(n)，空间复杂度 O(n)。
func maxProfitTable(prices []int) int {
	n := len(prices)
	dp := make([][2]int, n)
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
	}
	return dp[n-1][0]
}
