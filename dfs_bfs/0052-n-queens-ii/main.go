package nqueens2

// 52. N 皇后 II
// https://leetcode.cn/problems/n-queens-ii/
// 给一个整数 n，返回 n 皇后问题不同解决方案的数量。
//
// 核心思路：与 51 题一样逐行放置，但只计数不构造棋盘，于是可以用位运算压缩状态：
// cols / ld / rd 三个整数分别表示当前行被「列 / 左斜线 / 右斜线」攻击的位置，
// 可放位置 avail = ^(cols|ld|rd) & mask，每次用 avail & -avail 取最低位逐个尝试，
// 进入下一行时左斜线整体左移一位、右斜线整体右移一位。

// totalNQueens 解法一：位运算回溯 ⭐ 最优方案
// 时间复杂度 O(n!)，空间复杂度 O(n)（递归深度）。
func totalNQueens(n int) int {
	mask := 1<<n - 1
	var dfs func(cols, ld, rd int) int
	dfs = func(cols, ld, rd int) int {
		if cols == mask {
			return 1 // 每列都放了皇后，得到一个解
		}
		count := 0
		avail := ^(cols | ld | rd) & mask
		for avail != 0 {
			bit := avail & -avail // 取最低位的可放位置
			avail &= avail - 1
			count += dfs(cols|bit, (ld|bit)<<1&mask, (rd|bit)>>1)
		}
		return count
	}
	return dfs(0, 0, 0)
}

// totalNQueensEnumeration 解法二：回溯 + 逐个扫描已放皇后判冲突（对照用）
// 每行尝试每一列，与前面各行皇后逐一比较列和斜线是否冲突。
// 时间复杂度 O(n! * n)，空间复杂度 O(n)。
func totalNQueensEnumeration(n int) int {
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
	var dfs func(row int) int
	dfs = func(row int) int {
		if row == n {
			return 1
		}
		count := 0
		for c := 0; c < n; c++ {
			if valid(row, c) {
				queens[row] = c
				count += dfs(row + 1)
			}
		}
		return count
	}
	return dfs(0)
}
