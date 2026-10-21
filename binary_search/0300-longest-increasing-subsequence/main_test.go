package lis

import "testing"

var solutions = map[string]func([]int) int{
	"lengthOfLIS":            lengthOfLIS,
	"lengthOfLISDP":          lengthOfLISDP,
	"lengthOfLISEnumeration": lengthOfLISEnumeration,
}

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{"示例2", []int{0, 1, 0, 3, 2, 3}, 4},
		{"示例3-全相同", []int{7, 7, 7, 7, 7, 7, 7}, 1},
		{"空数组", []int{}, 0},
		{"单元素", []int{5}, 1},
		{"严格递增", []int{1, 2, 3, 4, 5}, 5},
		{"严格递减", []int{5, 4, 3, 2, 1}, 1},
		{"含负数", []int{-3, -1, -2, 0, -5, 4}, 4},
		{"重复中有递增", []int{1, 3, 3, 3, 2, 4}, 3},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}

// TestLengthOfLISAgainstEnumeration 用伪随机小数组和暴力枚举结果比对。
func TestLengthOfLISAgainstEnumeration(t *testing.T) {
	seed := 12345
	next := func() int {
		seed = (seed*1103515245 + 12345) & 0x7fffffff
		return seed%7 - 3 // [-3, 3]，制造大量重复
	}
	for round := 0; round < 200; round++ {
		n := round % 12
		nums := make([]int, n)
		for i := range nums {
			nums[i] = next()
		}
		want := lengthOfLISEnumeration(nums)
		for name, fn := range solutions {
			if got := fn(nums); got != want {
				t.Errorf("%s(%v) = %d, want %d", name, nums, got, want)
			}
		}
	}
}
