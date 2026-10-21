package courseschedule2

import "testing"

// validOrder 校验 order 是否为合法学习顺序：恰好包含 0..n-1 各一次，且每条先修关系 b 在 a 之前。
func validOrder(n int, prerequisites [][]int, order []int) bool {
	if len(order) != n {
		return false
	}
	pos := make([]int, n)
	for i := range pos {
		pos[i] = -1
	}
	for i, c := range order {
		if c < 0 || c >= n || pos[c] != -1 {
			return false
		}
		pos[c] = i
	}
	for _, p := range prerequisites {
		if pos[p[1]] > pos[p[0]] {
			return false
		}
	}
	return true
}

var solutions = map[string]func(int, [][]int) []int{
	"findOrder":    findOrder,
	"findOrderDFS": findOrderDFS,
}

func TestFindOrder(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		possible      bool // false 表示应返回空数组
	}{
		{"示例1-两门课", 2, [][]int{{1, 0}}, true},
		{"示例2-菱形依赖", 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true},
		{"示例3-单门课", 1, [][]int{}, true},
		{"两门互相依赖", 2, [][]int{{1, 0}, {0, 1}}, false},
		{"自环", 2, [][]int{{0, 0}}, false},
		{"长链", 5, [][]int{{1, 0}, {2, 1}, {3, 2}, {4, 3}}, true},
		{"多分量含环", 6, [][]int{{1, 0}, {3, 2}, {4, 3}, {2, 4}, {5, 0}}, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.numCourses, tt.prerequisites)
				if !tt.possible {
					if got == nil || len(got) != 0 {
						t.Errorf("%s(%d, %v) = %v, want []", name, tt.numCourses, tt.prerequisites, got)
					}
					return
				}
				if !validOrder(tt.numCourses, tt.prerequisites, got) {
					t.Errorf("%s(%d, %v) = %v, 不是合法拓扑序", name, tt.numCourses, tt.prerequisites, got)
				}
			})
		}
	}
}
