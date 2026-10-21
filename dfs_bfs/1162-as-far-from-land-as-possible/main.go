package farland

import "leetcode-go/template"

// 1162. 地图分析
// https://leetcode.cn/problems/as-far-from-land-as-possible/
// 给一个 n x n 网格，0 为海洋，1 为陆地。找出一个海洋格子，使它到离它最近的陆地的曼哈顿距离最大，
// 返回该距离；如果网格中只有陆地或只有海洋，返回 -1。
//
// 核心思路：把所有陆地同时作为起点做多源 BFS，向海洋一圈圈扩散，
// 最后被扩散到的海洋格子所在的层数就是答案（4 方向 BFS 的步数恰好等于曼哈顿距离）。

// maxDistance 解法一：多源 BFS ⭐ 最优方案
// 时间复杂度 O(n^2)，空间复杂度 O(n^2)。
func maxDistance(grid [][]int) int {
	n := len(grid)
	var queue [][2]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}
	if len(queue) == 0 || len(queue) == n*n {
		return -1
	}
	step := -1
	for len(queue) > 0 {
		var next [][2]int
		for _, cur := range queue {
			for _, d := range template.Dirs4 {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r < 0 || r >= n || c < 0 || c >= n || grid[r][c] != 0 {
					continue
				}
				grid[r][c] = 2 // 标记已访问
				next = append(next, [2]int{r, c})
			}
		}
		queue = next
		step++
	}
	return step
}

// maxDistanceDP 解法二：两遍动态规划（对照用）
// dist[i][j] 表示到最近陆地的距离：先从左上往右下用上、左更新，再从右下往左上用下、右更新。
// 时间复杂度 O(n^2)，空间复杂度 O(n^2)。
func maxDistanceDP(grid [][]int) int {
	n := len(grid)
	const inf = 1 << 30
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			if grid[i][j] == 0 {
				dist[i][j] = inf
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				dist[i][j] = min(dist[i][j], dist[i-1][j]+1)
			}
			if j > 0 {
				dist[i][j] = min(dist[i][j], dist[i][j-1]+1)
			}
		}
	}
	res := 0
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i < n-1 {
				dist[i][j] = min(dist[i][j], dist[i+1][j]+1)
			}
			if j < n-1 {
				dist[i][j] = min(dist[i][j], dist[i][j+1]+1)
			}
			if grid[i][j] == 0 {
				res = max(res, dist[i][j])
			}
		}
	}
	// 没有海洋（res==0）或没有陆地（res 仍为 inf 量级）
	if res == 0 || res >= inf {
		return -1
	}
	return res
}
