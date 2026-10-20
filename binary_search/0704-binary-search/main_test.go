package binarysearch

import "testing"

var solutions = map[string]func([]int, int) int{
	"search":          search,
	"searchHalfOpen":  searchHalfOpen,
	"searchRecursive": searchRecursive,
	"searchStdlib":    searchStdlib,
	"searchLinear":    searchLinear,
}

func TestSearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"示例1", []int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{"示例2-不存在", []int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{"首元素", []int{-1, 0, 3, 5, 9, 12}, -1, 0},
		{"尾元素", []int{-1, 0, 3, 5, 9, 12}, 12, 5},
		{"单元素-存在", []int{5}, 5, 0},
		{"单元素-不存在", []int{5}, -5, -1},
		{"两元素-取右", []int{1, 3}, 3, 1},
		{"小于所有元素", []int{2, 4, 6}, 1, -1},
		{"大于所有元素", []int{2, 4, 6}, 7, -1},
		{"空数组", []int{}, 1, -1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums, tt.target); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.nums, tt.target, got, tt.want)
				}
			})
		}
	}
}

// TestSearchAllTargets 对一个较长数组的每个元素及间隙值逐一验证，与暴力解对拍。
func TestSearchAllTargets(t *testing.T) {
	nums := make([]int, 0, 100)
	for i := range 100 {
		nums = append(nums, i*2-50) // 偶数：-50, -48, ..., 148
	}
	for target := -52; target <= 150; target++ {
		want := searchLinear(nums, target)
		for name, fn := range solutions {
			if got := fn(nums, target); got != want {
				t.Errorf("%s(target=%d) = %d, want %d", name, target, got, want)
			}
		}
	}
}
