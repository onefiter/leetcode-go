package atoi

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1", "42", 42},
		{"示例2-前导空格负号", "   -42", -42},
		{"示例3-尾部文字", "4193 with words", 4193},
		{"示例4-字母开头", "words and 987", 0},
		{"下溢裁剪", "-91283472332", -2147483648},
		{"上溢裁剪", "91283472332", 2147483647},
		{"空串", "", 0},
		{"只有空格", "   ", 0},
		{"只有符号", "+", 0},
		{"正号", "+42", 42},
		{"前导零", "0032", 32},
		{"中间有空格", "3.14159", 3},
		{"正负号连写", "+-12", 0},
		{"恰好上界", "2147483647", 2147483647},
		{"恰好下界", "-2147483648", -2147483648},
		{"小数点", "3.14", 3},
		{"负零", "-0", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := myAtoi(tt.s); got != tt.want {
				t.Errorf("myAtoi(%q) = %d, want %d", tt.s, got, tt.want)
			}
		})
	}
}
