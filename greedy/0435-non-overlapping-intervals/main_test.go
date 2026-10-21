package nonoverlapping

import (
	"slices"
	"testing"
)

// cloneIntervals 深拷贝二维切片，避免解法原地排序互相影响。
func cloneIntervals(src [][]int) [][]int {
	out := make([][]int, len(src))
	for i, iv := range src {
		out[i] = slices.Clone(iv)
	}
	return out
}

var solutions = map[string]func([][]int) int{
	"eraseOverlapIntervals":   eraseOverlapIntervals,
	"eraseOverlapIntervalsDP": eraseOverlapIntervalsDP,
}

func TestEraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{"示例1", [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{"示例2", [][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{"示例3", [][]int{{1, 2}, {2, 3}}, 0},
		{"单个区间", [][]int{{5, 9}}, 0},
		{"长区间覆盖多个短区间", [][]int{{1, 100}, {2, 3}, {4, 5}, {6, 7}}, 1},
		{"负数坐标", [][]int{{-5, -1}, {-3, 2}, {0, 4}, {3, 6}}, 2},
		{"极端值", [][]int{{-50000, 50000}, {-50000, 0}, {0, 50000}}, 1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(cloneIntervals(tt.intervals))
				if got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.intervals, got, tt.want)
				}
			})
		}
	}
}
