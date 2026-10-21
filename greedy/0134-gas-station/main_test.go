package gasstation

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, []int) int{
	"canCompleteCircuit":      canCompleteCircuit,
	"canCompleteCircuitBrute": canCompleteCircuitBrute,
}

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name string
		gas  []int
		cost []int
		want int
	}{
		{"示例1", []int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{"示例2", []int{2, 3, 4}, []int{3, 4, 3}, -1},
		{"单站够用", []int{5}, []int{4}, 0},
		{"单站不够", []int{1}, []int{2}, -1},
		{"恰好用完", []int{1, 1, 1}, []int{1, 1, 1}, 0},
		{"起点在末尾", []int{1, 1, 5}, []int{2, 2, 1}, 2},
		{"总量不足", []int{3, 1, 1}, []int{1, 2, 3}, -1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(slices.Clone(tt.gas), slices.Clone(tt.cost))
				if got != tt.want {
					t.Errorf("%s(%v, %v) = %d, want %d", name, tt.gas, tt.cost, got, tt.want)
				}
			})
		}
	}
}
