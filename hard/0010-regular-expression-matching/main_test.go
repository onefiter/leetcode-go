package regex

import "testing"

var solutions = map[string]func(string, string) bool{
	"isMatch":     isMatch,
	"isMatchMemo": isMatchMemo,
}

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name string
		s, p string
		want bool
	}{
		{"示例1", "aa", "a", false},
		{"示例2-星号重复", "aa", "a*", true},
		{"示例3-点星", "ab", ".*", true},
		{"星号匹配零次", "aab", "c*a*b", true},
		{"复杂不匹配", "mississippi", "mis*is*p*.", false},
		{"空串匹配星号模式", "", "a*b*", true},
		{"点号单字符", "abc", "a.c", true},
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
