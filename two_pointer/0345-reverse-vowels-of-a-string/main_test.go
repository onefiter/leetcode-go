package reversevowels

import "testing"

var solutions = map[string]func(string) string{
	"reverseVowels":            reverseVowels,
	"reverseVowelsEnumeration": reverseVowelsEnumeration,
}

func TestReverseVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"示例1", "IceCreAm", "AceCreIm"},
		{"示例2", "leetcode", "leotcede"},
		{"空串", "", ""},
		{"单个元音", "a", "a"},
		{"单个辅音", "b", "b"},
		{"没有元音", "bcdfg", "bcdfg"},
		{"全是元音", "aeiou", "uoiea"},
		{"大小写混合", "aA", "Aa"},
		{"含空格标点", "a.b,E!", "E.b,a!"},
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
