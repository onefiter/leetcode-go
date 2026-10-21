package grumpyowner

import "testing"

var solutions = map[string]func([]int, []int, int) int{
	"maxSatisfied":            maxSatisfied,
	"maxSatisfiedEnumeration": maxSatisfiedEnumeration,
}

func TestMaxSatisfied(t *testing.T) {
	tests := []struct {
		name      string
		customers []int
		grumpy    []int
		minutes   int
		want      int
	}{
		{"示例1", []int{1, 0, 1, 2, 1, 1, 7, 5}, []int{0, 1, 0, 1, 0, 1, 0, 1}, 3, 16},
		{"示例2-单元素", []int{1}, []int{0}, 1, 1},
		{"单元素生气", []int{5}, []int{1}, 1, 5},
		{"从不生气", []int{1, 2, 3}, []int{0, 0, 0}, 1, 6},
		{"一直生气", []int{3, 1, 4}, []int{1, 1, 1}, 2, 5},
		{"minutes等于n", []int{2, 3}, []int{1, 1}, 2, 5},
		{"全相同", []int{2, 2, 2, 2}, []int{1, 0, 1, 0}, 1, 6},
		{"挽回段在末尾", []int{1, 1, 1, 9, 9}, []int{0, 0, 0, 1, 1}, 2, 21},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.customers, tt.grumpy, tt.minutes); got != tt.want {
					t.Errorf("%s(%v, %v, %d) = %d, want %d", name, tt.customers, tt.grumpy, tt.minutes, got, tt.want)
				}
			})
		}
	}
}
