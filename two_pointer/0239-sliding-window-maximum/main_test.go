package slidingmax

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int) []int{
	"maxSlidingWindow":            maxSlidingWindow,
	"maxSlidingWindowHeap":        maxSlidingWindowHeap,
	"maxSlidingWindowEnumeration": maxSlidingWindowEnumeration,
}

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want []int
	}{
		{"示例1", []int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{"示例2-单元素", []int{1}, 1, []int{1}},
		{"空数组", []int{}, 1, []int{}},
		{"k=1", []int{4, -2, 7}, 1, []int{4, -2, 7}},
		{"k=n", []int{4, -2, 7, 1}, 4, []int{7}},
		{"全相同", []int{2, 2, 2, 2}, 2, []int{2, 2, 2}},
		{"严格递减", []int{9, 7, 5, 3, 1}, 2, []int{9, 7, 5, 3}},
		{"严格递增", []int{1, 3, 5, 7, 9}, 3, []int{5, 7, 9}},
		{"全负数", []int{-5, -1, -3, -4}, 2, []int{-1, -1, -3}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.nums), tt.k); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.nums, tt.k, got, tt.want)
				}
			})
		}
	}
}

// TestMaxSlidingWindowAgainstEnumeration 用伪随机数组和多种 k 核对各解法。
func TestMaxSlidingWindowAgainstEnumeration(t *testing.T) {
	nums := make([]int, 200)
	x := 42
	for i := range nums {
		x = (x*1103515245 + 12345) & 0x7fffffff
		nums[i] = x%21 - 10
	}
	for k := 1; k <= 25; k++ {
		want := maxSlidingWindowEnumeration(nums, k)
		for name, fn := range solutions {
			if got := fn(slices.Clone(nums), k); !slices.Equal(got, want) {
				t.Errorf("%s(k=%d) 与暴力结果不一致", name, k)
			}
		}
	}
}
