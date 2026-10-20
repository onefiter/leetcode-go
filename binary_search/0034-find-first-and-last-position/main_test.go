package findfirstlast

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int) []int{
	"searchRange":          searchRange,
	"searchRangeTwoBinary": searchRangeTwoBinary,
	"searchRangeStdlib":    searchRangeStdlib,
	"searchRangeLinear":    searchRangeLinear,
}

func TestSearchRange(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"示例1", []int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{"示例2-不存在", []int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{"示例3-空数组", []int{}, 0, []int{-1, -1}},
		{"单个元素-命中", []int{1}, 1, []int{0, 0}},
		{"单个元素-不命中", []int{1}, 2, []int{-1, -1}},
		{"全部相同-命中", []int{2, 2, 2, 2}, 2, []int{0, 3}},
		{"目标在开头", []int{1, 1, 2, 3}, 1, []int{0, 1}},
		{"目标在末尾", []int{1, 2, 3, 3}, 3, []int{2, 3}},
		{"比最小还小", []int{2, 4, 6}, 1, []int{-1, -1}},
		{"比最大还大", []int{2, 4, 6}, 9, []int{-1, -1}},
		{"含负数", []int{-3, -3, -1, 0, 0, 0, 5}, 0, []int{3, 5}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums, tt.target); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.nums, tt.target, got, tt.want)
				}
			})
		}
	}
}

// TestSearchRangeAgainstLinear 用顺序遍历的结果核对所有二分解法，
// 随机构造的非递减数组覆盖重复元素、缺口、边界等情况。
func TestSearchRangeAgainstLinear(t *testing.T) {
	// 构造一个有大量重复的非递减数组：每个值 v 出现 v 次左右。
	nums := []int{}
	for v := 0; v <= 20; v++ {
		for k := 0; k < v%4; k++ {
			nums = append(nums, v)
		}
	}
	for target := -2; target <= 22; target++ {
		want := searchRangeLinear(nums, target)
		for name, fn := range solutions {
			if got := fn(nums, target); !slices.Equal(got, want) {
				t.Errorf("%s(target=%d) = %v, want %v", name, target, got, want)
			}
		}
	}
}
