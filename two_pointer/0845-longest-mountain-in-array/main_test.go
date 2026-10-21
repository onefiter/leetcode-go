package longestmountain

import "testing"

var solutions = map[string]func([]int) int{
	"longestMountain":            longestMountain,
	"longestMountainDP":          longestMountainDP,
	"longestMountainEnumeration": longestMountainEnumeration,
}

func TestLongestMountain(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"示例1", []int{2, 1, 4, 7, 3, 2, 5}, 5},
		{"示例2-全相同", []int{2, 2, 2}, 0},
		{"空数组", []int{}, 0},
		{"单个元素", []int{1}, 0},
		{"两个元素", []int{1, 2}, 0},
		{"最小山脉", []int{1, 3, 2}, 3},
		{"只上坡", []int{1, 2, 3, 4}, 0},
		{"只下坡", []int{4, 3, 2, 1}, 0},
		{"山顶有平台", []int{1, 2, 2, 1}, 0},
		{"两座山共享山脚", []int{1, 3, 1, 4, 5, 2}, 4},
		{"平台后接山", []int{2, 2, 1, 4, 3, 2, 2}, 4},
		{"整个数组是山", []int{0, 1, 2, 3, 2, 1, 0}, 7},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.arr); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.arr, got, tt.want)
				}
			})
		}
	}
}
