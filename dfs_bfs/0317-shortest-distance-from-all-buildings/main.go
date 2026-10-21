package shortestdist

import (
	"math"

	"leetcode-go/template"
)

// 317. 离建筑物最近的距离
// https://leetcode.cn/problems/shortest-distance-from-all-buildings/
// 给一个 m x n 网格：0 为空地，1 为建筑，2 为障碍。要在某块空地上建房子，
// 使它到所有建筑的步行距离（只能上下左右走、不能穿过建筑和障碍）之和最小，
// 返回最小距离和；不存在这样的空地则返回 -1。
//
// 核心思路：反过来从每栋建筑出发做 BFS，把步数累加到每块空地上。
// 剪枝：第 k 次 BFS 只走「被前 k 栋建筑都到达过」的空地（用 walk 标记为 -k），
// 走过后标记为 -(k+1)。这样前面某栋建筑到不了的空地会自动被排除，也省去了 visited 数组。

// shortestDistance 解法一：从建筑出发 BFS + 可达性剪枝 ⭐ 最优方案
// 时间复杂度 O(B*m*n)，B 为建筑数量；空间复杂度 O(m*n)。
func shortestDistance(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	walk := make([][]int, m)  // 空地被多少栋建筑访问过（取负数）
	total := make([][]int, m) // 累计距离
	for i := range walk {
		walk[i] = make([]int, n)
		total[i] = make([]int, n)
		for j := range walk[i] {
			if grid[i][j] != 0 {
				walk[i][j] = 1 // 非空地，永远不会等于 -k
			}
		}
	}
	res := -1
	k := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 1 {
				continue
			}
			res = math.MaxInt
			queue := [][2]int{{i, j}}
			for step := 1; len(queue) > 0; step++ {
				var next [][2]int
				for _, cur := range queue {
					for _, d := range template.Dirs4 {
						r, c := cur[0]+d[0], cur[1]+d[1]
						if r < 0 || r >= m || c < 0 || c >= n || walk[r][c] != -k {
							continue
						}
						walk[r][c]--
						total[r][c] += step
						res = min(res, total[r][c])
						next = append(next, [2]int{r, c})
					}
				}
				queue = next
			}
			k++
			// 没有任何空地被目前所有建筑共同到达，后面也不可能有
			if res == math.MaxInt {
				return -1
			}
		}
	}
	return res
}

// shortestDistanceFromEmpty 解法二：从每块空地出发 BFS（对照用）
// 对每块空地 BFS 一遍（只走空地），建筑的距离 = 相邻空地的最小距离 + 1，
// 只有能到达全部建筑时才更新答案。
// 时间复杂度 O(E*m*n)，E 为空地数量（通常远大于建筑数）；空间复杂度 O(m*n)。
func shortestDistanceFromEmpty(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	buildings := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 1 {
				buildings++
			}
		}
	}
	res := -1
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				continue
			}
			dist := template.GridBFS(m, n, i, j, func(r, c int) bool { return grid[r][c] == 0 })
			sum, cnt := 0, 0
			for r := 0; r < m; r++ {
				for c := 0; c < n; c++ {
					if grid[r][c] != 1 {
						continue
					}
					best := -1
					for _, d := range template.Dirs4 {
						nr, nc := r+d[0], c+d[1]
						if nr >= 0 && nr < m && nc >= 0 && nc < n && dist[nr][nc] != -1 {
							if best == -1 || dist[nr][nc]+1 < best {
								best = dist[nr][nc] + 1
							}
						}
					}
					if best != -1 {
						cnt++
						sum += best
					}
				}
			}
			if cnt == buildings && (res == -1 || sum < res) {
				res = sum
			}
		}
	}
	return res
}
