package intervalintersection

import (
	"slices"
	"testing"
)

var solutions = map[string]func([][]int, [][]int) [][]int{
	"intervalIntersection":            intervalIntersection,
	"intervalIntersectionEnumeration": intervalIntersectionEnumeration,
}

func TestIntervalIntersection(t *testing.T) {
	tests := []struct {
		name          string
		first, second [][]int
		want          [][]int
	}{
		{
			"示例1",
			[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{"示例2-一方为空", [][]int{{1, 3}, {5, 9}}, [][]int{}, [][]int{}},
		{"都为空", [][]int{}, [][]int{}, [][]int{}},
		{"无交集", [][]int{{1, 2}}, [][]int{{3, 4}}, [][]int{}},
		{"端点相接", [][]int{{1, 3}}, [][]int{{3, 5}}, [][]int{{3, 3}}},
		{"完全相同", [][]int{{1, 5}}, [][]int{{1, 5}}, [][]int{{1, 5}}},
		{"包含关系", [][]int{{0, 10}}, [][]int{{1, 2}, {4, 6}, {8, 12}}, [][]int{{1, 2}, {4, 6}, {8, 10}}},
		{"单点区间", [][]int{{3, 3}}, [][]int{{1, 5}}, [][]int{{3, 3}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.first, tt.second)
				if !slices.EqualFunc(got, tt.want, slices.Equal) {
					t.Errorf("%s(%v, %v) = %v, want %v", name, tt.first, tt.second, got, tt.want)
				}
			})
		}
	}
}
