package obstacleelim

import "leetcode-go/template"

// 1293. 网格中的最短路径
// https://leetcode.cn/problems/shortest-path-in-a-grid-with-obstacles-elimination/
// 给一个 m x n 网格，0 为空地，1 为障碍物。每一步可上下左右移动一格，最多可以消除 k 个障碍物。
// 返回从左上角 (0,0) 走到右下角 (m-1,n-1) 的最少步数，无法到达返回 -1。
//
// 核心思路：状态 = (行, 列, 剩余消除次数)，在状态图上做 BFS。
// 优化一：若 k >= m+n-3，可直接走曼哈顿路径（中途最多 m+n-3 个格子），答案为 m+n-2。
// 优化二：对每个格子只记录到达时「剩余消除次数的最大值」，
// 由于 BFS 先到达的步数不会更大，剩余次数不更多的后到状态一定不优，直接剪掉。

// shortestPath 解法一：BFS + 剩余次数剪枝 ⭐ 最优方案
// 时间复杂度 O(m*n*k)，空间复杂度 O(m*n)。
func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if k >= m+n-3 {
		return m + n - 2
	}
	best := make([][]int, m) // 到达该格时剩余消除次数的最大值，-1 表示未到达
	for i := range best {
		best[i] = make([]int, n)
		for j := range best[i] {
			best[i][j] = -1
		}
	}
	best[0][0] = k
	queue := [][3]int{{0, 0, k}}
	for step := 0; len(queue) > 0; step++ {
		var next [][3]int
		for _, cur := range queue {
			if cur[0] == m-1 && cur[1] == n-1 {
				return step
			}
			for _, d := range template.Dirs4 {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r < 0 || r >= m || c < 0 || c >= n {
					continue
				}
				rest := cur[2] - grid[r][c]
				if rest < 0 || rest <= best[r][c] {
					continue
				}
				best[r][c] = rest
				next = append(next, [3]int{r, c, rest})
			}
		}
		queue = next
	}
	return -1
}

// shortestPathStateBFS 解法二：三维 visited 的朴素状态 BFS（对照用）
// visited[r][c][rest] 标记每个完整状态是否访问过，不做剩余次数的支配剪枝。
// 时间复杂度 O(m*n*k)，空间复杂度 O(m*n*k)。
func shortestPathStateBFS(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	visited := make([][][]bool, m)
	for i := range visited {
		visited[i] = make([][]bool, n)
		for j := range visited[i] {
			visited[i][j] = make([]bool, k+1)
		}
	}
	visited[0][0][k] = true
	queue := [][3]int{{0, 0, k}}
	for step := 0; len(queue) > 0; step++ {
		var next [][3]int
		for _, cur := range queue {
			if cur[0] == m-1 && cur[1] == n-1 {
				return step
			}
			for _, d := range template.Dirs4 {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r < 0 || r >= m || c < 0 || c >= n {
					continue
				}
				rest := cur[2] - grid[r][c]
				if rest < 0 || visited[r][c][rest] {
					continue
				}
				visited[r][c][rest] = true
				next = append(next, [3]int{r, c, rest})
			}
		}
		queue = next
	}
	return -1
}
