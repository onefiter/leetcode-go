package nextgreater2

import (
	"math/rand"
	"slices"
	"testing"
)

var solutions = map[string]func([]int) []int{
	"nextGreaterElements":            nextGreaterElements,
	"nextGreaterElementsEnumeration": nextGreaterElementsEnumeration,
}

func TestNextGreaterElements(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{"示例1", []int{1, 2, 1}, []int{2, -1, 2}},
		{"示例2", []int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
		{"空数组", []int{}, []int{}},
		{"单元素", []int{5}, []int{-1}},
		{"全相同", []int{2, 2, 2}, []int{-1, -1, -1}},
		{"严格递减", []int{3, 2, 1}, []int{-1, 3, 3}},
		{"严格递增", []int{1, 2, 3}, []int{2, 3, -1}},
		{"最大值在头", []int{5, 1, 2}, []int{-1, 2, 5}},
		{"含负数", []int{-1, -2, 0}, []int{0, 0, -1}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}

// TestAgainstEnumeration 随机数组对拍。
func TestAgainstEnumeration(t *testing.T) {
	r := rand.New(rand.NewSource(503))
	for k := 0; k < 500; k++ {
		nums := make([]int, r.Intn(20))
		for i := range nums {
			nums[i] = r.Intn(10) - 3
		}
		want := nextGreaterElementsEnumeration(nums)
		if got := nextGreaterElements(nums); !slices.Equal(got, want) {
			t.Fatalf("nextGreaterElements(%v) = %v, want %v", nums, got, want)
		}
	}
}
