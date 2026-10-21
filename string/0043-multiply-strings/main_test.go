package multiplystrings

import "testing"

func TestMultiply(t *testing.T) {
	tests := []struct {
		name string
		num1 string
		num2 string
		want string
	}{
		{"示例1", "2", "3", "6"},
		{"示例2", "123", "456", "56088"},
		{"含0-左", "0", "52", "0"},
		{"含0-右", "52", "0", "0"},
		{"两个0", "0", "0", "0"},
		{"单位乘", "9", "9", "81"},
		{"进位", "99", "99", "9801"},
		{"不等长", "123456789", "987654321", "121932631112635269"},
		{"带内部零", "100", "100", "10000"},
		{"一位乘多位", "7", "123", "861"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multiply(tt.num1, tt.num2); got != tt.want {
				t.Errorf("multiply(%q, %q) = %q, want %q", tt.num1, tt.num2, got, tt.want)
			}
		})
	}
}
