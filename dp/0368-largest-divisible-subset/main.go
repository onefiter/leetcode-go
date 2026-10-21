package divisible

import "slices"

// 368. 最大整除子集
// https://leetcode.cn/problems/largest-divisible-subset/
// 给一个由互不相同的正整数组成的集合 nums，找出最大的子集，使子集中任意一对 (a, b)
// 都满足 a % b == 0 或 b % a == 0。若有多个解，返回任意一个。
//
// 核心思路：排序后整除关系具有传递性——若 a | b 且 b | c，则 a | c。
// 因此只需保证「相邻元素整除」，问题就变成了类似最长递增子序列的 DP：
// dp[i] 表示以 nums[i] 为最大元素的最大整除子集大小，
// 转移方程 dp[i] = max(dp[j] + 1)（j < i 且 nums[i] % nums[j] == 0），初始 dp[i] = 1；
// 同时记录前驱 prev[i]，最后从 dp 最大的位置沿前驱回溯得到具体子集。

// largestDivisibleSubset 解法一：排序 + LIS 式 DP + 前驱回溯 ⭐ 最优方案
// 时间复杂度 O(n²)，空间复杂度 O(n)。
func largestDivisibleSubset(nums []int) []int {
	a := slices.Clone(nums)
	slices.Sort(a)
	n := len(a)
	dp := make([]int, n)
	prev := make([]int, n)
	best := 0
	for i := 0; i < n; i++ {
		dp[i], prev[i] = 1, -1
		for j := 0; j < i; j++ {
			if a[i]%a[j] == 0 && dp[j]+1 > dp[i] {
				dp[i], prev[i] = dp[j]+1, j
			}
		}
		if dp[i] > dp[best] {
			best = i
		}
	}
	res := []int{}
	for i := best; n > 0 && i != -1; i = prev[i] {
		res = append(res, a[i])
	}
	slices.Reverse(res)
	return res
}

// largestDivisibleSubsetDFS 解法二：排序 + 回溯枚举（对照用）
// 排序后对每个数决定选或不选，选的前提是能被当前子集最后一个元素整除，记录最长路径。
// 时间复杂度 O(2^n)（指数级，仅用于小规模对照），空间复杂度 O(n)。
func largestDivisibleSubsetDFS(nums []int) []int {
	a := slices.Clone(nums)
	slices.Sort(a)
	best := []int{}
	path := []int{}
	var dfs func(i int)
	dfs = func(i int) {
		if len(path)+len(a)-i <= len(best) {
			return // 剪枝：剩余全选也无法超过当前最优
		}
		if i == len(a) {
			best = slices.Clone(path)
			return
		}
		if len(path) == 0 || a[i]%path[len(path)-1] == 0 {
			path = append(path, a[i])
			dfs(i + 1)
			path = path[:len(path)-1]
		}
		dfs(i + 1)
	}
	dfs(0)
	return best
}
