package coinchange

// 322. 零钱兑换
// https://leetcode.cn/problems/coin-change/
// 给不同面额的硬币 coins（每种无限个）和总金额 amount，求凑成 amount 所需的最少硬币数，
// 无法凑成返回 -1。
//
// 核心思路：完全背包求最小值。dp[i] 表示凑成金额 i 所需的最少硬币数。
// 最后一枚硬币可以是任意 c（c <= i），转移方程 dp[i] = min(dp[i-c] + 1)；
// 初始 dp[0] = 0，其余设为 amount+1 表示「不可达」（硬币数不可能超过 amount）。
// 最终 dp[amount] > amount 说明凑不出，返回 -1。

// coinChange 解法一：一维完全背包 DP ⭐ 最优方案
// 时间复杂度 O(amount * len(coins))，空间复杂度 O(amount)。
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for _, c := range coins {
		for i := c; i <= amount; i++ { // 正序遍历，每种硬币可重复使用
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// coinChangeBFS 解法二：BFS 求最短路（对照用）
// 把每个金额看作图中节点，从 0 出发每加一枚硬币走一步，第一次到达 amount 的层数就是最少硬币数。
// 时间复杂度 O(amount * len(coins))，空间复杂度 O(amount)。
func coinChangeBFS(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	visited := make([]bool, amount+1)
	visited[0] = true
	queue := []int{0}
	for step := 1; len(queue) > 0; step++ {
		next := []int{}
		for _, cur := range queue {
			for _, c := range coins {
				v := cur + c
				if v == amount {
					return step
				}
				if v < amount && !visited[v] {
					visited[v] = true
					next = append(next, v)
				}
			}
		}
		queue = next
	}
	return -1
}
