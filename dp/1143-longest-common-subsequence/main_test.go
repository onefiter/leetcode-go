package lcs

import "testing"

var solutions = map[string]func(string, string) int{
	"longestCommonSubsequence":   longestCommonSubsequence,
	"longestCommonSubsequence2D": longestCommonSubsequence2D,
}

func TestLongestCommonSubsequence(t *testing.T) {
	tests := []struct {
		name         string
		text1, text2 string
		want         int
	}{
		{"示例1", "abcde", "ace", 3},
		{"示例2", "abc", "abc", 3},
		{"示例3", "abc", "def", 0},
		{"单字符相同", "a", "a", 1},
		{"顺序颠倒", "abc", "cba", 1},
		{"长度不等", "bl", "yby", 1},
		{"交错匹配", "ezupkr", "ubmrapg", 2},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.text1, tt.text2); got != tt.want {
					t.Errorf("%s(%q, %q) = %d, want %d", name, tt.text1, tt.text2, got, tt.want)
				}
			})
		}
	}
}
