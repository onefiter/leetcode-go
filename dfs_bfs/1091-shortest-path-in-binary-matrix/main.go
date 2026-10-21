package binarymatrix

// 1091. 二进制矩阵中的最短路径
// https://leetcode.cn/problems/shortest-path-in-binary-matrix/
// 给一个 n x n 的二进制矩阵 grid，返回从左上角 (0,0) 到右下角 (n-1,n-1) 的最短畅通路径长度。
// 畅通路径只经过值为 0 的格子，相邻格子 8 方向连通；路径长度为经过的格子数，不存在则返回 -1。
//
// 核心思路：无权图最短路 → BFS，只是邻居变成 8 个方向。
// 起点或终点为 1 时直接返回 -1；BFS 时直接把访问过的格子置 1，省去 visited 数组。

// dirs8 八个方向。
var dirs8 = [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

// shortestPathBinaryMatrix 解法一：8 方向 BFS ⭐ 最优方案
// 时间复杂度 O(n^2)，空间复杂度 O(n^2)。
func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	grid[0][0] = 1
	queue := [][2]int{{0, 0}}
	for length := 1; len(queue) > 0; length++ {
		var next [][2]int
		for _, cur := range queue {
			if cur[0] == n-1 && cur[1] == n-1 {
				return length
			}
			for _, d := range dirs8 {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r < 0 || r >= n || c < 0 || c >= n || grid[r][c] != 0 {
					continue
				}
				grid[r][c] = 1
				next = append(next, [2]int{r, c})
			}
		}
		queue = next
	}
	return -1
}

// shortestPathBinaryMatrixBidirectional 解法二：双向 BFS（对照用）
// 从起点和终点同时扩展，每次扩展较小的一侧，两侧相遇时路径长度 = 两侧距离之和。
// 最坏时间复杂度同样 O(n^2)，空间复杂度 O(n^2)；在开阔网格上访问的格子更少，但实现更繁琐。
func shortestPathBinaryMatrixBidirectional(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}
	side := make([][]int, n) // 0 未访问，1 起点侧，2 终点侧
	dist := make([][]int, n) // 从所在侧起点算起的格子数（含起点）
	for i := range side {
		side[i] = make([]int, n)
		dist[i] = make([]int, n)
	}
	side[0][0], dist[0][0] = 1, 1
	side[n-1][n-1], dist[n-1][n-1] = 2, 1
	front, back := [][2]int{{0, 0}}, [][2]int{{n - 1, n - 1}}
	for len(front) > 0 && len(back) > 0 {
		if len(front) > len(back) {
			front, back = back, front
		}
		var next [][2]int
		for _, cur := range front {
			s := side[cur[0]][cur[1]]
			for _, d := range dirs8 {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r < 0 || r >= n || c < 0 || c >= n || grid[r][c] != 0 || side[r][c] == s {
					continue
				}
				if side[r][c] != 0 {
					// 碰到另一侧访问过的格子，两段路径拼接
					return dist[cur[0]][cur[1]] + dist[r][c]
				}
				side[r][c], dist[r][c] = s, dist[cur[0]][cur[1]]+1
				next = append(next, [2]int{r, c})
			}
		}
		front = next
	}
	return -1
}
