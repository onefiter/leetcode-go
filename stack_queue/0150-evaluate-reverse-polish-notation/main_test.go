package evalrpn

import "testing"

var solutions = map[string]func([]string) int{
	"evalRPN":      evalRPN,
	"evalRPNArray": evalRPNArray,
}

func TestEvalRPN(t *testing.T) {
	tests := []struct {
		name   string
		tokens []string
		want   int
	}{
		{"示例1", []string{"2", "1", "+", "3", "*"}, 9},
		{"示例2", []string{"4", "13", "5", "/", "+"}, 6},
		{"示例3", []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, 22},
		{"单个数字", []string{"42"}, 42},
		{"单个负数", []string{"-7"}, -7},
		{"减法顺序", []string{"3", "5", "-"}, -2},
		{"除法顺序", []string{"8", "2", "/"}, 4},
		{"负数除法向零截断", []string{"-7", "2", "/"}, -3},
		{"乘以零", []string{"5", "0", "*"}, 0},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.tokens); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.tokens, got, tt.want)
				}
			})
		}
	}
}
