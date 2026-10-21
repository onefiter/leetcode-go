package combinations

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

var solutions = map[string]func(int, int) [][]int{
	"combine":            combine,
	"combineEnumeration": combineEnumeration,
}

func TestCombine(t *testing.T) {
	tests := []struct {
		name string
		n    int
		k    int
		want [][]int
	}{
		{"示例1", 4, 2, [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}},
		{"示例2", 1, 1, [][]int{{1}}},
		{"k=n 只有一种", 3, 3, [][]int{{1, 2, 3}}},
		{"k=1 每个数单独成组", 3, 1, [][]int{{1}, {2}, {3}}},
		{"n=5 k=3", 5, 3, [][]int{
			{1, 2, 3}, {1, 2, 4}, {1, 2, 5}, {1, 3, 4}, {1, 3, 5},
			{1, 4, 5}, {2, 3, 4}, {2, 3, 5}, {2, 4, 5}, {3, 4, 5},
		}},
		{"边界-n=2 k=2", 2, 2, [][]int{{1, 2}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := normalize(fn(tt.n, tt.k))
				want := normalize(tt.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("%s(%d, %d) = %v, want %v", name, tt.n, tt.k, got, want)
				}
			})
		}
	}
}
