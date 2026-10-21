package findminrotated

import (
	"testing"
)

var solutions = map[string]func([]int) int{
	"findMin":            findMin,
	"findMinEnumeration": findMinEnumeration,
}

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{3, 4, 5, 1, 2}, 1},
		{"示例2", []int{4, 5, 6, 7, 0, 1, 2}, 0},
		{"示例3-未旋转", []int{11, 13, 15, 17}, 11},
		{"单元素", []int{1}, 1},
		{"两元素-旋转", []int{2, 1}, 1},
		{"两元素-未旋转", []int{1, 2}, 1},
		{"旋转点在头部", []int{5, 1, 2, 3, 4}, 1},
		{"旋转点在尾部", []int{2, 3, 4, 5, 1}, 1},
		{"旋转点在中间", []int{4, 5, 6, 0, 1, 2, 3}, 0},
		{"含负数", []int{0, 1, 2, -3, -2, -1}, -3},
		{"全负数-旋转", []int{-1, -5, -4, -3, -2}, -5},
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

// TestFindMinAgainstEnumeration 用顺序遍历的结果核对二分解法，
// 使用互不相同的升序基数组，枚举所有旋转位置。
func TestFindMinAgainstEnumeration(t *testing.T) {
	base := []int{-7, -3, -1, 0, 2, 5, 8, 11, 14}
	n := len(base)
	for k := 0; k < n; k++ {
		rotated := append(append([]int{}, base[k:]...), base[:k]...)
		want := findMinEnumeration(rotated)
		if got := findMin(rotated); got != want {
			t.Errorf("findMin(%v) = %d, want %d", rotated, got, want)
		}
	}
}
