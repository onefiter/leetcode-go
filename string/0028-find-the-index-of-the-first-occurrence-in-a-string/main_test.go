package strstr

import "testing"

var solutions = map[string]func(string, string) int{
	"strStr":    strStr,
	"strStrKMP": strStrKMP,
}

func TestStrStr(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needle   string
		want     int
	}{
		{"示例1", "sadbutsad", "sad", 0},
		{"示例2-不存在", "leetcode", "leeto", -1},
		{"needle为空", "abc", "", 0},
		{"两者都空", "", "", 0},
		{"haystack为空", "", "a", -1},
		{"末尾匹配", "hello", "llo", 2},
		{"整串相等", "abc", "abc", 0},
		{"needle更长", "a", "ab", -1},
		{"单字符命中", "abc", "b", 1},
		{"单字符不命中", "abc", "d", -1},
		{"重复前缀KMP", "aaabaaaab", "aaaab", 4},
		{"多次出现取首个", "mississippi", "issip", 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.haystack, tt.needle); got != tt.want {
					t.Errorf("%s(%q, %q) = %d, want %d", name, tt.haystack, tt.needle, got, tt.want)
				}
			})
		}
	}
}
