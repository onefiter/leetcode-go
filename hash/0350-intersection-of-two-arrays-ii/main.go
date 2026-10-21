package intersectionii

import (
	"slices"
	"sort"
)

// 350. 两个数组的交集 II
// https://leetcode.cn/problems/intersection-of-two-arrays-ii/
// 给定两个数组 nums1 和 nums2，返回它们的交集。输出中每个元素出现的次数应与
// 它在「两个数组中都出现」的次数一致（取较小者）。顺序不限。
// 例如 nums1=[1,2,2,1], nums2=[2,2] -> [2,2]。
//
// 核心思路：统计较短数组每个值的出现次数，遍历另一个数组时若计数 >0 则收集并减一，
// 这样自然得到「两边出现次数的最小值」。

// intersect 【解法1】哈希计数 ⭐ 最优
// 为节省空间，对较短的数组建计数表。
// 时间 O(n+m)：分别遍历两个数组。
// 空间 O(min(n, m))：对较短数组计数。
func intersect(nums1, nums2 []int) []int {
	// 保证 nums1 为较短的，减少哈希表大小。
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	cnt := make(map[int]int, len(nums1))
	for _, x := range nums1 {
		cnt[x]++
	}
	res := []int{}
	for _, x := range nums2 {
		if cnt[x] > 0 {
			res = append(res, x)
			cnt[x]--
		}
	}
	return res
}

// intersectSort 【解法2】排序 + 双指针
// 两数组排序后双指针同步前进：相等收集并同时前进，否则较小一侧前进。
// 适合两个数组已排序或内存受限（无需哈希表）的场景。
// 时间 O(n*logn + m*logm)：排序开销。
// 空间 O(n+m)：克隆避免修改入参。
func intersectSort(nums1, nums2 []int) []int {
	a := slices.Clone(nums1)
	b := slices.Clone(nums2)
	sort.Ints(a)
	sort.Ints(b)
	res := []int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		switch {
		case a[i] < b[j]:
			i++
		case a[i] > b[j]:
			j++
		default:
			res = append(res, a[i])
			i++
			j++
		}
	}
	return res
}

// intersectEnumeration 【解法3】Enumeration 暴力（对照用）
// 用布尔标记数组记录 nums2 中已被匹配的位置，对 nums1 每个值线性找一个未用的相等项。
// 时间 O(n*m)：双重循环。
// 空间 O(m)：标记数组。
func intersectEnumeration(nums1, nums2 []int) []int {
	used := make([]bool, len(nums2))
	res := []int{}
	for _, x := range nums1 {
		for j := range nums2 {
			if !used[j] && nums2[j] == x {
				res = append(res, x)
				used[j] = true
				break
			}
		}
	}
	return res
}
