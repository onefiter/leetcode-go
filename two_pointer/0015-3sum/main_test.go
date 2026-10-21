package threesum

import (
	"reflect"
	"slices"
	"testing"
)

// normalize 把 [][]int 规范化：每个内层切片升序排序，外层按字典序排序。
// 用于比较「顺序不定」的三元组集合是否相等。
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
	if out == nil {
		out = [][]int{}
	}
	return out
}

var solutions = map[string]func([]int) [][]int{
	"threeSum":            threeSum,
	"threeSumEnumeration": threeSumEnumeration,
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"示例1", []int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"全零需去重", []int{0, 0, 0, 0}, [][]int{{0, 0, 0}}},
		{"无解", []int{1, 2, 3}, [][]int{}},
		{"空数组", []int{}, [][]int{}},
		{"全正数无解", []int{1, 1, 1}, [][]int{}},
		{"含大量重复", []int{-2, 0, 0, 2, 2}, [][]int{{-2, 0, 2}}},
		{"含负数多组", []int{-4, -2, -2, -2, 0, 1, 2, 2, 2, 3, 3, 4, 4, 6, 6}, [][]int{
			{-4, -2, 6}, {-4, 0, 4}, {-4, 1, 3}, {-4, 2, 2}, {-2, -2, 4}, {-2, 0, 2},
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
