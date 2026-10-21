package reversestring

import (
	"slices"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{"示例1", []byte("hello"), []byte("olleh")},
		{"示例2", []byte("Hannah"), []byte("hannaH")},
		{"空数组", []byte(""), []byte("")},
		{"单字符", []byte("a"), []byte("a")},
		{"两字符", []byte("ab"), []byte("ba")},
		{"偶数长度", []byte("abcd"), []byte("dcba")},
		{"回文", []byte("level"), []byte("level")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slices.Clone(tt.s)
			reverseString(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("reverseString(%q) = %q, want %q", tt.s, got, tt.want)
			}
		})
	}
}
