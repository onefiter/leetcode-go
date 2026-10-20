package template

import (
	"cmp"
	"slices"
)

// SortPairs 对二维数组按第一维升序、第一维相同按第二维升序排序（常用于区间题）。
func SortPairs(pairs [][]int) {
	slices.SortFunc(pairs, func(a, b []int) int {
		if c := cmp.Compare(a[0], b[0]); c != 0 {
			return c
		}
		return cmp.Compare(a[1], b[1])
	})
}

// SortDesc 降序排序。
func SortDesc[T cmp.Ordered](a []T) {
	slices.SortFunc(a, func(x, y T) int { return cmp.Compare(y, x) })
}

// SortedIndices 返回按 a[i] 升序排列的下标（稳定排序，不修改原数组）。
func SortedIndices[T cmp.Ordered](a []T) []int {
	idx := make([]int, len(a))
	for i := range idx {
		idx[i] = i
	}
	slices.SortStableFunc(idx, func(i, j int) int { return cmp.Compare(a[i], a[j]) })
	return idx
}
