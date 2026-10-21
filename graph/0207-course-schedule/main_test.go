package courseschedule

import "testing"

var solutions = map[string]func(int, [][]int) bool{
	"canFinish":    canFinish,
	"canFinishDFS": canFinishDFS,
}

func TestCanFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{"示例1-可完成", 2, [][]int{{1, 0}}, true},
		{"示例2-两门互相依赖", 2, [][]int{{1, 0}, {0, 1}}, false},
		{"无先修", 3, [][]int{}, true},
		{"自环", 1, [][]int{{0, 0}}, false},
		{"菱形依赖", 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true},
		{"三元环", 3, [][]int{{1, 0}, {2, 1}, {0, 2}}, false},
		{"环在独立分量里", 5, [][]int{{1, 0}, {3, 2}, {4, 3}, {2, 4}}, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.numCourses, tt.prerequisites); got != tt.want {
					t.Errorf("%s(%d, %v) = %v, want %v", name, tt.numCourses, tt.prerequisites, got, tt.want)
				}
			})
		}
	}
}
