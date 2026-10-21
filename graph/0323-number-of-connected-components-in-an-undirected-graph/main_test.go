package components

import "testing"

var solutions = map[string]func(int, [][]int) int{
	"countComponents":    countComponents,
	"countComponentsDFS": countComponentsDFS,
}

func TestCountComponents(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  int
	}{
		{"示例1-两个分量", 5, [][]int{{0, 1}, {1, 2}, {3, 4}}, 2},
		{"示例2-全连通", 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, 1},
		{"无边", 4, [][]int{}, 4},
		{"单节点", 1, [][]int{}, 1},
		{"含环", 4, [][]int{{0, 1}, {1, 2}, {2, 0}}, 2},
		{"重复边", 3, [][]int{{0, 1}, {1, 0}, {0, 1}}, 2},
		{"星形", 6, [][]int{{0, 1}, {0, 2}, {0, 3}, {0, 4}}, 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.n, tt.edges); got != tt.want {
					t.Errorf("%s(%d, %v) = %d, want %d", name, tt.n, tt.edges, got, tt.want)
				}
			})
		}
	}
}
