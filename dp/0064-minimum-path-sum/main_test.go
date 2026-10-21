package minpathsum

import (
	"slices"
	"testing"
)

var solutions = map[string]func([][]int) int{
	"minPathSum":   minPathSum,
	"minPathSum2D": minPathSum2D,
}

func cloneGrid(g [][]int) [][]int {
	res := make([][]int, len(g))
	for i := range g {
		res[i] = slices.Clone(g[i])
	}
	return res
}

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"示例1", [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{"示例2", [][]int{{1, 2, 3}, {4, 5, 6}}, 12},
		{"单格", [][]int{{5}}, 5},
		{"单行", [][]int{{1, 2, 3, 4}}, 10},
		{"单列", [][]int{{1}, {2}, {3}}, 6},
		{"全零", [][]int{{0, 0}, {0, 0}}, 0},
		{"绕路更优", [][]int{{1, 9, 1, 1}, {1, 9, 1, 9}, {1, 1, 1, 1}}, 6},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(cloneGrid(tt.grid)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.grid, got, tt.want)
				}
			})
		}
	}
}
