package getdecimalvalue

import (
	"testing"

	"leetcode-go/template"
)

var solutions = map[string]func(*ListNode) int{
	"getDecimalValue":            getDecimalValue,
	"getDecimalValueEnumeration": getDecimalValueEnumeration,
}

func TestGetDecimalValue(t *testing.T) {
	thirty := make([]int, 30)
	thirty[0] = 1
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"示例1", []int{1, 0, 1}, 5},
		{"示例2-单个0", []int{0}, 0},
		{"单个1", []int{1}, 1},
		{"前导零", []int{0, 0, 1, 1}, 3},
		{"全1", []int{1, 1, 1, 1}, 15},
		{"全0", []int{0, 0, 0}, 0},
		{"30位-最高位为1", thirty, 1 << 29},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(template.BuildList(tt.vals)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
