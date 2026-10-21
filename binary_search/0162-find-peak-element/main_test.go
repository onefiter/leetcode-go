package findpeak

import "testing"

var solutions = map[string]func([]int) int{
	"findPeakElement":            findPeakElement,
	"findPeakElementEnumeration": findPeakElementEnumeration,
}

// isPeak 判断下标 i 是否为峰值（两端视为 -∞）。
func isPeak(nums []int, i int) bool {
	if i < 0 || i >= len(nums) {
		return false
	}
	if i > 0 && nums[i] <= nums[i-1] {
		return false
	}
	if i < len(nums)-1 && nums[i] <= nums[i+1] {
		return false
	}
	return true
}

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{"示例1", []int{1, 2, 3, 1}},
		{"示例2-多个峰", []int{1, 2, 1, 3, 5, 6, 4}},
		{"单元素", []int{1}},
		{"两个元素-递增", []int{1, 2}},
		{"两个元素-递减", []int{2, 1}},
		{"严格递增", []int{1, 2, 3, 4, 5}},
		{"严格递减", []int{5, 4, 3, 2, 1}},
		{"锯齿", []int{1, 3, 2, 4, 3, 5, 4}},
		{"含负数与极值", []int{-2147483648, 2147483647, -2147483648}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums); !isPeak(tt.nums, got) {
					t.Errorf("%s(%v) = %d, not a peak", name, tt.nums, got)
				}
			})
		}
	}
}
