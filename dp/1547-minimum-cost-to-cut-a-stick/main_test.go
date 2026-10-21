package mincutcost

import (
	"slices"
	"testing"
)

var solutions = map[string]func(int, []int) int{
	"minCost":     minCost,
	"minCostMemo": minCostMemo,
}

func TestMinCost(t *testing.T) {
	tests := []struct {
		name string
		n    int
		cuts []int
		want int
	}{
		{"示例1", 7, []int{1, 3, 4, 5}, 16},
		{"示例2", 9, []int{5, 6, 1, 4, 2}, 22},
		{"只切一刀", 10, []int{5}, 10},
		{"最短木棍", 2, []int{1}, 2},
		{"两刀对称", 10, []int{2, 8}, 18},
		{"每个刻度都切", 5, []int{4, 1, 3, 2}, 12},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.n, slices.Clone(tt.cuts)); got != tt.want {
					t.Errorf("%s(%d, %v) = %d, want %d", name, tt.n, tt.cuts, got, tt.want)
				}
			})
		}
	}
}
