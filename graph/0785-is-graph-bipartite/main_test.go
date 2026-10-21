package bipartite

import "testing"

var solutions = map[string]func([][]int) bool{
	"isBipartite":          isBipartite,
	"isBipartiteUnionFind": isBipartiteUnionFind,
}

func TestIsBipartite(t *testing.T) {
	tests := []struct {
		name  string
		graph [][]int
		want  bool
	}{
		{"示例1-含三角形", [][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}, false},
		{"示例2-四元环", [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}, true},
		{"单个节点", [][]int{{}}, true},
		{"无边多节点", [][]int{{}, {}, {}}, true},
		{"不连通且第二块有奇环", [][]int{{1}, {0}, {3, 4}, {2, 4}, {2, 3}}, false},
		{"五元环", [][]int{{1, 4}, {0, 2}, {1, 3}, {2, 4}, {3, 0}}, false},
		{"星形图", [][]int{{1, 2, 3}, {0}, {0}, {0}}, true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.graph); got != tt.want {
					t.Errorf("%s(%v) = %v, want %v", name, tt.graph, got, tt.want)
				}
			})
		}
	}
}
