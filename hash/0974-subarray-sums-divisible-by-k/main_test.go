package subarraydivk

import "testing"

var solutions = map[string]func([]int, int) int{
	"subarraysDivByK":      subarraysDivByK,
	"subarraysDivByKBrute": subarraysDivByKBrute,
}

func TestSubarraysDivByK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"示例1", []int{4, 5, 0, -2, -3, 1}, 5, 7},
		{"示例2", []int{5}, 9, 0},
		{"单元素可整除", []int{5}, 5, 1},
		{"全零", []int{0, 0, 0}, 3, 6},
		{"含负数", []int{-1, 2, 9}, 2, 2},
		{"无解", []int{1, 2, 3}, 100, 0},
		{"全相同", []int{3, 3, 3}, 3, 6},
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
