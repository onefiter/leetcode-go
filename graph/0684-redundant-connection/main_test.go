package redundant

import (
	"slices"
	"testing"
)

var solutions = map[string]func([][]int) []int{
	"findRedundantConnection":    findRedundantConnection,
	"findRedundantConnectionDFS": findRedundantConnectionDFS,
}

func TestFindRedundantConnection(t *testing.T) {
	tests := []struct {
		name  string
		edges [][]int
		want  []int
	}{
		{"示例1", [][]int{{1, 2}, {1, 3}, {2, 3}}, []int{2, 3}},
		{"示例2", [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, []int{1, 4}},
		{"环在中间", [][]int{{1, 2}, {2, 3}, {3, 1}, {3, 4}}, []int{3, 1}},
		{"大环最后闭合", [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 1}}, []int{5, 1}},
		{"环不含节点2", [][]int{{1, 4}, {3, 4}, {1, 3}, {1, 2}, {4, 5}}, []int{1, 3}},
		{"多分支后成环", [][]int{{1, 2}, {1, 3}, {2, 4}, {3, 5}, {4, 6}, {5, 6}}, []int{5, 6}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.edges); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.edges, got, tt.want)
				}
			})
		}
	}
}
