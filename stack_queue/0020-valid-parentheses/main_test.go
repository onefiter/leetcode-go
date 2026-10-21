package validparentheses

import "testing"

var solutions = map[string]func(string) bool{
	"isValid":            isValid,
	"isValidEnumeration": isValidEnumeration,
}

func TestIsValid(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"示例1", "()", true},
		{"示例2", "()[]{}", true},
		{"示例3", "(]", false},
		{"示例4-嵌套", "([])", true},
		{"空串", "", true},
		{"单个左括号", "(", false},
		{"单个右括号", "]", false},
		{"顺序交错", "([)]", false},
		{"深层嵌套", "{[()()]}", true},
		{"只有左括号", "((((", false},
		{"右括号开头", "){}(", false},
		{"末尾多余左括号", "()[]{", false},
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
