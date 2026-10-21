package jumpgame2

// 45. 跳跃游戏 II
// https://leetcode.cn/problems/jump-game-ii/
// 给一个非负整数数组 nums，初始位于下标 0，nums[i] 表示从 i 最多能向后跳的长度，
// 返回到达最后一个下标的最少跳跃次数（题目保证一定能到达）。
//
// 核心思路：把跳跃过程看成「按层 BFS」——当前跳数能覆盖的区间是 [start, end]，
// 遍历这一层时维护下一跳能到达的最远位置 farthest；走到 end 时必须再跳一次，
// 下一层的边界就是 farthest。每一跳都选「能让下一层覆盖最远」的位置，
// 因为覆盖更远的区间包含覆盖更近区间的所有选择，所以不会更差，贪心正确。

// jump 解法一：贪心（按层扩展最远边界）⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func jump(nums []int) int {
	steps, end, farthest := 0, 0, 0
	// 只遍历到 n-2：到达最后一个下标后无需再跳
	for i := 0; i < len(nums)-1; i++ {
		farthest = max(farthest, i+nums[i])
		if i == end { // 当前层走完，必须再跳一次
			steps++
			end = farthest
			if end >= len(nums)-1 {
				break
			}
		}
	}
	return steps
}

// jumpDP 解法二：动态规划（对照用）
// dp[i] 表示到达下标 i 的最少跳数，从每个 i 向后更新它能到达的所有位置。
// 时间复杂度 O(n * max(nums))，最坏 O(n^2)，空间复杂度 O(n)。
func jumpDP(nums []int) int {
	n := len(nums)
	const inf = 1 << 60
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = inf
	}
	for i := 0; i < n; i++ {
		if dp[i] == inf {
			continue
		}
		for j := i + 1; j <= i+nums[i] && j < n; j++ {
			dp[j] = min(dp[j], dp[i]+1)
		}
	}
	return dp[n-1]
}
