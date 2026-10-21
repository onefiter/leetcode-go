package combinationsum3

import "slices"

// 216. 组合总和 III
// https://leetcode.cn/problems/combination-sum-iii/
// 找出所有相加之和为 n 的 k 个数的组合，只使用数字 1~9，每个数字最多用一次，
// 返回所有可能的有效组合，同一组合不能出现两次，顺序任意。
//
// 核心思路：从 1~9 中按升序选 k 个数回溯，天然不重复。
// 剪枝：当前数大于剩余和、或剩下的数字不够凑满 k 个时提前结束。

// combinationSum3 解法一：回溯 + 剪枝 ⭐ 最优方案
// 时间复杂度 O(C(9,k)*k)，空间复杂度 O(k)（不计结果）。
func combinationSum3(k int, n int) [][]int {
	res := [][]int{}
	path := make([]int, 0, k)
	var dfs func(start, remain int)
	dfs = func(start, remain int) {
		if len(path) == k {
			if remain == 0 {
				res = append(res, slices.Clone(path))
			}
			return
		}
		// 还需 k-len(path) 个数，i 最大到 9-(k-len(path))+1
		for i := start; i <= 9-(k-len(path))+1; i++ {
			if i > remain {
				break // 升序枚举，后面只会更大
			}
			path = append(path, i)
			dfs(i+1, remain-i)
			path = path[:len(path)-1]
		}
	}
	dfs(1, n)
	return res
}

// combinationSum3Enumeration 解法二：9 位二进制掩码枚举（对照用）
// mask 第 i 位为 1 表示选数字 i+1，枚举 2^9 种选法，筛出选了 k 个且和为 n 的。
// 时间复杂度 O(2^9 * 9)，空间复杂度 O(k)（不计结果）。
func combinationSum3Enumeration(k int, n int) [][]int {
	res := [][]int{}
	for mask := 0; mask < 1<<9; mask++ {
		comb := []int{}
		sum := 0
		for i := 0; i < 9; i++ {
			if mask>>i&1 == 1 {
				comb = append(comb, i+1)
				sum += i + 1
			}
		}
		if len(comb) == k && sum == n {
			res = append(res, comb)
		}
	}
	return res
}
