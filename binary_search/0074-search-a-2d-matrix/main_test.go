package searchmatrix

import "testing"

var solutions = map[string]func([][]int, int) bool{
	"searchMatrix":            searchMatrix,
	"searchMatrixTwoBinary":   searchMatrixTwoBinary,
	"searchMatrixEnumeration": searchMatrixEnumeration,
}

func TestSearchMatrix(t *testing.T) {
	sample := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{"示例1", sample, 3, true},
		{"示例2", sample, 13, false},
		{"第一个元素", sample, 1, true},
		{"最后一个元素", sample, 60, true},
		{"行首", sample, 23, true},
		{"比最小还小", sample, 0, false},
		{"比最大还大", sample, 61, false},
		{"行间缺口", sample, 8, false},
		{"空矩阵", [][]int{}, 1, false},
		{"空行", [][]int{{}}, 1, false},
		{"单元素-命中", [][]int{{1}}, 1, true},
		{"单元素-不命中", [][]int{{1}}, 2, false},
		{"单行", [][]int{{1, 3}}, 3, true},
		{"单列-命中", [][]int{{1}, {3}}, 3, true},
		{"单列-不命中", [][]int{{1}, {3}}, 2, false},
		{"含负数", [][]int{{-5, -3}, {-1, 0}}, -3, true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.matrix, tt.target); got != tt.want {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.matrix, tt.target, got, tt.want)
				}
			})
		}
	}
}

// TestSearchMatrixAgainstEnumeration 在每个可能的 target 上用暴力结果核对。
func TestSearchMatrixAgainstEnumeration(t *testing.T) {
	// 3x4 矩阵，元素为 0,3,6,...（步长 3，留出缺口）
	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 4)
		for j := range matrix[i] {
			matrix[i][j] = (i*4 + j) * 3
		}
	}
	for target := -2; target <= 36; target++ {
		want := searchMatrixEnumeration(matrix, target)
		for name, fn := range solutions {
			if got := fn(matrix, target); got != want {
				t.Errorf("%s(target=%d) = %v, want %v", name, target, got, want)
			}
		}
	}
}
