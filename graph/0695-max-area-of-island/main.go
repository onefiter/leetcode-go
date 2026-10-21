package maxarea

import "leetcode-go/template"

// 695. 岛屿的最大面积
// https://leetcode.cn/problems/max-area-of-island/
// 给一个由 0（水）和 1（陆地）组成的 m x n 网格，岛屿是上下左右相连的 1 组成的区域。
// 返回最大岛屿的面积（格子数），没有岛屿则返回 0。
//
// 核心思路：求网格中最大连通块的大小。
// DFS 从每块未访问陆地出发「淹没」整座岛并统计面积；
// 或用并查集把相邻陆地合并，最后取最大集合大小。

// maxAreaOfIsland 解法一：DFS 淹没岛屿 ⭐ 最优方案
// 访问过的陆地原地改为 0，避免额外 visited 数组（会修改入参）。
// 时间复杂度 O(mn)，空间复杂度 O(mn)（最坏递归深度）。
func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var dfs func(r, c int) int
	dfs = func(r, c int) int {
		if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 0 {
			return 0
		}
		grid[r][c] = 0
		area := 1
		for _, d := range template.Dirs4 {
			area += dfs(r+d[0], c+d[1])
		}
		return area
	}
	best := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				best = max(best, dfs(i, j))
			}
		}
	}
	return best
}

// maxAreaOfIslandUnionFind 解法二：并查集（对照用）
// 每块陆地只需与右、下邻居合并，最后取陆地所在集合的最大 Size。
// 不修改入参。时间复杂度 O(mn · α(mn))，空间复杂度 O(mn)。
func maxAreaOfIslandUnionFind(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	uf := template.NewUnionFind(m * n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			if i+1 < m && grid[i+1][j] == 1 {
				uf.Union(i*n+j, (i+1)*n+j)
			}
			if j+1 < n && grid[i][j+1] == 1 {
				uf.Union(i*n+j, i*n+j+1)
			}
		}
	}
	best := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				best = max(best, uf.Size(i*n+j))
			}
		}
	}
	return best
}
