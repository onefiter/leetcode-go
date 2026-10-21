package searchinsertposition

import "testing"

var solutions = map[string]func([]int, int) int{
	"searchInsert":         searchInsert,
	"searchInsertBinary":   searchInsertBinary,
	"searchInsertHalfOpen": searchInsertHalfOpen,
	"searchInsertStdlib":   searchInsertStdlib,
}

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{"示例1-找到", []int{1, 3, 5, 6}, 5, 2},
		{"示例2-插中间", []int{1, 3, 5, 6}, 2, 1},
		{"示例3-插末尾", []int{1, 3, 5, 6}, 7, 4},
		{"示例4-插开头", []int{1, 3, 5, 6}, 0, 0},
		{"单元素-相等", []int{1}, 1, 0},
		{"单元素-插入前面", []int{1}, 0, 0},
		{"单元素-插入后面", []int{1}, 2, 1},
		{"空数组", []int{}, 5, 0},
		{"负数目标", []int{-5, 0, 3, 9}, -6, 0},
		{"负数目标-插中间", []int{-5, 0, 3, 9}, -1, 1},
		{"全大于目标", []int{2, 4, 6}, 1, 0},
		{"全小于目标", []int{2, 4, 6}, 7, 3},
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

// TestSearchInsertAgainstLinear 用顺序遍历的结果核对二分类解法的正确性，
// 覆盖数组元素、元素之间、以及比所有元素都小/都大的各种 target。
func TestSearchInsertAgainstLinear(t *testing.T) {
	nums := make([]int, 0, 100)
	for i := range 100 {
		nums = append(nums, i*3-75) // -75, -72, ..., 222
	}
	for target := -80; target <= 230; target++ {
		want := searchInsert(nums, target)
		for name, fn := range solutions {
			if got := fn(nums, target); got != want {
				t.Errorf("%s(target=%d) = %d, want %d", name, target, got, want)
			}
		}
	}
}
