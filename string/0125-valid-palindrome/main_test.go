package validpalindrome

import "testing"

var solutions = map[string]func(string) bool{
	"isPalindrome":       isPalindrome,
	"isPalindromeFilter": isPalindromeFilter,
}

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"示例1", "A man, a plan, a canal: Panama", true},
		{"示例2", "race a car", false},
		{"空串", "", true},
		{"只有空格和符号", " ", true},
		{"只有符号", ".,", true},
		{"单字符", "a", true},
		{"单数字", "0", true},
		{"数字回文", "12321", true},
		{"数字非回文", "12345", false},
		{"字母数字混合", "0P", false},
		{"大小写忽略", "Aa", true},
		{"含下划线非字母数字", "a_a", true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s); got != tt.want {
					t.Errorf("%s(%q) = %v, want %v", name, tt.s, got, tt.want)
				}
			})
		}
	}
}
