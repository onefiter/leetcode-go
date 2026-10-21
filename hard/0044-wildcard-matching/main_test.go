package wildcard

import "testing"

var solutions = map[string]func(string, string) bool{
	"isMatch":   isMatch,
	"isMatchDP": isMatchDP,
}

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name string
		s, p string
		want bool
	}{
		{"示例1", "aa", "a", false},
		{"示例2-星号匹配全部", "aa", "*", true},
		{"示例3-问号不匹配", "cb", "?a", false},
		{"多个星号", "adceb", "*a*b", true},
		{"需要回溯", "acdcb", "a*c?b", false},
		{"空串与空模式", "", "", true},
		{"空串与全星号", "", "***", true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s, tt.p); got != tt.want {
					t.Errorf("%s(%q, %q) = %v, want %v", name, tt.s, tt.p, got, tt.want)
				}
			})
		}
	}
}
