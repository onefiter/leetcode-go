package shortestdist

import (
	"slices"
	"testing"
)

// cloneGrid 深拷贝二维切片，避免解法修改输入污染测试数据。
func cloneGrid(g [][]int) [][]int {
	out := make([][]int, len(g))
	for i := range g {
		out[i] = slices.Clone(g[i])
	}
	return out
}

var solutions = map[string]func([][]int) int{
	"shortestDistance":          shortestDistance,
	"shortestDistanceFromEmpty": shortestDistanceFromEmpty,
}

func TestShortestDistance(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"示例1", [][]int{{1, 0, 2, 0, 1}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}}, 7},
		{"示例2", [][]int{{1, 0}}, 1},
		{"示例3-无空地", [][]int{{1}}, -1},
		{"建筑被障碍隔开", [][]int{{1, 2, 0}}, -1},
		{"两栋建筑互不连通", [][]int{{1, 0, 2, 0, 1}}, -1},
		{"中间空地", [][]int{{1, 0, 1}}, 2},
		{"四角建筑", [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}, 8},
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
