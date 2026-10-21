package permutations

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
	"permute":          permute,
	"permuteIterative": permuteIterative,
}

func TestPermute(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"示例1", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{"示例2", []int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{"示例3-单元素", []int{1}, [][]int{{1}}},
		{"含负数", []int{-1, 2}, [][]int{{-1, 2}, {2, -1}}},
		{"乱序输入", []int{3, 1, 2}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{"边界-两位大数", []int{-10, 10}, [][]int{{-10, 10}, {10, -10}}},
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

	// 规模校验：4 个数字应有 4! = 24 个互不相同的排列
	for name, fn := range solutions {
		got := normalize(fn([]int{1, 2, 3, 4}))
		uniq := slices.CompactFunc(slices.Clone(got), slices.Equal[[]int])
		if len(got) != 24 || len(uniq) != 24 {
			t.Errorf("%s([1 2 3 4]) 结果数 = %d（去重后 %d），want 24", name, len(got), len(uniq))
		}
	}
}
