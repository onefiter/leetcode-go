package threesumclosest

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int) int{
	"threeSumClosest":            threeSumClosest,
	"threeSumClosestEnumeration": threeSumClosestEnumeration,
}

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"示例1", []int{-1, 2, 1, -4}, 1, 2},
		{"示例2", []int{0, 0, 0}, 1, 0},
		{"含重复元素", []int{1, 1, 1, 0}, -100, 2},
		{"含负数", []int{-3, -2, -5, 3, -4}, -1, -2},
		{"恰好命中", []int{1, 2, 3, 4}, 6, 6},
		{"全负数", []int{-5, -4, -3, -2, -1}, -8, -8},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.nums), tt.target); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.nums, tt.target, got, tt.want)
				}
			})
		}
	}
}
