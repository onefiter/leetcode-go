package maxarea

import (
	"slices"
	"testing"
)

// cloneGrid 深拷贝网格，解法一会原地修改入参。
func cloneGrid(grid [][]int) [][]int {
	out := make([][]int, len(grid))
	for i, row := range grid {
		out[i] = slices.Clone(row)
	}
	return out
}

var solutions = map[string]func([][]int) int{
	"maxAreaOfIsland":          maxAreaOfIsland,
	"maxAreaOfIslandUnionFind": maxAreaOfIslandUnionFind,
}

func TestMaxAreaOfIsland(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"示例1", [][]int{
			{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
			{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
		}, 6},
		{"示例2-全是水", [][]int{{0, 0, 0, 0, 0, 0, 0, 0}}, 0},
		{"单格陆地", [][]int{{1}}, 1},
		{"全是陆地", [][]int{{1, 1}, {1, 1}}, 4},
		{"对角不相连", [][]int{{1, 0}, {0, 1}}, 1},
		{"U 形岛屿", [][]int{{1, 0, 1}, {1, 0, 1}, {1, 1, 1}}, 7},
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
