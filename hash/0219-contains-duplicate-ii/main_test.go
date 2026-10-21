package containsduplicateii

import "testing"

var solutions = map[string]func([]int, int) bool{
	"containsNearbyDuplicate":            containsNearbyDuplicate,
	"containsNearbyDuplicateWindow":      containsNearbyDuplicateWindow,
	"containsNearbyDuplicateEnumeration": containsNearbyDuplicateEnumeration,
}

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want bool
	}{
		{"示例1", []int{1, 2, 3, 1}, 3, true},
		{"示例2", []int{1, 0, 1, 1}, 1, true},
		{"示例3", []int{1, 2, 3, 1, 2, 3}, 2, false},
		{"空数组", []int{}, 1, false},
		{"单元素", []int{1}, 1, false},
		{"k=0", []int{1, 1}, 0, false},
		{"k边界刚好", []int{1, 2, 1}, 2, true},
		{"k边界差一", []int{1, 2, 1}, 1, false},
		{"全相同小k", []int{5, 5, 5, 5}, 1, true},
		{"无重复", []int{1, 2, 3, 4, 5}, 10, false},
		{"含负数", []int{-1, -1}, 1, true},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums, tt.k); got != tt.want {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.nums, tt.k, got, tt.want)
				}
			})
		}
	}
}
