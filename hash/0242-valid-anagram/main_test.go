package validanagram

import "testing"

var solutions = map[string]func(string, string) bool{
	"isAnagram":      isAnagram,
	"isAnagramArray": isAnagramArray,
	"isAnagramSort":  isAnagramSort,
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		s    string
		t    string
		want bool
	}{
		{"示例1", "anagram", "nagaram", true},
		{"示例2", "rat", "car", false},
		{"长度不同", "ab", "abc", false},
		{"都为空", "", "", true},
		{"单字符相同", "a", "a", true},
		{"单字符不同", "a", "b", false},
		{"相同字母不同频次", "aacc", "ccac", false},
		{"全相同", "aaaa", "aaaa", true},
		{"顺序不同", "listen", "silent", true},
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
