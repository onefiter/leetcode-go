package triangulation

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int) int{
	"minScoreTriangulation":     minScoreTriangulation,
	"minScoreTriangulationMemo": minScoreTriangulationMemo,
}

func TestMinScoreTriangulation(t *testing.T) {
	tests := []struct {
		name   string
		values []int
		want   int
	}{
		{"示例1", []int{1, 2, 3}, 6},
		{"示例2", []int{3, 7, 4, 5}, 144},
		{"示例3", []int{1, 3, 1, 4, 1, 5}, 13},
		{"四边形全相等", []int{2, 2, 2, 2}, 16},
		{"五边形全为1", []int{1, 1, 1, 1, 1}, 3},
		{"四边形选短对角线", []int{1, 2, 3, 4}, 18},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(slices.Clone(tt.values)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.values, got, tt.want)
				}
			})
		}
	}
}
