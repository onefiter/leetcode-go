package permutations2

import (
	"reflect"
	"slices"
	"testing"
)

// normalize 只对外层按字典序排序（内层是排列，顺序有意义，不能排序）。
// 用于比较「排列出现顺序不定」的结果集合是否相等。
func normalize(groups [][]int) [][]int {
	out := make([][]int, len(groups))
	for i, g := range groups {
		out[i] = slices.Clone(g)
	}
	slices.SortFunc(out, func(a, b []int) int {
		return slices.Compare(a, b)
	})
	return out
}

var solutions = map[string]func([]int) [][]int{
	"permuteUnique":          permuteUnique,
	"permuteUniqueIterative": permuteUniqueIterative,
}

func TestPermuteUnique(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"示例1", []int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
		{"示例2-无重复", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{"单元素", []int{5}, [][]int{{5}}},
		{"全相同", []int{2, 2, 2}, [][]int{{2, 2, 2}}},
		{"两对重复", []int{1, 2, 1, 2}, [][]int{
			{1, 1, 2, 2}, {1, 2, 1, 2}, {1, 2, 2, 1}, {2, 1, 1, 2}, {2, 1, 2, 1}, {2, 2, 1, 1},
		}},
		{"乱序含负数", []int{0, -1, 0}, [][]int{{-1, 0, 0}, {0, -1, 0}, {0, 0, -1}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := normalize(fn(slices.Clone(tt.nums)))
				want := normalize(tt.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.nums, got, want)
				}
			})
		}
	}
}
