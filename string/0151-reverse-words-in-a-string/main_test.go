package reversewords

import "testing"

var solutions = map[string]func(string) string{
	"reverseWords":       reverseWords,
	"reverseWordsManual": reverseWordsManual,
}

func TestReverseWords(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"示例1", "the sky is blue", "blue is sky the"},
		{"示例2-前后空格", "  hello world  ", "world hello"},
		{"示例3-中间多空格", "a good   example", "example good a"},
		{"单个单词", "hello", "hello"},
		{"单个单词带空格", "  hello  ", "hello"},
		{"全空格", "     ", ""},
		{"空串", "", ""},
		{"两个单词", "Bob    Loves  Alice", "Alice Loves Bob"},
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
