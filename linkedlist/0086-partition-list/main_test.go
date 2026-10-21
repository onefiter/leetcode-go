package partitionlist

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode, int) *ListNode{
	"partition":            partition,
	"partitionEnumeration": partitionEnumeration,
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		head []int
		x    int
		want []int
	}{
		{"示例1", []int{1, 4, 3, 2, 5, 2}, 3, []int{1, 2, 2, 4, 3, 5}},
		{"示例2", []int{2, 1}, 2, []int{1, 2}},
		{"空链表", []int{}, 0, []int{}},
		{"单元素", []int{1}, 0, []int{1}},
		{"全部小于x", []int{1, 2, 3}, 9, []int{1, 2, 3}},
		{"全部大于等于x", []int{5, 6, 7}, 5, []int{5, 6, 7}},
		{"全相同", []int{3, 3, 3}, 3, []int{3, 3, 3}},
		{"大的在前小的在尾", []int{4, 5, 1}, 3, []int{1, 4, 5}},
		{"含负数", []int{0, -1, 2, -3}, 0, []int{-1, -3, 0, 2}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := template.ListToSlice(fn(template.BuildList(tt.head), tt.x))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.head, tt.x, got, tt.want)
				}
			})
		}
	}
}
