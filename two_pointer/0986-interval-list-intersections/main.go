package intervalintersection

// 986. 区间列表的交集
// https://leetcode.cn/problems/interval-list-intersections/
// 给定两个由闭区间组成的列表 firstList 和 secondList，每个列表内的区间两两不相交且已排序。
// 返回这两个列表的交集。例如 [[0,2],[5,10],[13,23],[24,25]] 与 [[1,5],[8,12],[15,24],[25,26]]
// -> [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]。
//
// 核心思路：双指针各指向一个区间，交集为 [max(起点), min(终点)]，非空就记录；
// 然后丢掉「结束得更早」的那个区间（它不可能再和后面的区间相交）。

// intervalIntersection 解法一：双指针归并 ⭐ 最优方案
// 时间复杂度 O(m+n)，空间复杂度 O(1)（不计返回结果）。
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	i, j := 0, 0
	for i < len(firstList) && j < len(secondList) {
		lo := max(firstList[i][0], secondList[j][0])
		hi := min(firstList[i][1], secondList[j][1])
		if lo <= hi {
			res = append(res, []int{lo, hi})
		}
		if firstList[i][1] < secondList[j][1] {
			i++
		} else {
			j++
		}
	}
	return res
}

// intervalIntersectionEnumeration 解法二：两两枚举（对照用）
// 枚举所有区间对，计算交集；因为两个列表都有序且各自不相交，
// 按 (i, j) 字典序产生的交集天然按起点有序。
// 时间复杂度 O(m·n)，空间复杂度 O(1)（不计返回结果）。
func intervalIntersectionEnumeration(firstList [][]int, secondList [][]int) [][]int {
	res := [][]int{}
	for _, a := range firstList {
		for _, b := range secondList {
			lo, hi := max(a[0], b[0]), min(a[1], b[1])
			if lo <= hi {
				res = append(res, []int{lo, hi})
			}
		}
	}
	return res
}
