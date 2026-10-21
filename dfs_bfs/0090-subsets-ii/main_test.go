package subsetsii

import (
	"reflect"
	"slices"
	"testing"
)

// normalize 把 [][]int 规范化：每个内层切片升序排序，外层按字典序排序。
// 用于比较「顺序不定」的子集集合是否相等（重复子集会被保留，从而能检测出去重错误）。
func normalize(groups [][]int) [][]int {
	out := make([][]int, len(groups))
	for i, g := range groups {
		cp := slices.Clone(g)
		slices.Sort(cp)
		if cp == nil {
			cp = []int{}
		}
		out[i] = cp
	}
	slices.SortFunc(out, func(a, b []int) int {
		return slices.Compare(a, b)
	})
	return out
}

var solutions = map[string]func([]int) [][]int{
	"subsetsWithDup":            subsetsWithDup,
	"subsetsWithDupEnumeration": subsetsWithDupEnumeration,
}

func TestSubsetsWithDup(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"示例1", []int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		{"示例2-单元素", []int{0}, [][]int{{}, {0}}},
		{"无重复元素", []int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{"全部相同", []int{2, 2, 2}, [][]int{{}, {2}, {2, 2}, {2, 2, 2}}},
		{"乱序含重复", []int{4, 4, 1, 4}, [][]int{
			{}, {1}, {4}, {1, 4}, {4, 4}, {1, 4, 4}, {4, 4, 4}, {1, 4, 4, 4},
		}},
		{"两组重复", []int{1, 1, 2, 2}, [][]int{
			{}, {1}, {1, 1}, {2}, {2, 2}, {1, 2}, {1, 2, 2}, {1, 1, 2}, {1, 1, 2, 2},
		}},
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
