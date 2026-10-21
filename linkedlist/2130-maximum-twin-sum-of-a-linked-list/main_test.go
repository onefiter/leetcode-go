package pairsum

import (
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode) int{
	"pairSum":            pairSum,
	"pairSumEnumeration": pairSumEnumeration,
}

func TestPairSum(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"示例1", []int{5, 4, 2, 1}, 6},
		{"示例2", []int{4, 2, 2, 3}, 7},
		{"示例3-两个节点", []int{1, 100000}, 100001},
		{"全相同", []int{3, 3, 3, 3, 3, 3}, 6},
		{"最大在头尾", []int{9, 1, 1, 9}, 18},
		{"最大在中间", []int{1, 8, 7, 1}, 15},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 解法一会原地反转后半段，每次重新构造。
				if got := fn(template.BuildList(tt.vals)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
