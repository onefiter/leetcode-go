package triangle

import (
	"slices"
	"testing"
)

var solutions = map[string]func([][]int) int{
	"minimumTotal":        minimumTotal,
	"minimumTotalTopDown": minimumTotalTopDown,
}

func cloneGrid(g [][]int) [][]int {
	res := make([][]int, len(g))
	for i := range g {
		res[i] = slices.Clone(g[i])
	}
	return res
}

func TestMinimumTotal(t *testing.T) {
	tests := []struct {
		name     string
		triangle [][]int
		want     int
	}{
		{"示例1", [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}, 11},
		{"示例2", [][]int{{-10}}, -10},
		{"两行", [][]int{{1}, {2, 3}}, 3},
		{"含负数", [][]int{{-1}, {2, 3}, {1, -1, -3}}, -1},
		{"贪心失效", [][]int{{1}, {2, 3}, {9, 9, 1}}, 5},
		{"全零", [][]int{{0}, {0, 0}, {0, 0, 0}}, 0},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(cloneGrid(tt.triangle)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.triangle, got, tt.want)
				}
			})
		}
	}
}
