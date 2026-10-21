package lettercombinations

import (
	"slices"
	"testing"
)

var solutions = map[string]func(string) []string{
	"letterCombinations":          letterCombinations,
	"letterCombinationsIterative": letterCombinationsIterative,
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name   string
		digits string
		want   []string
	}{
		{"示例1", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"示例2-空串", "", []string{}},
		{"示例3-单个数字", "2", []string{"a", "b", "c"}},
		{"四字母按键", "7", []string{"p", "q", "r", "s"}},
		{"重复数字", "22", []string{"aa", "ab", "ac", "ba", "bb", "bc", "ca", "cb", "cc"}},
		{"含四字母组合", "79", []string{
			"pw", "px", "py", "pz", "qw", "qx", "qy", "qz",
			"rw", "rx", "ry", "rz", "sw", "sx", "sy", "sz",
		}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.digits)
				slices.Sort(got)
				want := slices.Clone(tt.want)
				slices.Sort(want)
				if !slices.Equal(got, want) {
					t.Errorf("%s(%q) = %v, want %v", name, tt.digits, got, want)
				}
			})
		}
	}

	// 长度校验：4 位数字 "2345" 应有 3^4 = 81 种组合
	for name, fn := range solutions {
		if got := len(fn("2345")); got != 81 {
			t.Errorf("%s(\"2345\") 结果数 = %d, want 81", name, got)
		}
	}
}
