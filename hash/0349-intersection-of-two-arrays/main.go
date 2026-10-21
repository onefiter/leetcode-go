package intersection

import (
	"slices"
	"sort"
)

// 349. 两个数组的交集
// https://leetcode.cn/problems/intersection-of-two-arrays/
// 给定两个数组 nums1 和 nums2，返回它们的交集。输出结果中每个元素「唯一」，
// 不考虑输出顺序。例如 nums1=[1,2,2,1], nums2=[2,2] -> [2]。
//
// 核心思路：把其中一个数组的元素放进哈希集合，遍历另一个数组时若命中则收集，
// 并从集合里删除以保证结果去重。

// intersection 【解法1】哈希集合去重求交 ⭐ 最优
// 时间 O(n+m)：分别遍历两个数组。
// 空间 O(n)：哈希集合。
func intersection(nums1, nums2 []int) []int {
	set := make(map[int]struct{}, len(nums1))
	for _, x := range nums1 {
		set[x] = struct{}{}
	}
	res := []int{}
	for _, x := range nums2 {
		if _, ok := set[x]; ok {
			res = append(res, x)
			delete(set, x) // 删除避免重复加入
		}
	}
	return res
}

// intersectionSort 【解法2】排序 + 双指针
// 两个数组分别排序后，用双指针同步前进：相等则收集（并跳过各自的重复），
// 较小的一侧前进。
// 时间 O(n*logn + m*logm)：排序开销。
// 空间 O(n+m)：克隆避免修改入参。
func intersectionSort(nums1, nums2 []int) []int {
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
			// 相等，收集一次后跳过两侧的重复值。
			if len(res) == 0 || res[len(res)-1] != a[i] {
				res = append(res, a[i])
			}
			i++
			j++
		}
	}
	return res
}

// intersectionEnumeration 【解法3】Enumeration 暴力（对照用）
// 对 nums1 去重后的每个值，在 nums2 中线性查找是否存在。
// 时间 O(n*m)：双重循环。
// 空间 O(n)：去重集合。
func intersectionEnumeration(nums1, nums2 []int) []int {
	res := []int{}
	added := make(map[int]struct{})
	for _, x := range nums1 {
		if _, done := added[x]; done {
			continue
		}
		if slices.Contains(nums2, x) {
			res = append(res, x)
			added[x] = struct{}{}
		}
	}
	return res
}
