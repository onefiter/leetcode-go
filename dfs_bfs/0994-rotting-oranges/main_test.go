package rotting

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
	"orangesRotting":         orangesRotting,
	"orangesRottingSimulate": orangesRottingSimulate,
}

func TestOrangesRotting(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"示例1", [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{"示例2-有孤立橘子", [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{"示例3-没有新鲜橘子", [][]int{{0, 2}}, 0},
		{"全空", [][]int{{0}}, 0},
		{"只有新鲜橘子", [][]int{{1}}, -1},
		{"两端同时腐烂", [][]int{{2, 1, 1, 1, 2}}, 2},
		{"多源扩散", [][]int{{2, 1, 1}, {1, 1, 1}, {1, 1, 2}}, 2},
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
