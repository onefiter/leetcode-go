package minwindow

import "testing"

var solutions = map[string]func(string, string) string{
	"minWindow":      minWindow,
	"minWindowBrute": minWindowBrute,
}

func TestMinWindow(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want string
	}{
		{"示例1", "ADOBECODEBANC", "ABC", "BANC"},
		{"示例2-单字符", "a", "a", "a"},
		{"示例3-无解", "a", "aa", ""},
		{"空s", "", "a", ""},
		{"空t", "abc", "", ""},
		{"t比s长", "ab", "abc", ""},
		{"全相同覆盖", "aaaa", "aa", "aa"},
		{"整串即答案", "abc", "cba", "abc"},
		{"重复字符需求", "aabbccbb", "abc", "abbc"},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s, tt.t); got != tt.want {
					t.Errorf("%s(%q, %q) = %q, want %q", name, tt.s, tt.t, got, tt.want)
				}
			})
		}
	}
}
