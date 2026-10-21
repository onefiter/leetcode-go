package subarraysumk

import "testing"

var solutions = map[string]func([]int, int) int{
	"subarraySum":      subarraySum,
	"subarraySumBrute": subarraySumBrute,
}

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"示例1", []int{1, 1, 1}, 2, 2},
		{"示例2", []int{1, 2, 3}, 3, 2},
		{"空数组", []int{}, 0, 0},
		{"单元素命中", []int{5}, 5, 1},
		{"单元素不命中", []int{5}, 3, 0},
		{"含负数", []int{1, -1, 0}, 0, 3},
		{"全零求零", []int{0, 0, 0}, 0, 6},
		{"全相同", []int{2, 2, 2}, 4, 2},
		{"无解", []int{1, 2, 3}, 100, 0},
		{"负数k", []int{-1, -1, 1}, -2, 1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums, tt.k); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.nums, tt.k, got, tt.want)
				}
			})
		}
	}
}
