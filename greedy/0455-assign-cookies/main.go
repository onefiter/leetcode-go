package assigncookies

import "slices"

// 455. 分发饼干
// https://leetcode.cn/problems/assign-cookies/
// 每个孩子 i 有胃口值 g[i]，每块饼干 j 有尺寸 s[j]，当 s[j] >= g[i] 时饼干可以满足孩子，
// 每个孩子最多分一块饼干，求最多能满足多少个孩子。
//
// 核心思路：两个数组都升序排序，用最小的、能满足当前胃口最小孩子的饼干去喂他。
// 正确性（交换论证）：若某最优解给胃口最小的孩子分了更大的饼干，把它换成「刚好够用的最小饼干」不会让答案变差，
// 而省下来的大饼干还能留给胃口更大的孩子；所以「小饼干优先喂小胃口」总能达到最优。

// findContentChildren 解法一：排序 + 双指针贪心 ⭐ 最优方案
// 时间复杂度 O(n log n + m log m)，空间复杂度 O(log n + log m)（排序栈）。
func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)
	child := 0
	for j := 0; j < len(s) && child < len(g); j++ {
		if s[j] >= g[child] {
			child++ // 这块饼干喂饱当前胃口最小的孩子
		}
		// 否则这块饼干太小，谁也喂不饱，直接丢弃
	}
	return child
}

// findContentChildrenBrute 解法二：暴力匹配（对照用）
// 对每个孩子（按胃口升序）在所有未使用的饼干中找能满足他的最小饼干。
// 时间复杂度 O(n * m)，空间复杂度 O(m)。
func findContentChildrenBrute(g []int, s []int) int {
	slices.Sort(g)
	used := make([]bool, len(s))
	count := 0
	for _, need := range g {
		best := -1
		for j, size := range s {
			if !used[j] && size >= need && (best == -1 || size < s[best]) {
				best = j
			}
		}
		if best == -1 {
			break // 胃口更大的孩子更不可能被满足
		}
		used[best] = true
		count++
	}
	return count
}
