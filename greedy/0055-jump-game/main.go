package jumpgame

// 55. 跳跃游戏
// https://leetcode.cn/problems/jump-game/
// 给一个非负整数数组 nums，初始位于下标 0，nums[i] 表示从 i 最多能向后跳的长度，
// 判断能否到达最后一个下标。
//
// 核心思路：维护当前能到达的最远位置 reach。从左到右扫描，若 i > reach 说明 i 不可达，
// 后面更不可达，直接返回 false；否则用 i+nums[i] 更新 reach。
// 正确性：可达位置一定是连续前缀 [0, reach]（能跳到 x 就能跳到 x 之前的任意位置），
// 所以只需记录这个前缀的右端点即可，不必关心具体怎么跳。

// canJump 解法一：贪心维护最远可达位置 ⭐ 最优方案
// 时间复杂度 O(n)，空间复杂度 O(1)。
func canJump(nums []int) bool {
	reach := 0
	for i, v := range nums {
		if i > reach {
			return false // 出现断点，i 无法到达
		}
		reach = max(reach, i+v)
		if reach >= len(nums)-1 {
			return true
		}
	}
	return true
}

// canJumpDP 解法二：从后往前动态规划（对照用）
// good[i] 表示从 i 出发能否到达终点：只要 [i+1, i+nums[i]] 中存在 good 的位置即可。
// 时间复杂度 O(n * max(nums))，最坏 O(n^2)，空间复杂度 O(n)。
func canJumpDP(nums []int) bool {
	n := len(nums)
	good := make([]bool, n)
	good[n-1] = true
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j <= i+nums[i] && j < n; j++ {
			if good[j] {
				good[i] = true
				break
			}
		}
	}
	return good[0]
}
