package combinationsum

import "slices"

// 39. 组合总和
// https://leetcode.cn/problems/combination-sum/
// 给一个无重复元素的正整数数组 candidates 和目标值 target，找出所有和为 target 的不同组合，
// candidates 中的同一个数字可以无限制重复选取（顺序任意）。
//
// 核心思路：回溯时用 start 下标保证「只往后选」，避免 [2,3] 和 [3,2] 这类重复组合；
// 因为可以重复选，递归时下一层仍从 i 开始。先排序后，一旦 candidates[i] > 剩余值即可整体剪枝。

// combinationSum 解法一：排序 + 回溯剪枝 ⭐ 最优方案
// 时间复杂度 O(S)，S 为所有可行解长度之和（剪枝后接近输出规模），空间复杂度 O(target)（递归深度）。
func combinationSum(candidates []int, target int) [][]int {
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
				break // 已排序，后面更大，直接剪枝
			}
			path = append(path, nums[i])
			dfs(i, remain-nums[i]) // 可重复选，仍从 i 开始
			path = path[:len(path)-1]
		}
	}
	dfs(0, target)
	return res
}

// combinationSumEnumeration 解法二：选 / 不选 二叉枚举（对照用）
// 每层只决定当前数字「再选一次」还是「跳过它看下一个」，不排序、不做提前剪枝，
// 只在剩余值 < 0 时回退，搜索树更深更宽。
// 时间复杂度 O(n^(target/min))，空间复杂度 O(target/min + n)。
func combinationSumEnumeration(candidates []int, target int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(i, remain int)
	dfs = func(i, remain int) {
		if remain == 0 {
			res = append(res, slices.Clone(path))
			return
		}
		if i == len(candidates) || remain < 0 {
			return
		}
		// 跳过 candidates[i]
		dfs(i+1, remain)
		// 选 candidates[i]（下标不动，允许重复）
		path = append(path, candidates[i])
		dfs(i, remain-candidates[i])
		path = path[:len(path)-1]
	}
	dfs(0, target)
	return res
}
