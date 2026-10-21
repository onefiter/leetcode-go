package wordsearch

// 79. 单词搜索
// https://leetcode.cn/problems/word-search/
// 给一个 m x n 字符网格 board 和字符串 word，判断 word 是否存在于网格中。
// 单词必须由相邻（上下左右）单元格的字母按顺序构成，同一单元格不能重复使用。
//
// 核心思路：以每个格子为起点做 DFS 回溯，第 k 层匹配 word[k]。
// 访问时把格子原地改成 '#' 作标记，回溯时恢复，省去额外的 visited 数组。

// exist 解法一：DFS 回溯 + 原地标记 ⭐ 最优方案
// 时间复杂度 O(m*n*3^L)（L 为单词长度，除来路外最多 3 个方向），空间复杂度 O(L)。
func exist(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	m, n := len(board), len(board[0])
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		ch := board[i][j]
		board[i][j] = '#' // 标记已访问
		found := dfs(i+1, j, k+1) || dfs(i-1, j, k+1) || dfs(i, j+1, k+1) || dfs(i, j-1, k+1)
		board[i][j] = ch // 回溯恢复
		return found
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// existVisited 解法二：DFS + 独立 visited 数组（对照用）
// 不修改原网格，用额外的二维布尔数组记录访问状态，逻辑更直观但多 O(m*n) 空间。
// 时间复杂度 O(m*n*3^L)，空间复杂度 O(m*n)。
func existVisited(board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}
		visited[i][j] = true
		defer func() { visited[i][j] = false }()
		for _, d := range dirs {
			x, y := i+d[0], j+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] && dfs(x, y, k+1) {
				return true
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
