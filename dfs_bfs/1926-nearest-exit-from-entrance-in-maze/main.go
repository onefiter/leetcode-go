package nearestexit

import "leetcode-go/template"

// 1926. 迷宫中离入口最近的出口
// https://leetcode.cn/problems/nearest-exit-from-entrance-in-maze/
// 给一个 m x n 迷宫 maze，'.' 为空格子，'+' 为墙，以及入口坐标 entrance。
// 出口是位于迷宫边界上的空格子（入口本身不算出口）。每步可上下左右移动一格，
// 返回从入口到最近出口的最少步数，不存在则返回 -1。
//
// 核心思路：从入口出发做 BFS，第一个扩展到的边界空格子就是最近出口。
// 把访问过的格子改为 '+'，省去 visited 数组。

// nearestExit 解法一：从入口单源 BFS ⭐ 最优方案
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)；找到最近出口即可提前结束。
func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	maze[entrance[0]][entrance[1]] = '+'
	queue := [][2]int{{entrance[0], entrance[1]}}
	for step := 1; len(queue) > 0; step++ {
		var next [][2]int
		for _, cur := range queue {
			for _, d := range template.Dirs4 {
				r, c := cur[0]+d[0], cur[1]+d[1]
				if r < 0 || r >= m || c < 0 || c >= n || maze[r][c] == '+' {
					continue
				}
				if r == 0 || r == m-1 || c == 0 || c == n-1 {
					return step
				}
				maze[r][c] = '+'
				next = append(next, [2]int{r, c})
			}
		}
		queue = next
	}
	return -1
}

// nearestExitFromExits 解法二：从所有出口出发的多源 BFS（对照用）
// 把除入口外的所有边界空格子作为起点，反向扩散，第一次到达入口时的步数即为答案。
// 时间复杂度 O(m*n)，空间复杂度 O(m*n)；需要先扫描整个网格找出口，无法像解法一那样尽早结束。
func nearestExitFromExits(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	er, ec := entrance[0], entrance[1]
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var queue [][2]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			onBorder := i == 0 || i == m-1 || j == 0 || j == n-1
			if onBorder && maze[i][j] == '.' && !(i == er && j == ec) {
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
				if r < 0 || r >= m || c < 0 || c >= n || visited[r][c] || maze[r][c] == '+' {
					continue
				}
				if r == er && c == ec {
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
