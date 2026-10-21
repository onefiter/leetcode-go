package koko

import "slices"

// 875. 爱吃香蕉的珂珂
// https://leetcode.cn/problems/koko-eating-bananas/
// 有 n 堆香蕉 piles[i]，警卫 h 小时后回来。珂珂每小时选一堆吃 k 根，
// 不足 k 根就吃完这堆、这小时不再吃别的。求能在 h 小时内吃完的最小速度 k。
// 例如 piles=[3,6,7,11], h=8 -> 4。
//
// 一句话理解题意：速度越快用时越少，找「刚好来得及」的最慢速度。
//
// 核心思路：用时 hours(k) = Σ ceil(p/k) 随 k 单调不增，
// 「hours(k) <= h」在 [1, max(piles)] 上是后缀为真，二分答案找第一个满足的 k。
// 题目保证 len(piles) <= h，所以速度 max(piles) 一定可行。

// hoursNeeded 返回速度 k 下吃完所有香蕉需要的小时数。
func hoursNeeded(piles []int, k int) int {
	total := 0
	for _, p := range piles {
		total += (p + k - 1) / k // 向上取整
	}
	return total
}

// minEatingSpeed 解法一：二分答案 ⭐ 最优方案
// 在 [1, max(piles)] 上二分速度 mid：来得及就尝试更慢（hi = mid），否则加速（lo = mid+1）。
// 时间复杂度 O(n log M)，M 为最大堆大小；空间复杂度 O(1)。
func minEatingSpeed(piles []int, h int) int {
	lo, hi := 1, slices.Max(piles)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if hoursNeeded(piles, mid) <= h {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// minEatingSpeedEnumeration 解法二：从 1 开始线性枚举速度（对照用）
// 第一个来得及的速度就是答案。
// 时间复杂度 O(n · M)，空间复杂度 O(1)。
func minEatingSpeedEnumeration(piles []int, h int) int {
	maxPile := slices.Max(piles)
	for k := 1; k < maxPile; k++ {
		if hoursNeeded(piles, k) <= h {
			return k
		}
	}
	return maxPile
}
