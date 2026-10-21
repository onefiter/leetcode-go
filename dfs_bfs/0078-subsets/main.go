package subsets

import "slices"

// 78. 子集
// https://leetcode.cn/problems/subsets/
// 给一个元素互不相同的整数数组 nums，返回它所有可能的子集（幂集），
// 解集不能包含重复子集，顺序任意。
//
// 核心思路：每个元素只有「选 / 不选」两种状态。
// 回溯时 path 记录当前选择，进入每一层递归先把 path 记作一个答案，
// 再从 start 开始依次尝试加入后面的元素，保证组合不重复（只往后选）。

// subsets 解法一：回溯（按起点枚举下一个元素）⭐ 最优方案
// 时间复杂度 O(n*2^n)，空间复杂度 O(n)（不计结果）。
func subsets(nums []int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(start int)
	dfs = func(start int) {
		res = append(res, slices.Clone(path))
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return res
}

// subsetsEnumeration 解法二：二进制位掩码枚举（对照用）
// mask 的第 i 位为 1 表示选 nums[i]，枚举 0 ~ 2^n-1 即得全部子集。
// 时间复杂度 O(n*2^n)，空间复杂度 O(n)（不计结果）。
func subsetsEnumeration(nums []int) [][]int {
	n := len(nums)
	res := make([][]int, 0, 1<<n)
	for mask := 0; mask < 1<<n; mask++ {
		sub := []int{}
		for i := 0; i < n; i++ {
			if mask>>i&1 == 1 {
				sub = append(sub, nums[i])
			}
		}
		res = append(res, sub)
	}
	return res
}
