package getfood

import "leetcode-go/template"

// 1730. 获取食物的最短路径
// https://leetcode.cn/problems/shortest-path-to-get-food/
// 给一个 m x n 字符网格：'*' 是你的位置（恰好一个），'#' 是食物（可能多个），
// 'O' 是空地，'X' 是障碍。每步可上下左右移动，返回到达任意食物的最短步数，无法到达返回 -1。
//
// 核心思路：从 '*' 出发做 BFS，第一次扩展到 '#' 时的步数就是答案。
// 把走过的格子改为 'X'，省去 visited 数组。

// getFood 解法一：从起点单源 BFS ⭐ 最优方案
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)；找到最近的食物即可提前结束。
func getFood(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	var queue [][2]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '*' {
				queue = append(queue, [2]int{i, j})
				grid[i][j] = 'X'
			}
		}
	}
	for step := 1; len(queue) > 0; step++ {
		var next [][2]int
		for _, cur := range queue {
			for _, d := range template.Dirs4 {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r < 0 || r >= m || c < 0 || c >= n || grid[r][c] == 'X' {
					continue
				}
				if grid[r][c] == '#' {
					return step
				}
				grid[r][c] = 'X'
				next = append(next, [2]int{r, c})
			}
		}
		queue = next
	}
	return -1
}

// getFoodFromFood 解法二：从所有食物出发的多源 BFS（对照用）
// 把所有 '#' 作为起点反向扩散，第一次到达 '*' 时的步数即为答案。
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)；食物很多时需要先全部入队，通常比解法一访问更多格子。
func getFoodFromFood(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	var queue [][2]int
	for i := range visited {
		visited[i] = make([]bool, n)
		for j := range visited[i] {
			if grid[i][j] == '#' {
				visited[i][j] = true
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	for step := 1; len(queue) > 0; step++ {
		var next [][2]int
		for _, cur := range queue {
			for _, d := range template.Dirs4 {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] || grid[r][c] == 'X' {
					continue
				}
				if grid[r][c] == '*' {
					return step
				}
				visited[r][c] = true
				next = append(next, [2]int{r, c})
			}
		}
		queue = next
	}
	return -1
}
