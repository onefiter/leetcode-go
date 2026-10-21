package foursum

import (
	"reflect"
	"slices"
	"testing"
)

// normalize 把 [][]int 规范化：每个内层切片升序排序，外层按字典序排序。
// 用于比较「顺序不定」的四元组集合是否相等。
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

var solutions = map[string]func([]int, int) [][]int{
	"fourSum":            fourSum,
	"fourSumEnumeration": fourSumEnumeration,
}

func TestFourSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   [][]int
	}{
		{"示例1", []int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{"示例2-全2", []int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{"无解", []int{1, 2, 3, 4}, 100, [][]int{}},
		{"空数组", []int{}, 0, [][]int{}},
		{"不足四个", []int{1, 2, 3}, 6, [][]int{}},
		{"负 target", []int{-3, -2, -1, 0, 0, 1, 2, 3}, -4, [][]int{
			{-3, -2, -1, 2}, {-3, -2, 0, 1}, {-3, -1, 0, 0},
		}},
		{"含重复需去重", []int{0, 0, 0, 0, 0, 0}, 0, [][]int{{0, 0, 0, 0}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := normalize(fn(slices.Clone(tt.nums), tt.target))
				want := normalize(tt.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.nums, tt.target, got, want)
				}
			})
		}
	}
}
