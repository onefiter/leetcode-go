package rob2

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"rob":      rob,
	"robTable": robTable,
}

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{2, 3, 2}, 3},
		{"示例2", []int{1, 2, 3, 1}, 4},
		{"示例3", []int{1, 2, 3}, 3},
		{"单间房", []int{7}, 7},
		{"两间取大", []int{4, 9}, 9},
		{"首尾都大只能取一", []int{5, 1, 1, 5}, 6},
		{"较长环", []int{200, 3, 140, 20, 10}, 340},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.nums)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}
