package issubsequence

import "testing"

var solutions = map[string]func(string, string) bool{
	"isSubsequence":            isSubsequence,
	"isSubsequenceDP":          isSubsequenceDP,
	"isSubsequenceEnumeration": isSubsequenceEnumeration,
}

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		name string
		s, t string
		want bool
	}{
		{"示例1", "abc", "ahbgdc", true},
		{"示例2", "axc", "ahbgdc", false},
		{"s为空", "", "abc", true},
		{"都为空", "", "", true},
		{"t为空", "a", "", false},
		{"完全相同", "abc", "abc", true},
		{"s比t长", "abcd", "abc", false},
		{"顺序不对", "ba", "ab", false},
		{"全相同字符-够", "aaa", "aaaa", true},
		{"全相同字符-不够", "aaaa", "aaa", false},
		{"单字符命中", "c", "abc", true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s, tt.t); got != tt.want {
					t.Errorf("%s(%q, %q) = %v, want %v", name, tt.s, tt.t, got, tt.want)
				}
			})
		}
	}
}
