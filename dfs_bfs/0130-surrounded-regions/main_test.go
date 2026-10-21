package surrounded

import (
	"reflect"
	"testing"
)

// toBoard 把字符串切片转成 [][]byte，每次调用都会生成新的副本。
func toBoard(rows []string) [][]byte {
	board := make([][]byte, len(rows))
	for i, r := range rows {
		board[i] = []byte(r)
	}
	return board
}

var solutions = map[string]func([][]byte){
	"solve":    solve,
	"solveBFS": solveBFS,
}

func TestSolve(t *testing.T) {
	tests := []struct {
		name  string
		board []string
		want  []string
	}{
		{"示例1", []string{"XXXX", "XOOX", "XXOX", "XOXX"}, []string{"XXXX", "XXXX", "XXXX", "XOXX"}},
		{"示例2-单格", []string{"X"}, []string{"X"}},
		{"单个O在边界", []string{"O"}, []string{"O"}},
		{"全是O", []string{"OOO", "OOO", "OOO"}, []string{"OOO", "OOO", "OOO"}},
		{"中心被围", []string{"XXX", "XOX", "XXX"}, []string{"XXX", "XXX", "XXX"}},
		{"通过链连到边界", []string{"XXXXX", "XOOOX", "XOXOX", "XOXXX", "XOXXX"}, []string{"XXXXX", "XOOOX", "XOXOX", "XOXXX", "XOXXX"}},
		{"多区域混合", []string{"OXXOX", "XOOXO", "XXXOX", "OXOXX"}, []string{"OXXOX", "XXXXO", "XXXXX", "OXOXX"}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				board := toBoard(tt.board)
				fn(board)
				if want := toBoard(tt.want); !reflect.DeepEqual(board, want) {
					t.Errorf("%s(%v) = %q, want %q", name, tt.board, board, want)
				}
			})
		}
	}
}
