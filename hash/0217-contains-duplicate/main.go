package containsduplicate

import (
	"slices"
	"sort"
)

// 217. 存在重复元素
// https://leetcode.cn/problems/contains-duplicate/
// 给定一个整数数组 nums，如果任一值在数组中出现至少两次，返回 true；
// 如果每个元素互不相同，返回 false。
//
// 核心思路：边遍历边用哈希集合记录见过的值，再次遇到说明有重复。

// containsDuplicate 【解法1】哈希集合 ⭐ 最优
// 时间 O(n)：一次遍历。
// 空间 O(n)：哈希集合。
func containsDuplicate(nums []int) bool {
	seen := make(map[int]struct{}, len(nums))
	for _, x := range nums {
		if _, ok := seen[x]; ok {
			return true
		}
		seen[x] = struct{}{}
	}
	return false
}

// containsDuplicateSort 【解法2】排序后比较相邻元素
// 排序后重复元素必然相邻，扫描一遍即可。
// 时间 O(n*logn)：排序开销。
// 空间 O(n)：克隆避免修改入参。
func containsDuplicateSort(nums []int) bool {
	a := slices.Clone(nums)
	sort.Ints(a)
	for i := 1; i < len(a); i++ {
		if a[i] == a[i-1] {
			return true
		}
	}
	return false
}

// containsDuplicateEnumeration 【解法3】Enumeration 暴力两两比较（对照用）
// 时间 O(n^2)：双重循环。
// 空间 O(1)。
func containsDuplicateEnumeration(nums []int) bool {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
