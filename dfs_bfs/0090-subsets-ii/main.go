package subsetsii

import "slices"

// 90. 子集 II
// https://leetcode.cn/problems/subsets-ii/
// 给一个可能包含重复元素的整数数组 nums，返回所有可能的子集（幂集），
// 解集不能包含重复的子集，顺序任意。
//
// 核心思路：先排序让相同元素相邻，再按 78 题的方式回溯。
// 同一层（同一个 start 下的 for 循环）里，相同的值只选第一个，
// 即 i > start && nums[i] == nums[i-1] 时跳过，这样就去掉了重复子集。

// subsetsWithDup 解法一：排序 + 回溯同层剪枝 ⭐ 最优方案
// 时间复杂度 O(n*2^n)，空间复杂度 O(n)（不计结果）。
func subsetsWithDup(nums []int) [][]int {
	nums = slices.Clone(nums)
	slices.Sort(nums)
	res := [][]int{}
	path := []int{}
	var dfs func(start int)
	dfs = func(start int) {
		res = append(res, slices.Clone(path))
		for i := start; i < len(nums); i++ {
			if i > start && nums[i] == nums[i-1] {
				continue // 同层重复值只取第一个
			}
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

// subsetsWithDupEnumeration 解法二：排序 + 位掩码枚举（对照用）
// 排序后，若选了 nums[i] 而没选与它相等的 nums[i-1]，该掩码会和
// 「选 nums[i-1] 不选 nums[i]」产生同一子集，直接丢弃即可去重。
// 时间复杂度 O(n*2^n)，空间复杂度 O(n)（不计结果）。
func subsetsWithDupEnumeration(nums []int) [][]int {
	nums = slices.Clone(nums)
	slices.Sort(nums)
	n := len(nums)
	res := [][]int{}
	for mask := 0; mask < 1<<n; mask++ {
		sub := []int{}
		valid := true
		for i := 0; i < n; i++ {
			if mask>>i&1 == 0 {
				continue
			}
			if i > 0 && nums[i] == nums[i-1] && mask>>(i-1)&1 == 0 {
				valid = false
				break
			}
			sub = append(sub, nums[i])
		}
		if valid {
			res = append(res, sub)
		}
	}
	return res
}
