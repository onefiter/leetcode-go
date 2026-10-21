package combinationsum

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
	"combinationSum":            combinationSum,
	"combinationSumEnumeration": combinationSumEnumeration,
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{"示例1", []int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{"示例2", []int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{"示例3-无解", []int{2}, 1, [][]int{}},
		{"单元素重复选", []int{1}, 3, [][]int{{1, 1, 1}}},
		{"乱序输入", []int{7, 3, 2}, 7, [][]int{{2, 2, 3}, {7}}},
		{"全部大于目标", []int{5, 6}, 4, [][]int{}},
		{"多组解", []int{1, 2}, 4, [][]int{{1, 1, 1, 1}, {1, 1, 2}, {2, 2}}},
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
