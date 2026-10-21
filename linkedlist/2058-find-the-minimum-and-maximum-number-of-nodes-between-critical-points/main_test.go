package criticalpoints

import (
	"slices"
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode) []int{
	"nodesBetweenCriticalPoints":            nodesBetweenCriticalPoints,
	"nodesBetweenCriticalPointsEnumeration": nodesBetweenCriticalPointsEnumeration,
}

func TestNodesBetweenCriticalPoints(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want []int
	}{
		{"示例1-两个节点", []int{3, 1}, []int{-1, -1}},
		{"示例2", []int{5, 3, 1, 2, 5, 1, 2}, []int{1, 3}},
		{"示例3", []int{1, 3, 2, 2, 3, 2, 2, 2, 7}, []int{3, 3}},
		{"相等不算临界点", []int{2, 3, 3, 2}, []int{-1, -1}},
		{"只有一个临界点", []int{1, 2, 1}, []int{-1, -1}},
		{"全相同", []int{4, 4, 4, 4, 4}, []int{-1, -1}},
		{"锯齿", []int{1, 2, 1, 2, 1}, []int{1, 2}},
		{"单调递增", []int{1, 2, 3, 4, 5}, []int{-1, -1}},
		{"恰两个且相距远", []int{1, 5, 4, 3, 2, 6}, []int{3, 3}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(template.BuildList(tt.vals)); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
