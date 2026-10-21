package nqueens

import "strings"

// 51. N 皇后
// https://leetcode.cn/problems/n-queens/
// 将 n 个皇后放在 n×n 的棋盘上，使任意两个皇后都不在同一行、同一列或同一斜线上，
// 返回所有不同的解，每个解用 'Q' 表示皇后、'.' 表示空位（顺序任意）。
//
// 核心思路：逐行放置，每行恰好一个皇后，于是只需保证列、主对角线、副对角线不冲突。
// 同一主对角线上 row-col 相同，同一副对角线上 row+col 相同，
// 用三个布尔数组记录占用情况即可 O(1) 判断冲突，回溯时撤销标记。

// solveNQueens 解法一：回溯 + 列/对角线标记数组 ⭐ 最优方案
// 时间复杂度 O(n!)，空间复杂度 O(n)（不计结果）。
func solveNQueens(n int) [][]string {
	res := [][]string{}
	queens := make([]int, n) // queens[row] = 该行皇后所在列
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n) // 下标 row-col+n
	diag2 := make([]bool, 2*n) // 下标 row+col
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			res = append(res, buildBoard(queens))
			return
		}
		for c := 0; c < n; c++ {
			d1, d2 := row-c+n, row+c
			if cols[c] || diag1[d1] || diag2[d2] {
				continue
			}
			queens[row] = c
			cols[c], diag1[d1], diag2[d2] = true, true, true
			dfs(row + 1)
			cols[c], diag1[d1], diag2[d2] = false, false, false
		}
	}
	dfs(0)
	return res
}

// solveNQueensEnumeration 解法二：回溯 + 逐个扫描已放皇后判冲突（对照用）
// 不维护标记数组，每次尝试放置时都和前面所有行的皇后逐一比较列和斜线。
// 时间复杂度 O(n! * n)，空间复杂度 O(n)（不计结果）。
func solveNQueensEnumeration(n int) [][]string {
	res := [][]string{}
	queens := make([]int, n)
	valid := func(row, col int) bool {
		for r := 0; r < row; r++ {
			c := queens[r]
			if c == col || row-r == col-c || row-r == c-col {
				return false
			}
		}
		return true
	}
	var dfs func(row int)
	dfs = func(row int) {
		if row == n {
			res = append(res, buildBoard(queens))
			return
		}
		for c := 0; c < n; c++ {
			if valid(row, c) {
				queens[row] = c
				dfs(row + 1)
			}
		}
	}
	dfs(0)
	return res
}

// buildBoard 把每行皇后的列号转成题目要求的棋盘字符串。
func buildBoard(queens []int) []string {
	n := len(queens)
	board := make([]string, n)
	for r, c := range queens {
		board[r] = strings.Repeat(".", c) + "Q" + strings.Repeat(".", n-c-1)
	}
	return board
}
