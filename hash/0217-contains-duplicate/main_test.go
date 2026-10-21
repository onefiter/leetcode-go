package containsduplicate

import "testing"

var solutions = map[string]func([]int) bool{
	"containsDuplicate":            containsDuplicate,
	"containsDuplicateSort":        containsDuplicateSort,
	"containsDuplicateEnumeration": containsDuplicateEnumeration,
}

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"示例1-有重复", []int{1, 2, 3, 1}, true},
		{"示例2-无重复", []int{1, 2, 3, 4}, false},
		{"示例3-多重复", []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{"空数组", []int{}, false},
		{"单元素", []int{7}, false},
		{"全相同", []int{5, 5, 5}, true},
		{"含负数有重复", []int{-1, -2, -1}, true},
		{"含负数无重复", []int{-1, -2, -3}, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums); got != tt.want {
					t.Errorf("%s(%v) = %v, want %v", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}
