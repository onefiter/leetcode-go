package removeduplicateletters

import (
	"math/rand"
	"testing"
)

var solutions = map[string]func(string) string{
	"removeDuplicateLetters":            removeDuplicateLetters,
	"removeDuplicateLettersEnumeration": removeDuplicateLettersEnumeration,
}

func TestRemoveDuplicateLetters(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"示例1", "bcabc", "abc"},
		{"示例2", "cbacdcbc", "acdb"},
		{"单个字符", "a", "a"},
		{"全相同", "aaaa", "a"},
		{"无重复", "dcba", "dcba"},
		{"已有序", "abcabc", "abc"},
		{"后面不再出现不能弹", "bab", "ab"},
		{"栈顶不能弹", "cdadabcc", "adbc"},
		{"较长", "ecbacba", "eacb"},
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

// TestAgainstEnumeration 随机字符串对拍。
func TestAgainstEnumeration(t *testing.T) {
	r := rand.New(rand.NewSource(316))
	for k := 0; k < 500; k++ {
		n := 1 + r.Intn(20)
		b := make([]byte, n)
		for i := range b {
			b[i] = byte('a' + r.Intn(5))
		}
		s := string(b)
		want := removeDuplicateLettersEnumeration(s)
		if got := removeDuplicateLetters(s); got != want {
			t.Fatalf("removeDuplicateLetters(%q) = %q, want %q", s, got, want)
		}
	}
}
