package productlessk

import "testing"

var solutions = map[string]func([]int, int) int{
	"numSubarrayProductLessThanK":            numSubarrayProductLessThanK,
	"numSubarrayProductLessThanKEnumeration": numSubarrayProductLessThanKEnumeration,
}

func TestNumSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"示例1", []int{10, 5, 2, 6}, 100, 8},
		{"示例2-k为0", []int{1, 2, 3}, 0, 0},
		{"k为1", []int{1, 1, 1}, 1, 0},
		{"空数组", []int{}, 10, 0},
		{"单元素-满足", []int{5}, 6, 1},
		{"单元素-不满足", []int{5}, 5, 0},
		{"全为1", []int{1, 1, 1}, 2, 6},
		{"全部超过k", []int{10, 20, 30}, 5, 0},
		{"k很大全满足", []int{1, 2, 3, 4}, 1000, 10},
		{"中间有大数隔断", []int{1, 2, 100, 3}, 10, 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums, tt.k); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.nums, tt.k, got, tt.want)
				}
			})
		}
	}
}
