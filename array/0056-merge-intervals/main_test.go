package mergeintervals

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		{[][]int{}, [][]int{}},
		{[][]int{{1, 4}}, [][]int{{1, 4}}},
		// 重叠
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		// 端点相接视为重叠
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		// 完全包含
		{[][]int{{1, 10}, {2, 3}, {4, 5}}, [][]int{{1, 10}}},
		// 不重叠且乱序
		{[][]int{{3, 4}, {1, 2}}, [][]int{{1, 2}, {3, 4}}},
		// 全部重叠
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
	}
	for _, tt := range tests {
		got := merge(tt.intervals)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("merge(%v) = %v, want %v", tt.intervals, got, tt.want)
		}
	}
}
