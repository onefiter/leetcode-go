package decodestring

import "testing"

var solutions = map[string]func(string) string{
	"decodeString":            decodeString,
	"decodeStringRecursive":   decodeStringRecursive,
	"decodeStringEnumeration": decodeStringEnumeration,
}

func TestDecodeString(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"示例1", "3[a]2[bc]", "aaabcbc"},
		{"示例2-嵌套", "3[a2[c]]", "accaccacc"},
		{"示例3", "2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"示例4", "abc3[cd]xyz", "abccdcdcdxyz"},
		{"空串", "", ""},
		{"无括号", "leetcode", "leetcode"},
		{"单个字符", "a", "a"},
		{"重复一次", "1[x]", "x"},
		{"多位数", "10[a]", "aaaaaaaaaa"},
		{"深层嵌套", "2[a2[b2[c]]]", "abccbccabccbcc"},
		{"相邻括号组", "2[a]2[b]", "aabb"},
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
