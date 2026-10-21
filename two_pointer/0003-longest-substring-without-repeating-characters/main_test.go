package longestsubstring

import "testing"

var solutions = map[string]func(string) int{
	"lengthOfLongestSubstring":      lengthOfLongestSubstring,
	"lengthOfLongestSubstringBrute": lengthOfLongestSubstringBrute,
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"示例1", "abcabcbb", 3},
		{"示例2-全相同", "bbbbb", 1},
		{"示例3", "pwwkew", 3},
		{"空串", "", 0},
		{"单字符", "a", 1},
		{"全不同", "abcdef", 6},
		{"末尾重复", "abcda", 4},
		{"含空格符号", "a b!a", 4},
		{"跳跃更新左边界", "abba", 2},
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
