package scheduleiii

import (
	"slices"
	"testing"
)

var solutions = map[string]func([][]int) int{
	"scheduleCourse":   scheduleCourse,
	"scheduleCourseDP": scheduleCourseDP,
}

// deepClone 深拷贝二维切片，避免解法原地排序影响其他用例
func deepClone(src [][]int) [][]int {
	out := make([][]int, len(src))
	for i, row := range src {
		out[i] = slices.Clone(row)
	}
	return out
}

func TestScheduleCourse(t *testing.T) {
	tests := []struct {
		name    string
		courses [][]int
		want    int
	}{
		{"示例1", [][]int{{100, 200}, {200, 1300}, {1000, 1250}, {2000, 3200}}, 3},
		{"示例2", [][]int{{1, 2}}, 1},
		{"示例3", [][]int{{3, 2}, {4, 3}}, 0},
		{"需要反悔替换长课", [][]int{{5, 5}, {4, 6}, {2, 6}}, 2},
		{"全部可修", [][]int{{1, 2}, {2, 3}, {3, 6}}, 3},
		{"截止相同", [][]int{{2, 4}, {2, 4}, {2, 4}}, 2},
		{"乱序输入", [][]int{{7, 17}, {3, 12}, {10, 20}, {9, 10}, {5, 20}, {10, 19}, {4, 18}}, 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(deepClone(tt.courses))
				if got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.courses, got, tt.want)
				}
			})
		}
	}
}
