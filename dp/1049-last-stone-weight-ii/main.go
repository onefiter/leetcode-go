package laststone2

// 1049. 最后一块石头的重量 II
// https://leetcode.cn/problems/last-stone-weight-ii/
// 有一堆石头 stones，每次任选两块粉碎：重量相等则都消失，否则剩下差值。
// 最后至多剩一块石头，返回其可能的最小重量（没有剩下则返回 0）。
//
// 核心思路：每块石头最终带 '+' 或 '-' 号参与运算，相当于把石头分成两堆，答案是两堆差值的最小值。
// 设总和 sum，让一堆尽量接近 sum/2：转化为 01 背包求最大重量。
// 物品是每块石头（只能选一次），重量与价值都为 stones[i]，容量是 target = sum/2。
// dp[j] 表示容量 j 下能装的最大重量，转移：dp[j] = max(dp[j], dp[j-w] + w)。
// 01 背包内层容量「倒序」遍历。答案 = sum - 2*dp[target]。

// lastStoneWeightII 解法一：01 背包 + 一维滚动数组 ⭐ 最优方案
// 时间复杂度 O(n * sum)，空间复杂度 O(sum)。
func lastStoneWeightII(stones []int) int {
	sum := 0
	for _, w := range stones {
		sum += w
	}
	target := sum / 2
	dp := make([]int, target+1)
	for _, w := range stones {
		for j := target; j >= w; j-- { // 倒序：每块石头只用一次
			dp[j] = max(dp[j], dp[j-w]+w)
		}
	}
	return sum - 2*dp[target]
}

// lastStoneWeightIISet 解法二：可达和集合枚举（对照用）
// 维护所有可达的「带符号和」集合，每块石头把集合里的每个值分别 +w 和 -w，
// 最后取集合中绝对值最小的非负结果。
// 时间复杂度 O(n * sum)，空间复杂度 O(sum)。
func lastStoneWeightIISet(stones []int) int {
	reach := map[int]bool{0: true}
	for _, w := range stones {
		next := make(map[int]bool, len(reach)*2)
		for s := range reach {
			next[s+w] = true
			next[s-w] = true
		}
		reach = next
	}
	best := -1
	for s := range reach {
		if s >= 0 && (best < 0 || s < best) {
			best = s
		}
	}
	return best
}
