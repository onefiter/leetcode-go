package reversebetween

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode, int, int) *ListNode{
	"reverseBetween":            reverseBetween,
	"reverseBetweenEnumeration": reverseBetweenEnumeration,
}

func TestReverseBetween(t *testing.T) {
	tests := []struct {
		name        string
		head        []int
		left, right int
		want        []int
	}{
		{"示例1", []int{1, 2, 3, 4, 5}, 2, 4, []int{1, 4, 3, 2, 5}},
		{"示例2-单元素", []int{5}, 1, 1, []int{5}},
		{"left等于right", []int{1, 2, 3}, 2, 2, []int{1, 2, 3}},
		{"从表头开始", []int{1, 2, 3, 4}, 1, 3, []int{3, 2, 1, 4}},
		{"到表尾结束", []int{1, 2, 3, 4}, 2, 4, []int{1, 4, 3, 2}},
		{"整条反转", []int{1, 2, 3, 4, 5}, 1, 5, []int{5, 4, 3, 2, 1}},
		{"两个元素", []int{3, 5}, 1, 2, []int{5, 3}},
		{"全相同", []int{6, 6, 6}, 1, 3, []int{6, 6, 6}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := template.ListToSlice(fn(template.BuildList(tt.head), tt.left, tt.right))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %d, %d) = %v, want %v", name, tt.head, tt.left, tt.right, got, tt.want)
				}
			})
		}
	}
}
