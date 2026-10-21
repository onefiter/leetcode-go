package rotting

import "leetcode-go/template"

// 994. 腐烂的橘子
// https://leetcode.cn/problems/rotting-oranges/
// 网格中 0 为空，1 为新鲜橘子，2 为腐烂橘子。每分钟腐烂橘子会让上下左右相邻的新鲜橘子腐烂。
// 返回直到没有新鲜橘子为止所需的最少分钟数，不可能则返回 -1。
//
// 核心思路：所有腐烂橘子同时开始扩散 → 多源 BFS，层数就是分钟数。
// 统计新鲜橘子数量，每腐烂一个就减一，BFS 结束后若仍有剩余则返回 -1。

// orangesRotting 解法一：多源 BFS ⭐ 最优方案
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)。
func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	fresh := 0
	var queue [][2]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 1:
				fresh++
			case 2:
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	minutes := 0
	for len(queue) > 0 && fresh > 0 {
		var next [][2]int
		for _, cur := range queue {
			for _, d := range template.Dirs4 {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] != 1 {
					continue
				}
				grid[r][c] = 2
				fresh--
				next = append(next, [2]int{r, c})
			}
		}
		queue = next
		minutes++
	}
	if fresh > 0 {
		return -1
	}
	return minutes
}

// orangesRottingSimulate 解法二：逐分钟暴力模拟（对照用）
// 每一分钟扫描整张网格，把与腐烂橘子相邻的新鲜橘子标记为本轮腐烂，
// 直到某一轮没有新的橘子腐烂为止。
// 时间复杂度 O((m*n)^2)（最多 m*n 轮，每轮扫描 m*n），空间复杂度 O(m*n)。
func orangesRottingSimulate(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	minutes := 0
	for {
		var toRot [][2]int
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if grid[i][j] != 1 {
					continue
				}
				for _, d := range template.Dirs4 {
					r, c := i+d[0], j+d[1]
					if r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 2 {
						toRot = append(toRot, [2]int{i, j})
						break
					}
				}
			}
		}
		if len(toRot) == 0 {
			break
		}
		// 统一更新，避免本轮新腐烂的橘子在同一分钟继续传染
		for _, p := range toRot {
			grid[p[0]][p[1]] = 2
		}
		minutes++
	}
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}
	return minutes
}
