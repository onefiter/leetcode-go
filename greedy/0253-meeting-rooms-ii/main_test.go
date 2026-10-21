package meetingrooms2

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
	"minMeetingRooms":     minMeetingRooms,
	"minMeetingRoomsHeap": minMeetingRoomsHeap,
}

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      int
	}{
		{"示例1", [][]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		{"示例2", [][]int{{7, 10}, {2, 4}}, 1},
		{"空输入", [][]int{}, 0},
		{"首尾相接可复用", [][]int{{1, 5}, {5, 10}, {10, 15}}, 1},
		{"全部重叠", [][]int{{1, 10}, {2, 9}, {3, 8}}, 3},
		{"乱序部分重叠", [][]int{{9, 10}, {4, 9}, {4, 17}}, 2},
		{"嵌套与链式", [][]int{{1, 4}, {2, 5}, {4, 6}, {5, 7}}, 2},
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
