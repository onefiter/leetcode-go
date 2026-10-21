package minarrows

import (
	"math"
	"slices"
	"testing"
)

// clonePoints 深拷贝二维切片，避免解法原地排序互相影响。
func clonePoints(src [][]int) [][]int {
	out := make([][]int, len(src))
	for i, p := range src {
		out[i] = slices.Clone(p)
	}
	return out
}

var solutions = map[string]func([][]int) int{
	"findMinArrowShots":        findMinArrowShots,
	"findMinArrowShotsByStart": findMinArrowShotsByStart,
}

func TestFindMinArrowShots(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{"示例1", [][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}, 2},
		{"示例2", [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}, 4},
		{"示例3端点相接", [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 2},
		{"单个气球", [][]int{{1, 2}}, 1},
		{"大区间包含小区间", [][]int{{1, 10}, {2, 3}, {4, 5}}, 2},
		{"int32极端值防溢出", [][]int{{math.MinInt32, math.MaxInt32}, {math.MaxInt32, math.MaxInt32}, {math.MinInt32, math.MinInt32}}, 2},
		{"交集需收缩", [][]int{{1, 10}, {2, 3}, {3, 4}, {9, 12}}, 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(clonePoints(tt.points))
				if got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.points, got, tt.want)
				}
			})
		}
	}
}
