package wordsearch

import (
	"slices"
	"testing"
)

// cloneBoard 深拷贝网格，避免解法原地修改影响其他用例。
func cloneBoard(board [][]byte) [][]byte {
	out := make([][]byte, len(board))
	for i, row := range board {
		out[i] = slices.Clone(row)
	}
	return out
}

// toBoard 把字符串行转成 [][]byte，方便书写用例。
func toBoard(rows ...string) [][]byte {
	out := make([][]byte, len(rows))
	for i, r := range rows {
		out[i] = []byte(r)
	}
	return out
}

var solutions = map[string]func([][]byte, string) bool{
	"exist":        exist,
	"existVisited": existVisited,
}

func TestExist(t *testing.T) {
	classic := toBoard("ABCE", "SFCS", "ADEE")
	tests := []struct {
		name  string
		board [][]byte
		word  string
		want  bool
	}{
		{"示例1", classic, "ABCCED", true},
		{"示例2", classic, "SEE", true},
		{"示例3-不能重复使用", classic, "ABCB", false},
		{"单格匹配", toBoard("A"), "A", true},
		{"单格不匹配", toBoard("A"), "B", false},
		{"单词比网格长", toBoard("AB"), "ABA", false},
		{"需要回溯绕路", toBoard("AAB", "AAA", "BAA"), "AAAAB", true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(cloneBoard(tt.board), tt.word)
				if got != tt.want {
					t.Errorf("%s(%q, %q) = %v, want %v", name, tt.board, tt.word, got, tt.want)
				}
			})
		}
	}
}
