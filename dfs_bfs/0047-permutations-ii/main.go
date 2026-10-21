package permutations2

import "slices"

// 47. 全排列 II
// https://leetcode.cn/problems/permutations-ii/
// 给一个可能包含重复数字的序列 nums，按任意顺序返回所有不重复的全排列。
//
// 核心思路：在 46 题回溯的基础上先排序，让相同数字相邻；
// 规定相同数字只能按「从左到右」的顺序被使用：若 nums[i] == nums[i-1] 且 nums[i-1] 还没用，
// 说明在同一层已经以 nums[i-1] 开过头了，跳过 nums[i]，即可去掉重复排列。

// permuteUnique 解法一：排序 + 回溯 + 同层剪枝 ⭐ 最优方案
// 时间复杂度 O(n * n!)，空间复杂度 O(n)（不计结果）。
func permuteUnique(nums []int) [][]int {
	arr := slices.Clone(nums)
	slices.Sort(arr)
	n := len(arr)
	res := [][]int{}
	path := make([]int, 0, n)
	used := make([]bool, n)
	var dfs func()
	dfs = func() {
		if len(path) == n {
			res = append(res, slices.Clone(path))
			return
		}
		for i := 0; i < n; i++ {
			if used[i] {
				continue
			}
			if i > 0 && arr[i] == arr[i-1] && !used[i-1] {
				continue // 同层已用过相同值作为该位置的数
			}
			used[i] = true
			path = append(path, arr[i])
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return res
}

// permuteUniqueIterative 解法二：下一个排列迭代（对照用）
// 从升序（字典序最小）排列出发，反复调用 nextPermutation 直到变成降序（字典序最大），
// 由于按字典序严格递增生成，天然不会出现重复。
// 时间复杂度 O(n * n!)，空间复杂度 O(1)（不计结果）。
func permuteUniqueIterative(nums []int) [][]int {
	arr := slices.Clone(nums)
	slices.Sort(arr)
	res := [][]int{slices.Clone(arr)}
	for nextPermutation(arr) {
		res = append(res, slices.Clone(arr))
	}
	return res
}

// nextPermutation 原地把 arr 变成字典序的下一个排列，已是最大排列时返回 false。
func nextPermutation(arr []int) bool {
	i := len(arr) - 2
	for i >= 0 && arr[i] >= arr[i+1] {
		i--
	}
	if i < 0 {
		return false
	}
	j := len(arr) - 1
	for arr[j] <= arr[i] {
		j--
	}
	arr[i], arr[j] = arr[j], arr[i]
	slices.Reverse(arr[i+1:])
	return true
}
