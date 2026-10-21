package maxsquare

import (
	"slices"
	"testing"
)

var solutions = map[string]func([][]byte) int{
	"maximalSquare":      maximalSquare,
	"maximalSquareBrute": maximalSquareBrute,
}

func cloneMatrix(m [][]byte) [][]byte {
	res := make([][]byte, len(m))
	for i := range m {
		res[i] = slices.Clone(m[i])
	}
	return res
}

func TestMaximalSquare(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]byte
		want   int
	}{
		{"示例1", [][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}, 4},
		{"示例2", [][]byte{{'0', '1'}, {'1', '0'}}, 1},
		{"示例3单个零", [][]byte{{'0'}}, 0},
		{"单个一", [][]byte{{'1'}}, 1},
		{"全一3x3", [][]byte{
			{'1', '1', '1'},
			{'1', '1', '1'},
			{'1', '1', '1'},
		}, 9},
		{"单行", [][]byte{{'1', '1', '1', '1'}}, 1},
		{"中心缺口", [][]byte{
			{'1', '1', '1'},
			{'1', '0', '1'},
			{'1', '1', '1'},
		}, 1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(cloneMatrix(tt.matrix)); got != tt.want {
					t.Errorf("%s(%q) = %v, want %v", name, tt.matrix, got, tt.want)
				}
			})
		}
	}
}
