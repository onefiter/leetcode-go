package combinationsum2

import (
	"fmt"
	"slices"
)

// 40. 组合总和 II
// https://leetcode.cn/problems/combination-sum-ii/
// 给一个可能含重复数字的正整数数组 candidates 和目标值 target，找出所有和为 target 的组合，
// candidates 中每个数字在每个组合中只能使用一次，且解集不能包含重复组合。
//
// 核心思路：先排序让相同数字相邻，回溯时在「同一层」遇到与前一个相同的数字就跳过（i > start && nums[i] == nums[i-1]），
// 这样同值数字只会作为该层的第一个分支被展开，从根源上避免重复组合；
// 每个数字只能用一次，所以下一层从 i+1 开始。

// combinationSum2 解法一：排序 + 回溯 + 同层去重 ⭐ 最优方案
// 时间复杂度 O(2^n * n)，空间复杂度 O(n)（递归深度 + 路径）。
func combinationSum2(candidates []int, target int) [][]int {
	nums := slices.Clone(candidates)
	slices.Sort(nums)
	res := [][]int{}
	path := []int{}
	var dfs func(start, remain int)
	dfs = func(start, remain int) {
		if remain == 0 {
			res = append(res, slices.Clone(path))
			return
		}
		for i := start; i < len(nums); i++ {
			if nums[i] > remain {
				break // 已排序，后面只会更大
			}
			if i > start && nums[i] == nums[i-1] {
				continue // 同一层跳过重复值
			}
			path = append(path, nums[i])
			dfs(i+1, remain-nums[i])
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return res
}

// combinationSum2Enumeration 解法二：子集枚举 + 哈希去重（对照用）
// 排序后对每个数字做「选 / 不选」枚举所有子集，和为 target 的子集序列化成字符串放进哈希表去重。
// 时间复杂度 O(2^n * n)（无法利用同层去重剪枝），空间复杂度 O(2^n * n)（去重表）。
func combinationSum2Enumeration(candidates []int, target int) [][]int {
	nums := slices.Clone(candidates)
	slices.Sort(nums) // 排序后相同组合的序列化结果一致
	seen := map[string]bool{}
	res := [][]int{}
	path := []int{}
	var dfs func(i, remain int)
	dfs = func(i, remain int) {
		if remain < 0 {
			return
		}
		if i == len(nums) {
			if remain == 0 {
				key := fmt.Sprint(path)
				if !seen[key] {
					seen[key] = true
					res = append(res, slices.Clone(path))
				}
			}
			return
		}
		path = append(path, nums[i])
		dfs(i+1, remain-nums[i])
		path = path[:len(path)-1]
		dfs(i+1, remain)
	}
	dfs(0, target)
	return res
}
