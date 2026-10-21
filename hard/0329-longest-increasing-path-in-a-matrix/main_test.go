package lip

import "testing"

var solutions = map[string]func([][]int) int{
	"longestIncreasingPath":     longestIncreasingPath,
	"longestIncreasingPathTopo": longestIncreasingPathTopo,
}

func TestLongestIncreasingPath(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   int
	}{
		{"示例1", [][]int{{9, 9, 4}, {6, 6, 8}, {2, 1, 1}}, 4},
		{"示例2", [][]int{{3, 4, 5}, {3, 2, 6}, {2, 2, 1}}, 4},
		{"单个格子", [][]int{{1}}, 1},
		{"全部相等", [][]int{{7, 7}, {7, 7}}, 1},
		{"单行递增", [][]int{{1, 2, 3, 4, 5}}, 5},
		{"蛇形递增", [][]int{{1, 2, 3}, {6, 5, 4}, {7, 8, 9}}, 9},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.matrix); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.matrix, got, tt.want)
				}
			})
		}
	}
}
