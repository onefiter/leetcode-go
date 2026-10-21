package obstacleelim

import (
	"slices"
	"testing"
)

// cloneGrid 深拷贝二维切片。
func cloneGrid(g [][]int) [][]int {
	out := make([][]int, len(g))
	for i := range g {
		out[i] = slices.Clone(g[i])
	}
	return out
}

var solutions = map[string]func([][]int, int) int{
	"shortestPath":         shortestPath,
	"shortestPathStateBFS": shortestPathStateBFS,
}

func TestShortestPath(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		k    int
		want int
	}{
		{"示例1", [][]int{{0, 0, 0}, {1, 1, 0}, {0, 0, 0}, {0, 1, 1}, {0, 0, 0}}, 1, 6},
		{"示例2", [][]int{{0, 1, 1}, {1, 1, 1}, {1, 0, 0}}, 1, -1},
		{"单格", [][]int{{0}}, 0, 0},
		{"无需消除", [][]int{{0, 0}, {0, 0}}, 0, 2},
		{"k 足够大直接曼哈顿", [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 0}}, 3, 4},
		{"k=0 需绕路", [][]int{{0, 1, 0, 0}, {0, 1, 0, 1}, {0, 0, 0, 0}}, 0, 5},
		{"k=0 无路可走", [][]int{{0, 1}, {1, 0}}, 0, -1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(cloneGrid(tt.grid), tt.k)
				if got != tt.want {
					t.Errorf("%s(%v, k=%d) = %d, want %d", name, tt.grid, tt.k, got, tt.want)
				}
			})
		}
	}
}
