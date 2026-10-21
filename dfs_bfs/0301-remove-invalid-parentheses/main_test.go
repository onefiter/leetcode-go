package removeparens

import (
	"slices"
	"testing"
)

var solutions = map[string]func(string) []string{
	"removeInvalidParentheses":    removeInvalidParentheses,
	"removeInvalidParenthesesBFS": removeInvalidParenthesesBFS,
}

func TestRemoveInvalidParentheses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"示例1", "()())()", []string{"(())()", "()()()"}},
		{"示例2-含字母", "(a)())()", []string{"(a())()", "(a)()()"}},
		{"示例3-全删", ")(", []string{""}},
		{"本身合法", "(())", []string{"(())"}},
		{"只有字母", "abc", []string{"abc"}},
		{"多余左括号", "((()", []string{"()"}},
		{"字母与括号混合", "x(", []string{"x"}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := slices.Clone(fn(tt.s))
				slices.Sort(got)
				want := slices.Clone(tt.want)
				slices.Sort(want)
				if !slices.Equal(got, want) {
					t.Errorf("%s(%q) = %q, want %q", name, tt.s, got, want)
				}
			})
		}
	}
}
