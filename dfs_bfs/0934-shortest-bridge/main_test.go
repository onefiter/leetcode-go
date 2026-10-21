package shortestbridge

import (
	"slices"
	"testing"
)

// cloneGrid 深拷贝二维切片，解法会原地修改网格。
func cloneGrid(g [][]int) [][]int {
	out := make([][]int, len(g))
	for i := range g {
		out[i] = slices.Clone(g[i])
	}
	return out
}

var solutions = map[string]func([][]int) int{
	"shortestBridge":      shortestBridge,
	"shortestBridgeBrute": shortestBridgeBrute,
}

func TestShortestBridge(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"示例1", [][]int{{0, 1}, {1, 0}}, 1},
		{"示例2", [][]int{{0, 1, 0}, {0, 0, 0}, {0, 0, 1}}, 2},
		{"示例3-岛中岛", [][]int{{1, 1, 1, 1, 1}, {1, 0, 0, 0, 1}, {1, 0, 1, 0, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 1, 1}}, 1},
		{"对角最远", [][]int{{1, 0, 0}, {0, 0, 0}, {0, 0, 1}}, 3},
		{"同一行隔两格", [][]int{{1, 0, 0, 1}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}, 2},
		{"不规则岛", [][]int{{1, 1, 0, 0}, {1, 0, 0, 0}, {0, 0, 0, 1}, {0, 0, 1, 1}}, 3},
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
