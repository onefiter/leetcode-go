package laststone2

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"lastStoneWeightII":    lastStoneWeightII,
	"lastStoneWeightIISet": lastStoneWeightIISet,
}

func TestLastStoneWeightII(t *testing.T) {
	tests := []struct {
		name   string
		stones []int
		want   int
	}{
		{"示例1", []int{2, 7, 4, 1, 8, 1}, 1},
		{"示例2", []int{31, 26, 33, 21, 40}, 5},
		{"单块石头", []int{5}, 5},
		{"两块相等", []int{3, 3}, 0},
		{"两块不等", []int{1, 10}, 9},
		{"可完全抵消", []int{1, 2, 3}, 0},
		{"一块特别重", []int{1, 1, 1, 100}, 97},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.stones)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.stones, got, tt.want)
				}
			})
		}
	}
}
