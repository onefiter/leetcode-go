package palindromelist

import (
	"leetcode-go/template"
	"slices"
	"testing"
)

var solutions = map[string]func(*ListNode) bool{
	"isPalindrome":            isPalindrome,
	"isPalindromeEnumeration": isPalindromeEnumeration,
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want bool
	}{
		{"示例1-偶数回文", []int{1, 2, 2, 1}, true},
		{"示例2-非回文", []int{1, 2}, false},
		{"空链表", []int{}, true},
		{"单节点", []int{1}, true},
		{"奇数回文", []int{1, 2, 3, 2, 1}, true},
		{"奇数非回文", []int{1, 2, 3, 1, 1}, false},
		{"全部相同", []int{5, 5, 5, 5}, true},
		{"两个相同", []int{9, 9}, true},
		{"仅中间不同", []int{1, 2, 3, 4, 2, 1}, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				head := template.BuildList(tt.vals)
				if got := fn(head); got != tt.want {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
				// 判断后链表应保持原样。
				if after := template.ListToSlice(head); !slices.Equal(after, tt.vals) {
					t.Errorf("%s 修改了输入链表：%v -> %v", name, tt.vals, after)
				}
			})
		}
	}
}
