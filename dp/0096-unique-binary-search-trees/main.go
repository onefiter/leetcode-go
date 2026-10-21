package numtrees

// 96. 不同的二叉搜索树
// https://leetcode.cn/problems/unique-binary-search-trees/
// 给一个整数 n，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的二叉搜索树有多少种。
//
// 核心思路：G[k] 表示 k 个节点能组成的 BST 种数（只与节点个数有关，与具体值无关）。
// 枚举根节点 i（1 <= i <= k）：左子树有 i-1 个节点，右子树有 k-i 个节点，
// 左右组合数相乘。转移：G[k] = Σ G[i-1] * G[k-i]，G[0] = G[1] = 1（空树也算一种）。
// k 从小到大遍历，保证用到的 G[i-1]、G[k-i] 都已算好。

// numTrees 解法一：区间枚举根节点 DP ⭐ 最优方案
// 时间复杂度 O(n^2)，空间复杂度 O(n)。
func numTrees(n int) int {
	g := make([]int, n+1)
	g[0] = 1
	for k := 1; k <= n; k++ {
		for i := 1; i <= k; i++ {
			g[k] += g[i-1] * g[k-i] // 以 i 为根：左 i-1 个，右 k-i 个
		}
	}
	return g[n]
}

// numTreesCatalan 解法二：卡特兰数公式（对照用）
// G[n] 正是卡特兰数 C(n)，满足递推 C(0) = 1，C(k+1) = C(k) * 2(2k+1) / (k+2)。
// 先乘后除保证整除；n <= 19 时结果不超过 int64。
// 时间复杂度 O(n)，空间复杂度 O(1)。
func numTreesCatalan(n int) int {
	c := 1
	for k := 0; k < n; k++ {
		c = c * 2 * (2*k + 1) / (k + 2)
	}
	return c
}
