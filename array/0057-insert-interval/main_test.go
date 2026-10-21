package insertinterval

import (
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	tests := []struct {
		intervals   [][]int
		newInterval []int
		want        [][]int
	}{
		// 空列表
		{[][]int{}, []int{5, 7}, [][]int{{5, 7}}},
		// 插入到最前且不重叠
		{[][]int{{3, 5}, {8, 10}}, []int{1, 2}, [][]int{{1, 2}, {3, 5}, {8, 10}}},
		// 插入到最后且不重叠
		{[][]int{{1, 2}, {3, 5}}, []int{8, 10}, [][]int{{1, 2}, {3, 5}, {8, 10}}},
		// 与单个区间重叠
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		// 跨多个区间重叠
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		// 新区间包含所有
		{[][]int{{2, 3}, {5, 7}}, []int{1, 10}, [][]int{{1, 10}}},
		// 端点相接视为重叠
		{[][]int{{1, 5}}, []int{5, 8}, [][]int{{1, 8}}},
	}
	for _, tt := range tests {
		got := insert(tt.intervals, tt.newInterval)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("insert(%v, %v) = %v, want %v", tt.intervals, tt.newInterval, got, tt.want)
		}
	}
}
