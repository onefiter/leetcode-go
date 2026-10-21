package rangesumimmutable

import "testing"

func TestNumArray(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		queries [][2]int // {left, right}
		want    []int
	}{
		{
			name:    "典型用例",
			nums:    []int{-2, 0, 3, -5, 2, -1},
			queries: [][2]int{{0, 2}, {2, 5}, {0, 5}},
			want:    []int{1, -1, -3},
		},
		{
			name:    "单元素",
			nums:    []int{5},
			queries: [][2]int{{0, 0}},
			want:    []int{5},
		},
		{
			name:    "全相同",
			nums:    []int{4, 4, 4, 4},
			queries: [][2]int{{0, 0}, {1, 3}, {0, 3}},
			want:    []int{4, 12, 16},
		},
		{
			name:    "两端点相同",
			nums:    []int{1, 2, 3},
			queries: [][2]int{{1, 1}, {2, 2}},
			want:    []int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			na := Constructor(tt.nums)
			for i, q := range tt.queries {
				got := na.SumRange(q[0], q[1])
				if got != tt.want[i] {
					t.Errorf("SumRange(%d, %d) = %d, want %d", q[0], q[1], got, tt.want[i])
				}
			}
		})
	}
}
