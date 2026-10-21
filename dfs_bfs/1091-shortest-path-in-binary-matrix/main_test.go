package binarymatrix

import (
	"slices"
	"testing"
)

// cloneGrid 深拷贝二维切片，解法一会原地修改网格。
func cloneGrid(g [][]int) [][]int {
	out := make([][]int, len(g))
	for i := range g {
		out[i] = slices.Clone(g[i])
	}
	return out
}

var solutions = map[string]func([][]int) int{
	"shortestPathBinaryMatrix":              shortestPathBinaryMatrix,
	"shortestPathBinaryMatrixBidirectional": shortestPathBinaryMatrixBidirectional,
}

func TestShortestPathBinaryMatrix(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"示例1", [][]int{{0, 1}, {1, 0}}, 2},
		{"示例2", [][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}}, 4},
		{"示例3-起点被堵", [][]int{{1, 0, 0}, {1, 1, 0}, {1, 1, 0}}, -1},
		{"单格可通", [][]int{{0}}, 1},
		{"终点被堵", [][]int{{0, 0}, {0, 1}}, -1},
		{"对角直达", [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}, 3},
		{"被墙隔断", [][]int{{0, 1, 0}, {1, 1, 0}, {0, 0, 0}}, -1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(cloneGrid(tt.grid))
				if got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.grid, got, tt.want)
				}
			})
		}
	}
}
