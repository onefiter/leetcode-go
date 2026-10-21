package genemutation

import (
	"slices"
	"testing"
)

var solutions = map[string]func(string, string, []string) int{
	"minMutation":      minMutation,
	"minMutationGraph": minMutationGraph,
}

func TestMinMutation(t *testing.T) {
	tests := []struct {
		name  string
		start string
		end   string
		bank  []string
		want  int
	}{
		{"示例1", "AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}, 1},
		{"示例2", "AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
		{"示例3", "AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}, 3},
		{"起点等于终点", "AAAAAAAA", "AAAAAAAA", []string{}, 0},
		{"终点不在基因库", "AACCGGTT", "AACCGGTA", []string{}, -1},
		{"中间断链", "AAAAAAAA", "CCCCCCCC", []string{"AAAAAAAC", "CCCCCCCC"}, -1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.start, tt.end, slices.Clone(tt.bank))
				if got != tt.want {
					t.Errorf("%s(%q, %q, %v) = %d, want %d", name, tt.start, tt.end, tt.bank, got, tt.want)
				}
			})
		}
	}
}
