package lip

// 329. 矩阵中的最长递增路径
// https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/
// 给一个 m x n 整数矩阵，每个格子可以向上下左右四个方向移动，
// 返回最长「严格递增」路径的长度（不能走对角线，也不能出界）。
//
// 核心思路：严格递增保证不会走回头路，把「小 -> 大」看成有向边，矩阵就是一张 DAG，
// 问题变成求 DAG 最长路。
//   - 记忆化 DFS：memo[i][j] = 1 + max(memo[邻居])，邻居值更大才能走；
//   - 拓扑排序：出度为 0 的格子（周围没有更大值）是终点，从它们开始一层层 BFS 剥洋葱，
//     层数就是最长路径长度。

var dirs = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

// longestIncreasingPath 解法一：记忆化 DFS ⭐ 最优方案
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)。
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	memo := make([][]int, m) // memo[i][j]：从 (i,j) 出发的最长递增路径长度，0 表示未计算
	for i := range memo {
		memo[i] = make([]int, n)
	}
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		if memo[i][j] != 0 {
			return memo[i][j]
		}
		best := 1
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] > matrix[i][j] {
				best = max(best, dfs(x, y)+1)
			}
		}
		memo[i][j] = best
		return best
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans = max(ans, dfs(i, j))
		}
	}
	return ans
}

// longestIncreasingPathTopo 解法二：拓扑排序 BFS（对照用）
// outdeg[i][j] = 周围比它大的邻居个数；出度为 0 的入队，
// 每弹出一层，就把比它小的邻居出度减一，减到 0 入队。BFS 的层数即答案。
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)。
func longestIncreasingPathTopo(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	outdeg := make([][]int, m)
	var queue [][2]int
	for i := 0; i < m; i++ {
		outdeg[i] = make([]int, n)
		for j := 0; j < n; j++ {
			for _, d := range dirs {
				x, y := i+d[0], j+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] > matrix[i][j] {
					outdeg[i][j]++
				}
			}
			if outdeg[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	levels := 0
	for len(queue) > 0 {
		levels++
		var next [][2]int
		for _, c := range queue {
			for _, d := range dirs {
				x, y := c[0]+d[0], c[1]+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && matrix[x][y] < matrix[c[0]][c[1]] {
					outdeg[x][y]--
					if outdeg[x][y] == 0 {
						next = append(next, [2]int{x, y})
					}
				}
			}
		}
		queue = next
	}
	return levels
}
