package partitionlabels

import (
	"slices"
	"testing"
)

var solutions = map[string]func(string) []int{
	"partitionLabels":      partitionLabels,
	"partitionLabelsMerge": partitionLabelsMerge,
}

func TestPartitionLabels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []int
	}{
		{"示例1", "ababcbacadefegdehijhklij", []int{9, 7, 8}},
		{"示例2", "eccbbbbdec", []int{10}},
		{"单字符", "a", []int{1}},
		{"全不同", "abc", []int{1, 1, 1}},
		{"全相同", "aaaa", []int{4}},
		{"首尾相同", "abca", []int{4}},
		{"两段", "aabbcb", []int{2, 4}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.s)
				if !slices.Equal(got, tt.want) {
					t.Errorf("%s(%q) = %v, want %v", name, tt.s, got, tt.want)
				}
			})
		}
	}
}
