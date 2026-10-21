package shortestbridge

import "leetcode-go/template"

// 934. 最短的桥
// https://leetcode.cn/problems/shortest-bridge/
// 给一个 n x n 的二进制矩阵 grid，1 表示陆地，0 表示水域，矩阵中恰好有两座岛（4 方向连通）。
// 可以把 0 翻转成 1 来连接两座岛，返回最少需要翻转的 0 的数目。
//
// 核心思路：先 DFS 找到第一座岛并全部标记为 2，同时把这些格子作为多源 BFS 的起点；
// 然后向外一层层扩展水域，第一次碰到值为 1 的格子（第二座岛）时，已扩展的层数就是答案。

// shortestBridge 解法一：DFS 标记一座岛 + 多源 BFS ⭐ 最优方案
// 时间复杂度 O(n^2)，空间复杂度 O(n^2)。
func shortestBridge(grid [][]int) int {
	n := len(grid)
	var queue [][2]int
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= n || c < 0 || c >= n || grid[r][c] != 1 {
			return
		}
		grid[r][c] = 2
		queue = append(queue, [2]int{r, c})
		for _, d := range template.Dirs4 {
			dfs(r+d[0], c+d[1])
		}
	}
	// 找到第一块陆地并标记整座岛
	found := false
	for i := 0; i < n && !found; i++ {
		for j := 0; j < n && !found; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				found = true
			}
		}
	}
	for step := 0; len(queue) > 0; step++ {
		var next [][2]int
		for _, cur := range queue {
			for _, d := range template.Dirs4 {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r < 0 || r >= n || c < 0 || c >= n || grid[r][c] == 2 {
					continue
				}
				if grid[r][c] == 1 {
					return step
				}
				grid[r][c] = 2 // 水域被「填平」，并入第一座岛
				next = append(next, [2]int{r, c})
			}
		}
		queue = next
	}
	return -1
}

// shortestBridgeBrute 解法二：枚举两岛格子对求最小曼哈顿距离（对照用）
// 用 DFS 分出两座岛的全部格子，两两计算 |dr|+|dc|-1 取最小值。
// 距离最近的一对格子之间的曼哈顿路径不会经过任何陆地（否则会有更近的一对），所以该值就是答案。
// 时间复杂度 O(A*B)，最坏 O(n^4)；空间复杂度 O(n^2)。
func shortestBridgeBrute(grid [][]int) int {
	n := len(grid)
	var islands [][][2]int
	var dfs func(r, c int, cells *[][2]int)
	dfs = func(r, c int, cells *[][2]int) {
		if r < 0 || r >= n || c < 0 || c >= n || grid[r][c] != 1 {
			return
		}
		grid[r][c] = 2
		*cells = append(*cells, [2]int{r, c})
		for _, d := range template.Dirs4 {
			dfs(r+d[0], c+d[1], cells)
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				var cells [][2]int
				dfs(i, j, &cells)
				islands = append(islands, cells)
			}
		}
	}
	if len(islands) < 2 {
		return -1
	}
	res := 2 * n
	for _, a := range islands[0] {
		for _, b := range islands[1] {
			res = min(res, abs(a[0]-b[0])+abs(a[1]-b[1])-1)
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
