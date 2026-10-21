package subsets

import (
	"reflect"
	"slices"
	"testing"
)

// normalize 把 [][]int 规范化：每个内层切片升序排序，外层按字典序排序。
// 用于比较「顺序不定」的子集集合是否相等。
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
	"subsets":            subsets,
	"subsetsEnumeration": subsetsEnumeration,
}

func TestSubsets(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"示例1", []int{1, 2, 3}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{"示例2-单元素", []int{0}, [][]int{{}, {0}}},
		{"空数组", []int{}, [][]int{{}}},
		{"两个元素", []int{5, 9}, [][]int{{}, {5}, {9}, {5, 9}}},
		{"含负数", []int{-1, 2}, [][]int{{}, {-1}, {2}, {-1, 2}}},
		{"乱序输入", []int{3, 1, 2}, [][]int{{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}},
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
