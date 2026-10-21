package combinationsum3

import (
	"reflect"
	"slices"
	"testing"
)

// normalize 把 [][]int 规范化：每个内层切片升序排序，外层按字典序排序。
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

var solutions = map[string]func(int, int) [][]int{
	"combinationSum3":            combinationSum3,
	"combinationSum3Enumeration": combinationSum3Enumeration,
}

func TestCombinationSum3(t *testing.T) {
	tests := []struct {
		name string
		k, n int
		want [][]int
	}{
		{"示例1", 3, 7, [][]int{{1, 2, 4}}},
		{"示例2", 3, 9, [][]int{{1, 2, 6}, {1, 3, 5}, {2, 3, 4}}},
		{"示例3-无解", 4, 1, [][]int{}},
		{"k=1", 1, 5, [][]int{{5}}},
		{"用满九个数", 9, 45, [][]int{{1, 2, 3, 4, 5, 6, 7, 8, 9}}},
		{"和过大无解", 2, 18, [][]int{}},
		{"两数之和", 2, 10, [][]int{{1, 9}, {2, 8}, {3, 7}, {4, 6}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := normalize(fn(tt.k, tt.n))
				want := normalize(tt.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("%s(%d, %d) = %v, want %v", name, tt.k, tt.n, got, want)
				}
			})
		}
	}
}
