package basiccalculator

import "testing"

var solutions = map[string]func(string) int{
	"calculate":            calculate,
	"calculateStack":       calculateStack,
	"calculateEnumeration": calculateEnumeration,
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1", "1 + 1", 2},
		{"示例2", " 2-1 + 2 ", 3},
		{"示例3", "(1+(4+5+2)-3)+(6+8)", 23},
		{"单个数字", "7", 7},
		{"多位数", "123", 123},
		{"一元负号", "-2+1", -1},
		{"括号前负号", "-(2+3)", -5},
		{"嵌套负号", "1-(-(2-3))", 0},
		{"括号内负号", "2-(5-6)", 3},
		{"多层括号", "((((10))))", 10},
		{"大量空格", "  30 -  ( 4 + 6 ) ", 20},
		{"连续减法", "10-2-3", 5},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s); got != tt.want {
					t.Errorf("%s(%q) = %d, want %d", name, tt.s, got, tt.want)
				}
			})
		}
	}
}
