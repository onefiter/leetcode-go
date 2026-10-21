package validtree

import "testing"

var solutions = map[string]func(int, [][]int) bool{
	"validTree":    validTree,
	"validTreeBFS": validTreeBFS,
}

func TestValidTree(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		edges [][]int
		want  bool
	}{
		{"示例1-合法树", 5, [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}}, true},
		{"示例2-有环", 5, [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}}, false},
		{"单节点", 1, [][]int{}, true},
		{"不连通", 4, [][]int{{0, 1}, {2, 3}}, false},
		{"边数对但有环且不连通", 4, [][]int{{0, 1}, {1, 2}, {2, 0}}, false},
		{"链", 4, [][]int{{0, 1}, {1, 2}, {2, 3}}, true},
		{"两节点无边", 2, [][]int{}, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.n, tt.edges); got != tt.want {
					t.Errorf("%s(%d, %v) = %v, want %v", name, tt.n, tt.edges, got, tt.want)
				}
			})
		}
	}
}
