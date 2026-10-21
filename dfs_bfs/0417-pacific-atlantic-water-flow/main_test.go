package pacificatlantic

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

// sortCoords 把坐标按 (行, 列) 字典序排序，便于比较顺序不定的结果。
func sortCoords(coords [][]int) [][]int {
	out := cloneGrid(coords)
	slices.SortFunc(out, func(a, b []int) int { return slices.Compare(a, b) })
	return out
}

var solutions = map[string]func([][]int) [][]int{
	"pacificAtlantic":    pacificAtlantic,
	"pacificAtlanticBFS": pacificAtlanticBFS,
}

func TestPacificAtlantic(t *testing.T) {
	tests := []struct {
		name    string
		heights [][]int
		want    [][]int
	}{
		{"示例1",
			[][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}},
			[][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}},
		{"示例2-单格", [][]int{{1}}, [][]int{{0, 0}}},
		{"全部等高", [][]int{{1, 1}, {1, 1}}, [][]int{{0, 0}, {0, 1}, {1, 0}, {1, 1}}},
		{"单行", [][]int{{1, 2, 3}}, [][]int{{0, 0}, {0, 1}, {0, 2}}},
		{"中心高峰", [][]int{{1, 1, 1}, {1, 9, 1}, {1, 1, 1}},
			[][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 1}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
		{"中心低谷", [][]int{{3, 3, 3}, {3, 1, 3}, {3, 3, 3}},
			[][]int{{0, 0}, {0, 1}, {0, 2}, {1, 0}, {1, 2}, {2, 0}, {2, 1}, {2, 2}}},
		{"斜坡", [][]int{{1, 2}, {2, 3}}, [][]int{{0, 1}, {1, 0}, {1, 1}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := sortCoords(fn(cloneGrid(tt.heights)))
				want := sortCoords(tt.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.heights, got, want)
				}
			})
		}
	}
}
