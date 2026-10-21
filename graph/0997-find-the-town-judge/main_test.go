package townjudge

import "testing"

var solutions = map[string]func(int, [][]int) int{
	"findJudge":           findJudge,
	"findJudgeBruteForce": findJudgeBruteForce,
}

func TestFindJudge(t *testing.T) {
	tests := []struct {
		name  string
		n     int
		trust [][]int
		want  int
	}{
		{"示例1", 2, [][]int{{1, 2}}, 2},
		{"示例2", 3, [][]int{{1, 3}, {2, 3}}, 3},
		{"示例3-法官信任别人", 3, [][]int{{1, 3}, {2, 3}, {3, 1}}, -1},
		{"只有一个人", 1, [][]int{}, 1},
		{"两人无信任", 2, [][]int{}, -1},
		{"有人不信任候选人", 3, [][]int{{1, 3}}, -1},
		{"多人互信但存在法官", 4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}, 3},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.n, tt.trust); got != tt.want {
					t.Errorf("%s(n=%d, %v) = %d, want %d", name, tt.n, tt.trust, got, tt.want)
				}
			})
		}
	}
}
