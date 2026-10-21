package updatematrix

import (
	"reflect"
	"slices"
	"testing"
)

// cloneGrid 深拷贝二维切片。
func cloneGrid(g [][]int) [][]int {
	out := make([][]int, len(g))
	for i := range g {
		out[i] = slices.Clone(g[i])
	}
	return out
}

var solutions = map[string]func([][]int) [][]int{
	"updateMatrix":   updateMatrix,
	"updateMatrixDP": updateMatrixDP,
}

func TestUpdateMatrix(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want [][]int
	}{
		{"示例1", [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		{"示例2", [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
		{"单个0", [][]int{{0}}, [][]int{{0}}},
		{"单行", [][]int{{1, 1, 0, 1, 1, 1}}, [][]int{{2, 1, 0, 1, 2, 3}}},
		{"单列", [][]int{{1}, {1}, {0}}, [][]int{{2}, {1}, {0}}},
		{"角落唯一的0", [][]int{{0, 1, 1}, {1, 1, 1}, {1, 1, 1}}, [][]int{{0, 1, 2}, {1, 2, 3}, {2, 3, 4}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(cloneGrid(tt.mat))
				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.mat, got, tt.want)
				}
			})
		}
	}
}
