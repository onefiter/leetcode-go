package updatematrix

import "leetcode-go/template"

// 542. 01 矩阵
// https://leetcode.cn/problems/01-matrix/
// 给一个由 0 和 1 组成的矩阵 mat，返回同样大小的矩阵，
// 其中每个格子是它到最近的 0 的距离（相邻格子距离为 1）。
//
// 核心思路：把所有 0 同时作为起点做多源 BFS，第一次到达某个 1 时的层数就是答案。
// 也可以用 DP：最近的 0 要么在左上方向、要么在右下方向，正反各扫一遍取最小值。

// updateMatrix 解法一：多源 BFS ⭐ 最优方案
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)。
func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	dist := make([][]int, m)
	var queue [][2]int
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if mat[i][j] == 0 {
				queue = append(queue, [2]int{i, j})
			} else {
				dist[i][j] = -1 // 未访问
			}
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range template.Dirs4 {
			r, c := cur[0]+d[0], cur[1]+d[1]
			if r < 0 || r >= m || c < 0 || c >= n || dist[r][c] != -1 {
				continue
			}
			dist[r][c] = dist[cur[0]][cur[1]] + 1
			queue = append(queue, [2]int{r, c})
		}
	}
	return dist
}

// updateMatrixDP 解法二：两遍动态规划（对照用）
// 第一遍从左上到右下，用上方和左方更新；第二遍从右下到左上，用下方和右方更新。
// 时间复杂度 O(m*n)，空间复杂度 O(1)（不计输出）。
func updateMatrixDP(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	const inf = 1 << 30
	dist := make([][]int, m)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if mat[i][j] != 0 {
				dist[i][j] = inf
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
			}
			if j > 0 {
				dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i < m-1 {
				dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
			}
			if j < n-1 {
				dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
			}
		}
	}
	return dist
}
