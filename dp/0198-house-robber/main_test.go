package rob

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"rob":     rob,
	"robMemo": robMemo,
}

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{1, 2, 3, 1}, 4},
		{"示例2", []int{2, 7, 9, 3, 1}, 12},
		{"单间房", []int{5}, 5},
		{"两间取大", []int{2, 9}, 9},
		{"隔两间更优", []int{2, 1, 1, 2}, 4},
		{"全零", []int{0, 0, 0}, 0},
		{"空数组", []int{}, 0},
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
