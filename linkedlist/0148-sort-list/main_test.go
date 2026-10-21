package sortlist

import (
	"leetcode-go/template"
	"slices"
	"testing"
)

var solutions = map[string]func(*ListNode) *ListNode{
	"sortList":            sortList,
	"sortListTopDown":     sortListTopDown,
	"sortListEnumeration": sortListEnumeration,
}

func TestSortList(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"示例1", []int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{"示例2", []int{-1, 5, 3, 4, 0}, []int{-1, 0, 3, 4, 5}},
		{"空链表", []int{}, []int{}},
		{"单个节点", []int{1}, []int{1}},
		{"两个节点逆序", []int{2, 1}, []int{1, 2}},
		{"全部相同", []int{3, 3, 3}, []int{3, 3, 3}},
		{"已升序", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{"降序奇数长度", []int{7, 6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7}},
		{"含重复和负数", []int{0, -2, 5, -2, 3, 0, 9, 1}, []int{-2, -2, 0, 0, 1, 3, 5, 9}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 解法会原地修改链表，每次重新构造输入。
				got := template.ListToSlice(fn(template.BuildList(tt.vals)))
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}

// TestSortListAgainstEnumeration 用不同长度的伪随机序列核对各解法。
func TestSortListAgainstEnumeration(t *testing.T) {
	seed := 7
	for n := 0; n <= 33; n++ {
		vals := make([]int, n)
		for i := range vals {
			seed = (seed*1103515245 + 12345) % 2147483648
			vals[i] = seed%21 - 10
		}
		want := slices.Clone(vals)
		slices.Sort(want)
		for name, fn := range solutions {
			got := template.ListToSlice(fn(template.BuildList(vals)))
			if !slices.Equal(got, want) {
				t.Errorf("%s(%v) = %v, want %v", name, vals, got, want)
			}
		}
	}
}
