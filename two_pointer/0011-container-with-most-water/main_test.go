package containerwater

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"maxArea":            maxArea,
	"maxAreaEnumeration": maxAreaEnumeration,
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"示例1", []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{"示例2", []int{1, 1}, 1},
		{"递增", []int{1, 2, 3, 4, 5}, 6},
		{"递减", []int{5, 4, 3, 2, 1}, 6},
		{"含零", []int{0, 2, 0}, 0},
		{"两端最高", []int{8, 1, 1, 1, 8}, 32},
		{"空数组", []int{}, 0},
		{"单个元素", []int{5}, 0},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.height)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.height, got, tt.want)
				}
			})
		}
	}
}
