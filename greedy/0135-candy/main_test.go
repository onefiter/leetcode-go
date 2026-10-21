package candy

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"candy":      candy,
	"candyBrute": candyBrute,
}

func TestCandy(t *testing.T) {
	tests := []struct {
		name    string
		ratings []int
		want    int
	}{
		{"示例1", []int{1, 0, 2}, 5},
		{"示例2", []int{1, 2, 2}, 4},
		{"单个孩子", []int{7}, 1},
		{"严格递增", []int{1, 2, 3, 4}, 10},
		{"严格递减", []int{4, 3, 2, 1}, 10},
		{"全部相同", []int{3, 3, 3}, 3},
		{"山峰形", []int{1, 3, 4, 5, 2}, 11},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(slices.Clone(tt.ratings))
				if got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.ratings, got, tt.want)
				}
			})
		}
	}
}
