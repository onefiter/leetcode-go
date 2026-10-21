package minsubarraylen

import "testing"

var solutions = map[string]func(int, []int) int{
	"minSubArrayLen":             minSubArrayLen,
	"minSubArrayLenBinarySearch": minSubArrayLenBinarySearch,
	"minSubArrayLenEnumeration":  minSubArrayLenEnumeration,
}

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{"示例1", 7, []int{2, 3, 1, 2, 4, 3}, 2},
		{"示例2", 4, []int{1, 4, 4}, 1},
		{"示例3-无解", 11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
		{"空数组", 1, []int{}, 0},
		{"单个元素-够", 3, []int{5}, 1},
		{"单个元素-不够", 6, []int{5}, 0},
		{"全部相同", 6, []int{2, 2, 2, 2}, 3},
		{"需要整个数组", 15, []int{1, 2, 3, 4, 5}, 5},
		{"答案在末尾", 11, []int{1, 2, 3, 4, 5}, 3},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.target, tt.nums); got != tt.want {
					t.Errorf("%s(%d, %v) = %d, want %d", name, tt.target, tt.nums, got, tt.want)
				}
			})
		}
	}
}
