package wordsearch

import (
	"reflect"
	"sort"
	"testing"
)

func toBoard(rows ...string) [][]byte {
	board := make([][]byte, len(rows))
	for i, r := range rows {
		board[i] = []byte(r)
	}
	return board
}

func cloneBoard(board [][]byte) [][]byte {
	c := make([][]byte, len(board))
	for i := range board {
		c[i] = append([]byte(nil), board[i]...)
	}
	return c
}

// normalize 排序结果，nil 与空切片视为相同。
func normalize(s []string) []string {
	out := append([]string{}, s...)
	sort.Strings(out)
	return out
}

func TestFindWords(t *testing.T) {
	tests := []struct {
		name  string
		board [][]byte
		words []string
		want  []string
	}{
		{"示例1", toBoard("oaan", "etae", "ihkr", "iflv"),
			[]string{"oath", "pea", "eat", "rain"}, []string{"eat", "oath"}},
		{"示例2", toBoard("ab", "cd"), []string{"abcb"}, []string{}},
		{"单格子", toBoard("a"), []string{"a", "b", "aa"}, []string{"a"}},
		{"格子不能重复使用", toBoard("ab"), []string{"aba", "ab", "ba"}, []string{"ab", "ba"}},
		{"共享前缀", toBoard("abc", "aed", "afg"),
			[]string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"},
			[]string{"abcdefg", "befa", "eaabcdgfa", "gfedcbaaa"}},
		{"重复单词只返回一次", toBoard("aa"), []string{"a", "a", "aa"}, []string{"a", "aa"}},
	}
	funcs := []struct {
		name string
		fn   func([][]byte, []string) []string
	}{
		{"前缀树回溯", findWords},
		{"逐词DFS", findWordsEach},
	}
	for _, f := range funcs {
		for _, tt := range tests {
			t.Run(f.name+"/"+tt.name, func(t *testing.T) {
				got := normalize(f.fn(cloneBoard(tt.board), tt.words))
				want := normalize(tt.want)
				if !reflect.DeepEqual(got, want) {
					t.Errorf("%s() = %v, want %v", f.name, got, want)
				}
			})
		}
	}
}
