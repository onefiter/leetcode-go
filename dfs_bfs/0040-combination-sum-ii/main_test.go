package combinationsum2

import (
	"reflect"
	"slices"
	"testing"
)

// normalize 把 [][]int 规范化：每个内层切片升序排序，外层按字典序排序。
// 用于比较「顺序不定」的组合集合是否相等。
func normalize(groups [][]int) [][]int {
	out := make([][]int, len(groups))
	for i, g := range groups {
		cp := slices.Clone(g)
		slices.Sort(cp)
		out[i] = cp
	}
	slices.SortFunc(out, func(a, b []int) int {
		return slices.Compare(a, b)
	})
	return out
}

var solutions = map[string]func([]int, int) [][]int{
	"combinationSum2":            combinationSum2,
	"combinationSum2Enumeration": combinationSum2Enumeration,
}

func TestCombinationSum2(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{"示例1", []int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{"示例2", []int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
		{"无解", []int{3, 4}, 2, [][]int{}},
		{"全相同需去重", []int{1, 1, 1, 1, 1}, 3, [][]int{{1, 1, 1}}},
		{"单元素命中", []int{4}, 4, [][]int{{4}}},
		{"每个数只能用一次", []int{2}, 4, [][]int{}},
		{"全部选中", []int{1, 2, 3}, 6, [][]int{{1, 2, 3}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := normalize(fn(slices.Clone(tt.candidates), tt.target))
				want := normalize(tt.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.candidates, tt.target, got, want)
				}
			})
		}
	}
}
