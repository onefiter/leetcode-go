package twodistinct

import "testing"

var solutions = map[string]func(string) int{
	"lengthOfLongestSubstringTwoDistinct":            lengthOfLongestSubstringTwoDistinct,
	"lengthOfLongestSubstringTwoDistinctEnumeration": lengthOfLongestSubstringTwoDistinctEnumeration,
}

func TestLengthOfLongestSubstringTwoDistinct(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1", "eceba", 3},
		{"示例2", "ccaabbb", 5},
		{"空串", "", 0},
		{"单字符", "a", 1},
		{"两个字符", "ab", 2},
		{"全相同", "aaaa", 4},
		{"三种交替", "abcabcabc", 2},
		{"末尾最长", "abcbbbbcccbdddadacb", 10},
		{"整体合法", "ababab", 6},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.s); got != tt.want {
					t.Errorf("%s(%q) = %d, want %d", name, tt.s, got, tt.want)
				}
			})
		}
	}
}
