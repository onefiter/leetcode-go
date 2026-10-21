package maxaverage

import (
	"math"
	"testing"
)

var solutions = map[string]func([]int, int) float64{
	"findMaxAverage":            findMaxAverage,
	"findMaxAveragePrefixSum":   findMaxAveragePrefixSum,
	"findMaxAverageEnumeration": findMaxAverageEnumeration,
}

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want float64
	}{
		{"示例1", []int{1, 12, -5, -6, 50, 3}, 4, 12.75},
		{"示例2-单元素", []int{5}, 1, 5.0},
		{"k等于n", []int{1, 2, 3, 4}, 4, 2.5},
		{"k为1取最大值", []int{3, -1, 7, 2}, 1, 7.0},
		{"全负数", []int{-5, -3, -8, -1}, 2, -4.0},
		{"全相同", []int{4, 4, 4, 4, 4}, 3, 4.0},
		{"最大在末尾", []int{0, 0, 0, 9, 9}, 2, 9.0},
		{"非整除平均", []int{1, 2, 2}, 3, 5.0 / 3.0},
		{"单负数", []int{-1}, 1, -1.0},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums, tt.k); math.Abs(got-tt.want) > 1e-5 {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.nums, tt.k, got, tt.want)
				}
			})
		}
	}
}
