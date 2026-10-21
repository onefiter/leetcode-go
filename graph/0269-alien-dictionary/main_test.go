package aliendict

import "testing"

// validAlienOrder 校验 order：恰好包含 words 中出现的全部字母各一次，
// 且与每对相邻单词第一个不同字符给出的约束一致。
func validAlienOrder(words []string, order string) bool {
	pos := map[byte]int{}
	for i := 0; i < len(order); i++ {
		if _, dup := pos[order[i]]; dup {
			return false
		}
		pos[order[i]] = i
	}
	letters := map[byte]bool{}
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			letters[w[i]] = true
		}
	}
	if len(letters) != len(pos) {
		return false
	}
	for c := range letters {
		if _, ok := pos[c]; !ok {
			return false
		}
	}
	for i := 0; i+1 < len(words); i++ {
		w1, w2 := words[i], words[i+1]
		j := 0
		for j < len(w1) && j < len(w2) && w1[j] == w2[j] {
			j++
		}
		if j < len(w1) && j < len(w2) && pos[w1[j]] > pos[w2[j]] {
			return false
		}
	}
	return true
}

var solutions = map[string]func([]string) string{
	"alienOrder":    alienOrder,
	"alienOrderDFS": alienOrderDFS,
}

func TestAlienOrder(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		possible bool // false 表示应返回 ""
	}{
		{"示例1", []string{"wrt", "wrf", "er", "ett", "rftt"}, true},
		{"示例2", []string{"z", "x"}, true},
		{"示例3-有环", []string{"z", "x", "z"}, false},
		{"前缀在后不合法", []string{"abc", "ab"}, false},
		{"前缀在前合法", []string{"ab", "abc"}, true},
		{"单个单词", []string{"zyx"}, true},
		{"多分量字母", []string{"ac", "ab", "zc", "zb"}, true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.words)
				if !tt.possible {
					if got != "" {
						t.Errorf("%s(%v) = %q, want \"\"", name, tt.words, got)
					}
					return
				}
				if !validAlienOrder(tt.words, got) {
					t.Errorf("%s(%v) = %q, 不是合法字母顺序", name, tt.words, got)
				}
			})
		}
	}
}
