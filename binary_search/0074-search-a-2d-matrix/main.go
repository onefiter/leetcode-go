package searchmatrix

// 74. 搜索二维矩阵
// https://leetcode.cn/problems/search-a-2d-matrix/
// 给定 m x n 矩阵：每行从左到右递增，且每行第一个数大于上一行最后一个数。
// 判断 target 是否在矩阵中。
//
// 一句话理解题意：把矩阵一行接一行拼起来，就是一个完整的有序数组。
//
// 核心思路：把下标 k（0..m*n-1）映射成 (k/n, k%n)，直接在「虚拟一维数组」上二分；
// 或者先二分找所在行，再在行内二分。

// searchMatrix 解法一：展平成一维二分 ⭐ 最优方案
// 虚拟数组第 k 个元素 = matrix[k/n][k%n]，在 [0, m*n-1] 上做标准二分。
// 时间复杂度 O(log(mn))，空间复杂度 O(1)。
func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 || len(matrix[0]) == 0 {
		return false
	}
	n := len(matrix[0])
	lo, hi := 0, m*n-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		v := matrix[mid/n][mid%n]
		switch {
		case v == target:
			return true
		case v < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return false
}

// searchMatrixTwoBinary 解法二：两次二分
// 第一次在第一列上二分，找最后一个「行首 <= target」的行；
// 第二次在该行内二分查找 target。
// 时间复杂度 O(log m + log n)，空间复杂度 O(1)。
func searchMatrixTwoBinary(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 || len(matrix[0]) == 0 {
		return false
	}
	row := -1
	lo, hi := 0, m-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if matrix[mid][0] <= target {
			row = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	if row == -1 {
		return false // target 比所有元素都小
	}
	nums := matrix[row]
	lo, hi = 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		switch {
		case nums[mid] == target:
			return true
		case nums[mid] < target:
			lo = mid + 1
		default:
			hi = mid - 1
		}
	}
	return false
}

// searchMatrixEnumeration 解法三：暴力遍历（对照用）
// 逐个元素比较。
// 时间复杂度 O(mn)，空间复杂度 O(1)。
func searchMatrixEnumeration(matrix [][]int, target int) bool {
	for _, row := range matrix {
		for _, v := range row {
			if v == target {
				return true
			}
		}
	}
	return false
}
