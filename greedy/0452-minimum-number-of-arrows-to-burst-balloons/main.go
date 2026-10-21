package minarrows

import (
	"math"
	"slices"
)

// 452. 用最少数量的箭引爆气球
// https://leetcode.cn/problems/minimum-number-of-arrows-to-burst-balloons/
// 每个气球用水平区间 points[i] = [xstart, xend] 表示，在 x 处竖直射出的箭
// 能引爆所有满足 xstart <= x <= xend 的气球（端点相接也会被同一支箭引爆），求最少箭数。
//
// 核心思路：按右端点升序排序，第一支箭射在第一个气球的右端点上——
// 这是能引爆它的最靠右位置，能顺带引爆尽可能多的后续气球。
// 之后遇到左端点 > 当前箭位置的气球，说明它没被引爆，再在它的右端点补一支箭。
// 正确性同活动选择：箭数 = 最多互不相交的气球数（闭区间意义下）。
// 注意坐标范围是 [-2^31, 2^31-1]，排序比较必须用比较运算而不能用 a-b，否则会溢出。

// cmpInt 用比较而非减法，避免 int32 极端值相减溢出。
func cmpInt(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

// findMinArrowShots 解法一：按右端点排序 + 贪心射在右端点 ⭐ 最优方案
// 时间复杂度 O(n log n)，空间复杂度 O(log n)（排序栈）。
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	slices.SortFunc(points, func(a, b []int) int { return cmpInt(a[1], b[1]) })
	arrows, pos := 1, points[0][1]
	for _, p := range points[1:] {
		if p[0] > pos { // 当前箭够不到，补一支射在它的右端点
			arrows++
			pos = p[1]
		}
	}
	return arrows
}

// findMinArrowShotsByStart 解法二：按左端点排序 + 维护重叠区间交集（对照用）
// 按左端点扫描，维护当前一组气球的公共交集右界 minEnd；
// 新气球左端点 <= minEnd 则可共用一支箭并收缩交集，否则开新组。
// 时间复杂度 O(n log n)，空间复杂度 O(log n)。
func findMinArrowShotsByStart(points [][]int) int {
	if len(points) == 0 {
		return 0
	}
	slices.SortFunc(points, func(a, b []int) int { return cmpInt(a[0], b[0]) })
	arrows, minEnd := 0, math.MinInt
	for i, p := range points {
		if i == 0 || p[0] > minEnd { // 与当前组无交集，开新的一支箭
			arrows++
			minEnd = p[1]
		} else {
			minEnd = min(minEnd, p[1]) // 收缩交集
		}
	}
	return arrows
}
