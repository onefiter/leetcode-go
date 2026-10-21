package rangesum2dimmutable

import "testing"

func TestNumMatrix(t *testing.T) {
	tests := []struct {
		name    string
		matrix  [][]int
		queries [][4]int // {row1, col1, row2, col2}
		want    []int
	}{
		{
			name: "典型用例",
			matrix: [][]int{
				{3, 0, 1, 4, 2},
				{5, 6, 3, 2, 1},
				{1, 2, 0, 1, 5},
				{4, 1, 0, 1, 7},
				{1, 0, 3, 0, 5},
			},
			queries: [][4]int{{2, 1, 4, 3}, {1, 1, 2, 2}, {1, 2, 2, 4}},
			want:    []int{8, 11, 12},
		},
		{
			name:    "单元素",
			matrix:  [][]int{{5}},
			queries: [][4]int{{0, 0, 0, 0}},
			want:    []int{5},
		},
		{
			name:    "全相同",
			matrix:  [][]int{{2, 2}, {2, 2}},
			queries: [][4]int{{0, 0, 1, 1}, {0, 0, 0, 0}, {1, 1, 1, 1}},
			want:    []int{8, 2, 2},
		},
		{
			name:    "单行",
			matrix:  [][]int{{1, 2, 3, 4}},
			queries: [][4]int{{0, 0, 0, 3}, {0, 1, 0, 2}},
			want:    []int{10, 5},
		},
		{
			name:    "单列",
			matrix:  [][]int{{1}, {2}, {3}},
			queries: [][4]int{{0, 0, 2, 0}, {1, 0, 2, 0}},
			want:    []int{6, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			nm := Constructor(tt.matrix)
			for i, q := range tt.queries {
				got := nm.SumRegion(q[0], q[1], q[2], q[3])
				if got != tt.want[i] {
					t.Errorf("SumRegion(%d, %d, %d, %d) = %d, want %d", q[0], q[1], q[2], q[3], got, tt.want[i])
				}
			}
		})
	}
}
