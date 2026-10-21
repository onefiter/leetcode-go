package repeateddna

import (
	"slices"
	"strings"
	"testing"
)

var solutions = map[string]func(string) []string{
	"findRepeatedDnaSequences":            findRepeatedDnaSequences,
	"findRepeatedDnaSequencesHash":        findRepeatedDnaSequencesHash,
	"findRepeatedDnaSequencesEnumeration": findRepeatedDnaSequencesEnumeration,
}

func TestFindRepeatedDnaSequences(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"示例1", "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT", []string{"AAAAACCCCC", "CCCCCAAAAA"}},
		{"示例2-全相同", "AAAAAAAAAAAAA", []string{"AAAAAAAAAA"}},
		{"空串", "", []string{}},
		{"单字符", "A", []string{}},
		{"恰好10位", "ACGTACGTAC", []string{}},
		{"11位全相同", "AAAAAAAAAAA", []string{"AAAAAAAAAA"}},
		{"无重复", "ACGTTGCAACGGTCA", []string{}},
		{"周期4", "ACGTACGTACGTAC", []string{"ACGTACGTAC"}},
		{"多个周期答案", "ACGTACGTACGTACGT", []string{"ACGTACGTAC", "CGTACGTACG", "GTACGTACGT"}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.s)
				slices.Sort(got) // 结果顺序不唯一，排序后比较
				want := slices.Clone(tt.want)
				slices.Sort(want)
				if !slices.Equal(got, want) {
					t.Errorf("%s(%q) = %v, want %v", name, tt.s, got, want)
				}
			})
		}
	}
}

// TestFindRepeatedDnaSequencesAgainstEnumeration 用较长的伪随机序列核对各解法。
func TestFindRepeatedDnaSequencesAgainstEnumeration(t *testing.T) {
	var sb strings.Builder
	x := 7
	for i := 0; i < 600; i++ {
		x = (x*1103515245 + 12345) & 0x7fffffff
		sb.WriteByte("ACGT"[x>>16&3])
	}
	s := sb.String()
	want := findRepeatedDnaSequencesEnumeration(s)
	slices.Sort(want)
	for name, fn := range solutions {
		got := fn(s)
		slices.Sort(got)
		if !slices.Equal(got, want) {
			t.Errorf("%s 结果与暴力不一致: got %d 个, want %d 个", name, len(got), len(want))
		}
	}
}
