package islands

import "leetcode-go/template"

// 200. 岛屿数量
// https://leetcode.cn/problems/number-of-islands/
// 给你一个由 '1'（陆地）和 '0'（水）组成的二维网格，计算网格中岛屿的数量。
// 岛屿由水平或竖直方向相邻的陆地连接而成，网格四周都被水包围。
//
// 核心思路：每遇到一块未访问的陆地，岛屿数 +1，并把与它相连的整片陆地「淹掉」
// （改成 '0'），保证同一座岛只被计数一次。也可以用并查集合并相邻陆地。

// numIslands 解法一：DFS 淹没 ⭐ 最优方案
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)（最坏递归深度）。
func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != '1' {
			return
		}
		grid[r][c] = '0' // 原地标记已访问
		for _, d := range template.Dirs4 {
			dfs(r+d[0], c+d[1])
		}
	}
	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

// numIslandsUnionFind 解法二：并查集（对照用）
// 每个格子编号 i*n+j，只与右、下相邻陆地合并；最终连通分量数减去水格子数即为答案。
// 时间复杂度 O(m*n*α(m*n))，空间复杂度 O(m*n)。不修改原网格。
func numIslandsUnionFind(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	uf := template.NewUnionFind(m * n)
	water := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != '1' {
				water++
				continue
			}
			if i+1 < m && grid[i+1][j] == '1' {
				uf.Union(i*n+j, (i+1)*n+j)
			}
			if j+1 < n && grid[i][j+1] == '1' {
				uf.Union(i*n+j, i*n+j+1)
			}
		}
	}
	return uf.Count - water
}
