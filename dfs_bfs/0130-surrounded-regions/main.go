package surrounded

import "leetcode-go/template"

// 130. 被围绕的区域
// https://leetcode.cn/problems/surrounded-regions/
// 给一个由 'X' 和 'O' 组成的矩阵 board，把所有被 'X' 完全围绕的 'O' 区域原地改为 'X'。
// 与边界上的 'O' 相连的 'O' 不会被围绕。
//
// 核心思路：反向思考——从边界上的 'O' 出发搜索，标记所有能连到边界的 'O'（临时改为 '#'），
// 最后遍历一次：剩下的 'O' 都是被围绕的，改为 'X'；'#' 恢复为 'O'。

// solve 解法一：从边界 DFS 标记 ⭐ 最优方案
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)（递归栈最坏情况）。
func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] != 'O' {
			return
		}
		board[r][c] = '#'
		for _, d := range template.Dirs4 {
			dfs(r+d[0], c+d[1])
		}
	}
	for i := 0; i < m; i++ {
		dfs(i, 0)
		dfs(i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(0, j)
		dfs(m-1, j)
	}
	restore(board)
}

// solveBFS 解法二：从边界多源 BFS 标记（对照用）
// 把所有边界 'O' 一起入队，逐层扩散，避免深递归导致栈过深。
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)。
func solveBFS(board [][]byte) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	var queue [][2]int
	mark := func(r, c int) {
		if board[r][c] == 'O' {
			board[r][c] = '#'
			queue = append(queue, [2]int{r, c})
		}
	}
	for i := 0; i < m; i++ {
		mark(i, 0)
		mark(i, n-1)
	}
	for j := 0; j < n; j++ {
		mark(0, j)
		mark(m-1, j)
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range template.Dirs4 {
			r, c := cur[0]+d[0], cur[1]+d[1]
			if r >= 0 && r < m && c >= 0 && c < n {
				mark(r, c)
			}
		}
	}
	restore(board)
}

// restore 把未标记的 'O' 填成 'X'，把标记 '#' 恢复为 'O'。
func restore(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			switch board[i][j] {
			case 'O':
				board[i][j] = 'X'
			case '#':
				board[i][j] = 'O'
			}
		}
	}
}
