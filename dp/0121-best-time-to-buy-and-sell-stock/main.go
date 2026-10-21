package maxprofit

// 121. 买卖股票的最佳时机
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/
// 给一个数组 prices，prices[i] 是股票第 i 天的价格。只能选某一天买入、之后某一天卖出（最多一次交易），
// 返回能获得的最大利润；无法获利则返回 0。
//
// 核心思路：dp[i][0] 表示第 i 天结束时「不持股」的最大利润，dp[i][1] 表示「持股」的最大利润。
// 转移：dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])（今天卖出）；
//       dp[i][1] = max(dp[i-1][1], -prices[i])（只能买一次，所以买入前利润一定是 0）。
// 等价于：遍历时维护历史最低价 minPrice，今天卖出的利润就是 prices[i]-minPrice。

// maxProfit 解法一：滚动 DP（维护历史最低价） ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func maxProfit(prices []int) int {
	minPrice, best := prices[0], 0
	for _, p := range prices[1:] {
		best = max(best, p-minPrice) // 今天卖出
		minPrice = min(minPrice, p)  // 更新最佳买入点
	}
	return best
}

// maxProfitBrute 解法二：暴力枚举买卖日（对照用）
// 枚举所有 i < j 的买入日 i 与卖出日 j，取 prices[j]-prices[i] 的最大值。
// 时间复杂度 O(n^2)，空间复杂度 O(1)。
func maxProfitBrute(prices []int) int {
	best := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			best = max(best, prices[j]-prices[i])
		}
	}
	return best
}
