package largestnumber

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int) string{
	"largestNumber":       largestNumber,
	"largestNumberString": largestNumberString,
}

func TestLargestNumber(t *testing.T) {
	tests := []struct {
		name   string
		cost   []int
		target int
		want   string
	}{
		{"示例1", []int{4, 3, 2, 5, 6, 7, 2, 5, 5}, 9, "7772"},
		{"示例2", []int{7, 6, 5, 5, 5, 6, 8, 7, 8}, 12, "85"},
		{"示例3无解", []int{2, 4, 6, 2, 4, 6, 4, 4, 4}, 5, "0"},
		{"示例4", []int{6, 10, 15, 40, 40, 40, 40, 40, 40}, 47, "32211"},
		{"成本全相同", []int{1, 1, 1, 1, 1, 1, 1, 1, 1}, 3, "999"},
		{"恰好一个数字", []int{9, 9, 9, 9, 9, 9, 9, 9, 3}, 3, "9"},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.cost), tt.target); got != tt.want {
					t.Errorf("%s(%v, %d) = %q, want %q", name, tt.cost, tt.target, got, tt.want)
				}
			})
		}
	}
}
