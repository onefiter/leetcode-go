package maxprofit3

// 123. 买卖股票的最佳时机 III
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
// 给一个数组 prices，prices[i] 是股票第 i 天的价格。最多可以完成两笔交易，
// 且不能同时参与多笔交易（再次买入前必须卖出），返回能获得的最大利润。
//
// 核心思路：每天结束时只可能处于 4 个状态之一（外加「什么都没做」利润恒为 0）：
// buy1 = 第一次买入后的最大利润，sell1 = 第一次卖出后，buy2 = 第二次买入后，sell2 = 第二次卖出后。
// 转移：buy1 = max(buy1, -p)，sell1 = max(sell1, buy1+p)，
//       buy2 = max(buy2, sell1-p)，sell2 = max(sell2, buy2+p)。
// 同一天内按顺序用新值更新也正确：相当于同一天买入再卖出，利润为 0，不会产生错误答案。
// 答案为 sell2（最多两笔，少做一笔的情况已被「同天买卖」覆盖）。

// maxProfit 解法一：四状态滚动 DP ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func maxProfit(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for _, p := range prices[1:] {
		buy1 = max(buy1, -p)
		sell1 = max(sell1, buy1+p)
		buy2 = max(buy2, sell1-p)
		sell2 = max(sell2, buy2+p)
	}
	return sell2
}

// maxProfitSplit 解法二：前后缀分割（对照用）
// 两笔交易一定可以用某个分割点 i 分开：第一笔在 [0, i] 内，第二笔在 [i, n-1] 内。
// left[i] = 前 i 天内一笔交易的最大利润（正向维护最低价），
// right[i] = 从第 i 天起一笔交易的最大利润（反向维护最高价），答案为 max(left[i] + right[i])。
// 时间复杂度 O(n)，空间复杂度 O(n)。
func maxProfitSplit(prices []int) int {
	n := len(prices)
	left := make([]int, n)
	minPrice := prices[0]
	for i := 1; i < n; i++ {
		minPrice = min(minPrice, prices[i])
		left[i] = max(left[i-1], prices[i]-minPrice)
	}
	right := make([]int, n)
	maxPrice := prices[n-1]
	for i := n - 2; i >= 0; i-- {
		maxPrice = max(maxPrice, prices[i])
		right[i] = max(right[i+1], maxPrice-prices[i])
	}
	best := 0
	for i := 0; i < n; i++ {
		best = max(best, left[i]+right[i])
	}
	return best
}
