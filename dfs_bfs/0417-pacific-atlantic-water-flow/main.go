package pacificatlantic

import "leetcode-go/template"

// 417. 太平洋大西洋水流问题
// https://leetcode.cn/problems/pacific-atlantic-water-flow/
// 给一个 m x n 的高度矩阵，左边界和上边界是太平洋，右边界和下边界是大西洋。
// 水只能从高处流向相邻的不高于自己的格子，返回既能流到太平洋又能流到大西洋的所有坐标。
//
// 核心思路：逆向思考——从海洋边界出发「逆流而上」（只走到不低于当前高度的格子），
// 分别标记能到达太平洋和大西洋的格子，两者交集即为答案。每个格子每片海最多访问一次。

// pacificAtlantic 解法一：从两大洋边界逆流 DFS ⭐ 最优方案
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)。
func pacificAtlantic(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	pac, atl := newGrid(m, n), newGrid(m, n)
	var dfs func(r, c int, ocean [][]bool)
	dfs = func(r, c int, ocean [][]bool) {
		ocean[r][c] = true
		for _, d := range template.Dirs4 {
			nr, nc := r+d[0], c+d[1]
			if nr < 0 || nr >= m || nc < 0 || nc >= n || ocean[nr][nc] || heights[nr][nc] < heights[r][c] {
				continue
			}
			dfs(nr, nc, ocean)
		}
	}
	for i := 0; i < m; i++ {
		if !pac[i][0] {
			dfs(i, 0, pac)
		}
		if !atl[i][n-1] {
			dfs(i, n-1, atl)
		}
	}
	for j := 0; j < n; j++ {
		if !pac[0][j] {
			dfs(0, j, pac)
		}
		if !atl[m-1][j] {
			dfs(m-1, j, atl)
		}
	}
	return intersect(pac, atl)
}

// pacificAtlanticBFS 解法二：从两大洋边界多源 BFS（对照用）
// 思路同解法一，把边界格子一次性入队逐层扩展，避免深递归。
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)。
func pacificAtlanticBFS(heights [][]int) [][]int {
	m, n := len(heights), len(heights[0])
	bfs := func(starts [][2]int) [][]bool {
		ocean := newGrid(m, n)
		var queue [][2]int
		for _, s := range starts {
			if !ocean[s[0]][s[1]] {
				ocean[s[0]][s[1]] = true
				queue = append(queue, s)
			}
		}
		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]
			for _, d := range template.Dirs4 {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r < 0 || r >= m || c < 0 || c >= n || ocean[r][c] || heights[r][c] < heights[cur[0]][cur[1]] {
					continue
				}
				ocean[r][c] = true
				queue = append(queue, [2]int{r, c})
			}
		}
		return ocean
	}
	var pacStarts, atlStarts [][2]int
	for i := 0; i < m; i++ {
		pacStarts = append(pacStarts, [2]int{i, 0})
		atlStarts = append(atlStarts, [2]int{i, n - 1})
	}
	for j := 0; j < n; j++ {
		pacStarts = append(pacStarts, [2]int{0, j})
		atlStarts = append(atlStarts, [2]int{m - 1, j})
	}
	return intersect(bfs(pacStarts), bfs(atlStarts))
}

func newGrid(m, n int) [][]bool {
	g := make([][]bool, m)
	for i := range g {
		g[i] = make([]bool, n)
	}
	return g
}

// intersect 收集两张标记表都为 true 的坐标。
func intersect(a, b [][]bool) [][]int {
	res := [][]int{}
	for i := range a {
		for j := range a[i] {
			if a[i][j] && b[i][j] {
				res = append(res, []int{i, j})
			}
		}
	}
	return res
}
