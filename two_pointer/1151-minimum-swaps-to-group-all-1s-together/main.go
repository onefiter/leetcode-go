package minswaps

// 1151. 最少交换次数来组合所有的 1
// https://leetcode.cn/problems/minimum-swaps-to-group-all-1s-together/
// 给出一个二进制数组 data，每次可以交换任意两个位置的元素，
// 返回将数组中所有 1 组合到一起（连续排列）所需的最少交换次数。
// 例如 data=[1,0,1,0,1] -> 1（交换后得到 [0,0,1,1,1] 之类）。
//
// 一句话理解题意：设 1 的总数为 m，最终所有 1 会挤在某个长度为 m 的窗口里，
// 窗口里原本的每个 0 都要和窗口外的某个 1 交换一次。
//
// 核心思路：交换次数 = m - 窗口内 1 的个数。用定长为 m 的滑动窗口找「1 最多」的窗口。

// minSwaps 解法一：定长滑动窗口 ⭐ 最优方案
// 先统计 1 的总数 m，再在长度为 m 的窗口上滑动，维护窗口内 1 的个数 cur，记录最大值 best。
// 答案为 m - best；m 为 0 时不需要交换。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func minSwaps(data []int) int {
	m := 0
	for _, x := range data {
		m += x
	}
	if m == 0 {
		return 0
	}
	cur, best := 0, 0
	for i, x := range data {
		cur += x
		if i >= m {
			cur -= data[i-m] // 移出窗口左侧元素
		}
		if i >= m-1 {
			best = max(best, cur)
		}
	}
	return m - best
}

// minSwapsEnumeration 解法二：枚举每个长度为 m 的窗口重新计数（对照用）
// 对每个起点 s，直接数 data[s..s+m-1] 中 0 的个数，即该窗口需要的交换次数，取最小值。
// 时间复杂度 O(n·m)，空间复杂度 O(1)。
func minSwapsEnumeration(data []int) int {
	m := 0
	for _, x := range data {
		m += x
	}
	if m == 0 {
		return 0
	}
	ans := m
	for s := 0; s+m <= len(data); s++ {
		zeros := 0
		for j := s; j < s+m; j++ {
			if data[j] == 0 {
				zeros++
			}
		}
		ans = min(ans, zeros)
	}
	return ans
}
