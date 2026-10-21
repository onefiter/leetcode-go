package addtwonumbers

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode, *ListNode) *ListNode{
	"addTwoNumbers":            addTwoNumbers,
	"addTwoNumbersRecursive":   addTwoNumbersRecursive,
	"addTwoNumbersEnumeration": addTwoNumbersEnumeration,
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name   string
		l1, l2 []int
		want   []int
	}{
		{"示例1", []int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}},
		{"示例2-都为0", []int{0}, []int{0}, []int{0}},
		{"示例3-连续进位", []int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}, []int{8, 9, 9, 9, 0, 0, 0, 1}},
		{"单元素无进位", []int{1}, []int{2}, []int{3}},
		{"单元素有进位", []int{5}, []int{5}, []int{0, 1}},
		{"长短不一", []int{1, 8}, []int{0}, []int{1, 8}},
		{"短加长", []int{0}, []int{3, 7, 1}, []int{3, 7, 1}},
		{"最高位进位", []int{9, 9}, []int{1}, []int{0, 0, 1}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := template.ListToSlice(fn(template.BuildList(tt.l1), template.BuildList(tt.l2)))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %v) = %v, want %v", name, tt.l1, tt.l2, got, tt.want)
				}
			})
		}
	}
}
