package largestrectangle

import (
	"math/rand"
	"testing"
)

var solutions = map[string]func([]int) int{
	"largestRectangleArea":            largestRectangleArea,
	"largestRectangleAreaEnumeration": largestRectangleAreaEnumeration,
}

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		name    string
		heights []int
		want    int
	}{
		{"示例1", []int{2, 1, 5, 6, 2, 3}, 10},
		{"示例2", []int{2, 4}, 4},
		{"空数组", []int{}, 0},
		{"单个柱子", []int{5}, 5},
		{"单个零", []int{0}, 0},
		{"全相同", []int{3, 3, 3, 3}, 12},
		{"严格递增", []int{1, 2, 3, 4, 5}, 9},
		{"严格递减", []int{5, 4, 3, 2, 1}, 9},
		{"中间有零", []int{2, 0, 2}, 2},
		{"山峰", []int{1, 3, 5, 3, 1}, 9},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.heights); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.heights, got, tt.want)
				}
			})
		}
	}
}

// TestAgainstEnumeration 随机数组对拍：单调栈结果应与暴力一致。
func TestAgainstEnumeration(t *testing.T) {
	r := rand.New(rand.NewSource(84))
	for k := 0; k < 500; k++ {
		n := r.Intn(30)
		h := make([]int, n)
		for i := range h {
			h[i] = r.Intn(10)
		}
		want := largestRectangleAreaEnumeration(h)
		if got := largestRectangleArea(h); got != want {
			t.Fatalf("largestRectangleArea(%v) = %d, want %d", h, got, want)
		}
	}
}
