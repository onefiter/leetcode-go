package singlenonduplicate

import "testing"

var solutions = map[string]func([]int) int{
	"singleNonDuplicate":            singleNonDuplicate,
	"singleNonDuplicateEnumeration": singleNonDuplicateEnumeration,
}

func TestSingleNonDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"示例1", []int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{"示例2", []int{3, 3, 7, 7, 10, 11, 11}, 10},
		{"单元素", []int{1}, 1},
		{"在开头", []int{1, 2, 2, 3, 3}, 1},
		{"在末尾", []int{1, 1, 2, 2, 3}, 3},
		{"含零和负数", []int{-3, -3, 0, 5, 5}, 0},
		{"三个元素-开头", []int{1, 2, 2}, 1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.nums); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.nums, got, tt.want)
				}
			})
		}
	}
}

// TestSingleNonDuplicateAllPositions 单一元素依次放在每个位置上校验。
// 构造方式：值 0..pairs，除 single 外每个值出现两次，天然有序。
func TestSingleNonDuplicateAllPositions(t *testing.T) {
	for pairs := 0; pairs <= 8; pairs++ {
		for single := 0; single <= pairs; single++ {
			nums := []int{}
			for v := 0; v <= pairs; v++ {
				nums = append(nums, v)
				if v != single {
					nums = append(nums, v)
				}
			}
			for name, fn := range solutions {
				if got := fn(nums); got != single {
					t.Errorf("%s(%v) = %d, want %d", name, nums, got, single)
				}
			}
		}
	}
}
