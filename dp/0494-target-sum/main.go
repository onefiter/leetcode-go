package targetsum

// 494. 目标和
// https://leetcode.cn/problems/target-sum/
// 给一个非负整数数组 nums 和整数 target，在每个数前添加 '+' 或 '-' 串联成表达式，
// 返回运算结果等于 target 的不同表达式的数目。
//
// 核心思路：设取正号的数之和为 P，取负号的为 N，则 P - N = target，P + N = sum，
// 得 P = (sum + target) / 2。问题转化为 01 背包求方案数：从 nums 中选若干个数恰好凑出 P。
// 若 sum + target 为奇数或 |target| > sum 则无解。
// 物品是每个数字（只能选一次），容量是 P。
// dp[j] 表示凑出和 j 的方案数，dp[0] = 1，转移：dp[j] += dp[j-num]。
// 01 背包内层容量「倒序」遍历。注意 num 可以为 0，此时方案数翻倍，倒序写法天然处理。

// findTargetSumWays 解法一：01 背包求方案数 + 一维滚动数组 ⭐ 最优方案
// 时间复杂度 O(n * P)，空间复杂度 O(P)。
func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, x := range nums {
		sum += x
	}
	if target > sum || -target > sum || (sum+target)%2 != 0 {
		return 0
	}
	p := (sum + target) / 2
	dp := make([]int, p+1)
	dp[0] = 1
	for _, num := range nums {
		for j := p; j >= num; j-- { // 倒序：每个数只用一次
			dp[j] += dp[j-num]
		}
	}
	return dp[p]
}

// findTargetSumWaysMemo 解法二：记忆化 DFS（对照用）
// dfs(i, cur) 表示处理到第 i 个数、当前和为 cur 时，后续能凑到 target 的方案数，
// 每个数分别尝试 '+' 和 '-'，用哈希表按 (i, cur) 缓存。
// 时间复杂度 O(n * sum)，空间复杂度 O(n * sum)。
func findTargetSumWaysMemo(nums []int, target int) int {
	type key struct{ i, cur int }
	memo := map[key]int{}
	var dfs func(i, cur int) int
	dfs = func(i, cur int) int {
		if i == len(nums) {
			if cur == target {
				return 1
			}
			return 0
		}
		k := key{i, cur}
		if v, ok := memo[k]; ok {
			return v
		}
		res := dfs(i+1, cur+nums[i]) + dfs(i+1, cur-nums[i])
		memo[k] = res
		return res
	}
	return dfs(0, 0)
}
