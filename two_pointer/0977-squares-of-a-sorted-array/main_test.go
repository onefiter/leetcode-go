package sortedsquares

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) []int{
	"sortedSquares":            sortedSquares,
	"sortedSquaresEnumeration": sortedSquaresEnumeration,
}

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"示例1", []int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{"示例2", []int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{"空数组", []int{}, []int{}},
		{"单个元素", []int{-5}, []int{25}},
		{"全负数", []int{-3, -2, -1}, []int{1, 4, 9}},
		{"全正数", []int{1, 2, 3}, []int{1, 4, 9}},
		{"全相同", []int{2, 2, 2}, []int{4, 4, 4}},
		{"绝对值对称", []int{-2, -1, 1, 2}, []int{1, 1, 4, 4}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}
