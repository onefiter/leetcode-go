package nqueens

import (
	"reflect"
	"slices"
	"testing"
)

// normalize 对外层棋盘按字典序排序（棋盘内的行顺序有意义，不排序）。
func normalize(boards [][]string) [][]string {
	out := make([][]string, len(boards))
	for i, b := range boards {
		out[i] = slices.Clone(b)
	}
	slices.SortFunc(out, func(a, b []string) int {
		return slices.Compare(a, b)
	})
	return out
}

var solutions = map[string]func(int) [][]string{
	"solveNQueens":            solveNQueens,
	"solveNQueensEnumeration": solveNQueensEnumeration,
}

func TestSolveNQueens(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want [][]string
	}{
		{"示例1-n=4", 4, [][]string{
			{".Q..", "...Q", "Q...", "..Q."},
			{"..Q.", "Q...", "...Q", ".Q.."},
		}},
		{"示例2-n=1", 1, [][]string{{"Q"}}},
		{"n=2 无解", 2, [][]string{}},
		{"n=3 无解", 3, [][]string{}},
		{"n=6 四组解", 6, [][]string{
			{".Q....", "...Q..", ".....Q", "Q.....", "..Q...", "....Q."},
			{"..Q...", ".....Q", ".Q....", "....Q.", "Q.....", "...Q.."},
			{"...Q..", "Q.....", "....Q.", ".Q....", ".....Q", "..Q..."},
			{"....Q.", "..Q...", "Q.....", ".....Q", "...Q..", ".Q...."},
		}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := normalize(fn(tt.n))
				want := normalize(tt.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("%s(%d) = %v, want %v", name, tt.n, got, want)
				}
			})
		}
	}

	// 解的数量校验，并且两种解法结果集合一致
	counts := map[int]int{5: 10, 7: 40, 8: 92}
	for n, cnt := range counts {
		var base [][]string
		for name, fn := range solutions {
			got := normalize(fn(n))
			if len(got) != cnt {
				t.Errorf("%s(%d) 解的数量 = %d, want %d", name, n, len(got), cnt)
			}
			if base == nil {
				base = got
			} else if !reflect.DeepEqual(base, got) {
				t.Errorf("n=%d 时两种解法结果不一致", n)
			}
		}
	}
}
