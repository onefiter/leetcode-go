package longestpalindrome

import "testing"

var solutions = map[string]func(string) string{
	"longestPalindrome":      longestPalindrome,
	"longestPalindromeBrute": longestPalindromeBrute,
	"longestPalindromeDP":    longestPalindromeDP,
}

// isPalindromeStr 判断整个字符串是否回文（测试辅助）。
func isPalindromeStr(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		wantLen int // 最长回文的长度（答案可能不唯一，故只校验长度 + 回文性）
	}{
		{"示例1", "babad", 3},
		{"示例2", "cbbd", 2},
		{"空串", "", 0},
		{"单字符", "a", 1},
		{"两字符相同", "aa", 2},
		{"两字符不同", "ab", 1},
		{"全部相同", "aaaa", 4},
		{"整体回文奇数", "racecar", 7},
		{"整体回文偶数", "abba", 4},
		{"无重复", "abcde", 1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.s)
				if len(got) != tt.wantLen {
					t.Errorf("%s(%q) = %q, 长度 %d, want 长度 %d", name, tt.s, got, len(got), tt.wantLen)
				}
				if !isPalindromeStr(got) {
					t.Errorf("%s(%q) = %q, 不是回文串", name, tt.s, got)
				}
				// 结果必须是原串的子串
				if !containsSub(tt.s, got) {
					t.Errorf("%s(%q) = %q, 不是原串的子串", name, tt.s, got)
				}
			})
		}
	}
}

// containsSub 判断 sub 是否为 s 的连续子串（测试辅助）。
func containsSub(s, sub string) bool {
	if len(sub) == 0 {
		return true
	}
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return true
		}
	}
	return false
}
