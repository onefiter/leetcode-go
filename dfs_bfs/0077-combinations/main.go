package combinations

import (
	"math/bits"
	"slices"
)

// 77. 组合
// https://leetcode.cn/problems/combinations/
// 给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合（顺序任意）。
//
// 核心思路：回溯时用 start 保证组合内元素递增，避免 [1,2] 与 [2,1] 重复；
// 剪枝：还需要选 k-len(path) 个数，所以当前数最多取到 n-(k-len(path))+1，再往后数量不够。

// combine 解法一：回溯 + 剩余数量剪枝 ⭐ 最优方案
// 时间复杂度 O(C(n,k) * k)，空间复杂度 O(k)（不计结果，递归深度 + 路径）。
func combine(n int, k int) [][]int {
	res := [][]int{}
	path := make([]int, 0, k)
	var dfs func(start int)
	dfs = func(start int) {
		if len(path) == k {
			res = append(res, slices.Clone(path))
			return
		}
		// 剩余可选数必须够填满 path
		for i := start; i <= n-(k-len(path))+1; i++ {
			path = append(path, i)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return res
}

// combineEnumeration 解法二：二进制枚举所有子集（对照用）
// 用 n 位掩码表示 [1, n] 的每个子集，只保留恰好有 k 个 1 的掩码并还原成组合。
// 时间复杂度 O(2^n * n)，空间复杂度 O(k)（不计结果）。
func combineEnumeration(n int, k int) [][]int {
	res := [][]int{}
	for mask := 0; mask < 1<<n; mask++ {
		if bits.OnesCount(uint(mask)) != k {
			continue
		}
		comb := make([]int, 0, k)
		for i := 0; i < n; i++ {
			if mask>>i&1 == 1 {
				comb = append(comb, i+1)
			}
		}
		res = append(res, comb)
	}
	return res
}
