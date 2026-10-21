package addtwonumbers

import (
	"leetcode-go/template"
	"slices"
	"testing"
)

var solutions = map[string]func(*ListNode, *ListNode) *ListNode{
	"addTwoNumbers":            addTwoNumbers,
	"addTwoNumbersReverse":     addTwoNumbersReverse,
	"addTwoNumbersEnumeration": addTwoNumbersEnumeration,
}

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name   string
		l1, l2 []int
		want   []int
	}{
		{"示例1", []int{7, 2, 4, 3}, []int{5, 6, 4}, []int{7, 8, 0, 7}},
		{"示例2", []int{2, 4, 3}, []int{5, 6, 4}, []int{8, 0, 7}},
		{"示例3-两个0", []int{0}, []int{0}, []int{0}},
		{"单位数无进位", []int{3}, []int{4}, []int{7}},
		{"单位数有进位", []int{5}, []int{5}, []int{1, 0}},
		{"连续进位到最高位", []int{9, 9, 9}, []int{1}, []int{1, 0, 0, 0}},
		{"长度差异大", []int{1}, []int{9, 9, 9, 9, 9}, []int{1, 0, 0, 0, 0, 0}},
		{"加0", []int{1, 2, 3}, []int{0}, []int{1, 2, 3}},
		{"全部为9", []int{9, 9}, []int{9, 9}, []int{1, 9, 8}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 反转解法会原地修改输入，每次重新构造输入。
				got := template.ListToSlice(fn(template.BuildList(tt.l1), template.BuildList(tt.l2)))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v, %v) = %v, want %v", name, tt.l1, tt.l2, got, tt.want)
				}
			})
		}
	}
}
