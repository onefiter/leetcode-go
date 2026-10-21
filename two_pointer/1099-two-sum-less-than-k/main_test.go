package twosumlessk

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int) int{
	"twoSumLessThanK":            twoSumLessThanK,
	"twoSumLessThanKEnumeration": twoSumLessThanKEnumeration,
}

func TestTwoSumLessThanK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"示例1", []int{34, 23, 1, 24, 75, 33, 54, 8}, 60, 58},
		{"示例2-无解", []int{10, 20, 30}, 15, -1},
		{"空数组", []int{}, 10, -1},
		{"单个元素", []int{1}, 10, -1},
		{"两个元素-可行", []int{1, 2}, 4, 3},
		{"两个元素-等于k不行", []int{1, 2}, 3, -1},
		{"全部相同", []int{5, 5, 5}, 11, 10},
		{"全部相同-无解", []int{5, 5, 5}, 10, -1},
		{"同值不能重复使用下标", []int{3, 10}, 7, -1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				nums := slices.Clone(tt.nums)
				if got := fn(nums, tt.k); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.nums, tt.k, got, tt.want)
				}
			})
		}
	}
}
