package removeelement

import (
	"slices"
	"testing"
)

var solutions = map[string]func([]int, int) int{
	"removeElement":            removeElement,
	"removeElementEnumeration": removeElementEnumeration,
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		want []int // 保留的元素（顺序无关）
	}{
		{[]int{3, 2, 2, 3}, 3, []int{2, 2}},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, []int{0, 0, 1, 3, 4}},
		{[]int{}, 1, []int{}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			nums := slices.Clone(tt.nums)
			k := fn(nums, tt.val)
			got := slices.Clone(nums[:k])
			slices.Sort(got)
			if !slices.Equal(got, tt.want) {
				t.Errorf("%s(%v, %d) kept %v, want %v", name, tt.nums, tt.val, got, tt.want)
			}
		}
	}
}
