package permutations

import "slices"

// 46. 全排列
// https://leetcode.cn/problems/permutations/
// 给一个不含重复数字的数组 nums，返回其所有可能的全排列（顺序任意）。
//
// 核心思路：回溯时每一层决定「第 k 个位置放谁」，用 used 数组标记已经放过的数，
// 递归到 path 长度等于 n 时得到一个排列，返回时撤销选择（used 置回 false、path 弹出）。

// permute 解法一：回溯 + used 标记 ⭐ 最优方案
// 时间复杂度 O(n * n!)，空间复杂度 O(n)（不计结果，递归深度 + used + path）。
func permute(nums []int) [][]int {
	n := len(nums)
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
			used[i] = true
			path = append(path, nums[i])
			dfs()
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	dfs()
	return res
}

// permuteIterative 解法二：插入法迭代构造（对照用）
// 从 [[]] 开始，每来一个新数字，就把它插入到已有每个排列的每个空隙（共 len+1 个位置）。
// 时间复杂度 O(n * n!)，空间复杂度 O(n * n!)（需要同时保存上一轮所有排列）。
func permuteIterative(nums []int) [][]int {
	res := [][]int{{}}
	for _, x := range nums {
		next := make([][]int, 0, len(res)*(len(res[0])+1))
		for _, p := range res {
			for pos := 0; pos <= len(p); pos++ {
				np := make([]int, 0, len(p)+1)
				np = append(np, p[:pos]...)
				np = append(np, x)
				np = append(np, p[pos:]...)
				next = append(next, np)
			}
		}
		res = next
	}
	return res
}
