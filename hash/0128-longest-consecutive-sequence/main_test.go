package longestconsecutive

import "testing"

var solutions = map[string]func([]int) int{
	"longestConsecutive":     longestConsecutive,
	"longestConsecutiveSort": longestConsecutiveSort,
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{100, 4, 200, 1, 3, 2}, 4},
		{"示例2", []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},
		{"空数组", []int{}, 0},
		{"单元素", []int{5}, 1},
		{"全相同", []int{1, 1, 1, 1}, 1},
		{"无连续", []int{10, 20, 30}, 1},
		{"含负数", []int{-3, -2, -1, 0, 1}, 5},
		{"有重复的连续", []int{1, 2, 2, 3}, 3},
		{"两段取较长", []int{1, 2, 3, 10, 11}, 3},
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
