package removeouterparens

import "testing"

var solutions = map[string]func(string) string{
	"removeOuterParentheses":            removeOuterParentheses,
	"removeOuterParenthesesEnumeration": removeOuterParenthesesEnumeration,
}

func TestRemoveOuterParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"示例1", "(()())(())", "()()()"},
		{"示例2", "(()())(())(()(()))", "()()()()(())"},
		{"示例3", "()()", ""},
		{"空串", "", ""},
		{"单个原语", "()", ""},
		{"深层嵌套", "((()))", "(())"},
		{"并列加嵌套", "(()())", "()()"},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s); got != tt.want {
					t.Errorf("%s(%q) = %q, want %q", name, tt.s, got, tt.want)
				}
			})
		}
	}
}
