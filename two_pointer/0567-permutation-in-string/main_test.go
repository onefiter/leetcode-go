package checkinclusion

import "testing"

var solutions = map[string]func(string, string) bool{
	"checkInclusion":             checkInclusion,
	"checkInclusionArrayCompare": checkInclusionArrayCompare,
	"checkInclusionEnumeration":  checkInclusionEnumeration,
}

func TestCheckInclusion(t *testing.T) {
	tests := []struct {
		name   string
		s1, s2 string
		want   bool
	}{
		{"示例1", "ab", "eidbaooo", true},
		{"示例2", "ab", "eidboaoo", false},
		{"s1比s2长", "abc", "ab", false},
		{"完全相同", "abc", "abc", true},
		{"整串是排列", "abc", "cba", true},
		{"单字符命中", "a", "a", true},
		{"单字符不命中", "a", "b", false},
		{"全相同字母", "aaa", "aaaa", true},
		{"个数不够", "aaa", "aabaa", false},
		{"在末尾", "abc", "xxxxcab", true},
		{"在开头", "abc", "bcaxxxx", true},
		{"重复字母", "aab", "abab", true},
		{"进出同字母", "ab", "aaab", true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s1, tt.s2); got != tt.want {
					t.Errorf("%s(%q, %q) = %v, want %v", name, tt.s1, tt.s2, got, tt.want)
				}
			})
		}
	}
}
