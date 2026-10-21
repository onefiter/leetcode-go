package containsduplicateiii

import "testing"

var solutions = map[string]func([]int, int, int) bool{
	"containsNearbyAlmostDuplicate":             containsNearbyAlmostDuplicate,
	"containsNearbyAlmostDuplicateSortedWindow": containsNearbyAlmostDuplicateSortedWindow,
	"containsNearbyAlmostDuplicateEnumeration":  containsNearbyAlmostDuplicateEnumeration,
}

func TestContainsNearbyAlmostDuplicate(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		indexDiff int
		valueDiff int
		want      bool
	}{
		{"示例1", []int{1, 2, 3, 1}, 3, 0, true},
		{"示例2", []int{1, 5, 9, 1, 5, 9}, 2, 3, false},
		{"空数组", []int{}, 1, 1, false},
		{"单元素", []int{1}, 1, 1, false},
		{"indexDiff为0", []int{1, 1}, 0, 0, false},
		{"valueDiff边界命中", []int{1, 3}, 1, 2, true},
		{"valueDiff差一不命中", []int{1, 4}, 1, 2, false},
		{"含负数命中", []int{-3, -1}, 1, 2, true},
		{"溢出场景", []int{2147483647, -2147483647}, 1, 2147483647, false},
		{"相邻桶命中", []int{4, 2}, 1, 2, true},
		{"距离超出", []int{1, 2, 3, 1}, 2, 0, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums, tt.indexDiff, tt.valueDiff); got != tt.want {
					t.Errorf("%s(%v, %d, %d) = %v, want %v",
						name, tt.nums, tt.indexDiff, tt.valueDiff, got, tt.want)
				}
			})
		}
	}
}
