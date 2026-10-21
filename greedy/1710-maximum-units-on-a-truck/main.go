package maximumunits

import "slices"

// 1710. 卡车上的最大单元数
// https://leetcode.cn/problems/maximum-units-on-a-truck/
// boxTypes[i] = [numberOfBoxes, numberOfUnitsPerBox] 表示第 i 类箱子的数量和每个箱子装的单元数，
// 卡车最多装 truckSize 个箱子，返回卡车能装载的最大单元总数。
//
// 核心思路：每个箱子占用的「容量」都是 1，只是价值（单元数）不同，
// 所以按每箱单元数降序装，先把价值最高的箱子装满，再装次高的，直到卡车装满。
// 正确性：任何解里若装了低价值箱子而留下高价值箱子，交换二者总价值不减（分数背包的特例）。

// maximumUnits 解法一：按单元数降序排序 + 贪心 ⭐ 最优方案
// 时间复杂度 O(n log n)，空间复杂度 O(log n)（排序栈）。
func maximumUnits(boxTypes [][]int, truckSize int) int {
	slices.SortFunc(boxTypes, func(a, b []int) int { return b[1] - a[1] })
	total := 0
	for _, bt := range boxTypes {
		if truckSize == 0 {
			break
		}
		take := min(bt[0], truckSize)
		total += take * bt[1]
		truckSize -= take
	}
	return total
}

// maximumUnitsBucket 解法二：计数排序 / 桶（对照用）
// 题目约束每箱单元数 <= 1000，按单元数分桶累加箱子数量，再从高到低取，免去比较排序。
// 时间复杂度 O(n + U)，空间复杂度 O(U)，U 为单元数上界。
func maximumUnitsBucket(boxTypes [][]int, truckSize int) int {
	maxUnit := 0
	for _, bt := range boxTypes {
		maxUnit = max(maxUnit, bt[1])
	}
	bucket := make([]int, maxUnit+1) // bucket[u]：每箱 u 个单元的箱子总数
	for _, bt := range boxTypes {
		bucket[bt[1]] += bt[0]
	}
	total := 0
	for u := maxUnit; u > 0 && truckSize > 0; u-- {
		take := min(bucket[u], truckSize)
		total += take * u
		truckSize -= take
	}
	return total
}
