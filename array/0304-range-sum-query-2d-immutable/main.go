package rangesum2dimmutable

// NumMatrix 304. 二维区域和检索 - 矩阵不可变
// https://leetcode.cn/problems/range-sum-query-2d-immutable/
// 矩阵初始化后不变，需多次查询由 (row1,col1) 左上角到 (row2,col2) 右下角的子矩阵和。

// NumMatrix 【解法一】二维前缀和设计 ⭐
// 预处理 prefix[i][j] = 以 (0,0) 为左上、(i-1,j-1) 为右下的子矩阵元素和。
// 递推式 prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + matrix[i][j]。
// 查询时用容斥：
// region = prefix[r2+1][c2+1] - prefix[r1][c2+1] - prefix[r2+1][c1] + prefix[r1][c1]。
// 构造时间 O(m*n)；每次查询时间 O(1)、空间 O(m*n)。
type NumMatrix struct {
	prefix [][]int
}

// Constructor 预处理二维前缀和。
func Constructor(matrix [][]int) NumMatrix {
	m := len(matrix)
	n := 0
	if m > 0 {
		n = len(matrix[0])
	}
	prefix := make([][]int, m+1)
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{prefix: prefix}
}

// SumRegion 返回左上角 (row1,col1) 到右下角 (row2,col2) 子矩阵的元素和。
func (nm *NumMatrix) SumRegion(row1, col1, row2, col2 int) int {
	p := nm.prefix
	return p[row2+1][col2+1] - p[row1][col2+1] - p[row2+1][col1] + p[row1][col1]
}
