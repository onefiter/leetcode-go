package minswaps

import "testing"

var solutions = map[string]func([]int) int{
	"minSwaps":            minSwaps,
	"minSwapsEnumeration": minSwapsEnumeration,
}

func TestMinSwaps(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"示例1", []int{1, 0, 1, 0, 1}, 1},
		{"示例2", []int{0, 0, 0, 1, 0}, 0},
		{"示例3", []int{1, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1}, 3},
		{"空数组", []int{}, 0},
		{"单个1", []int{1}, 0},
		{"单个0", []int{0}, 0},
		{"全为0", []int{0, 0, 0}, 0},
		{"全为1", []int{1, 1, 1}, 0},
		{"两端各一个1", []int{1, 0, 0, 1}, 1},
		{"已经连续", []int{0, 1, 1, 1, 0}, 0},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.data); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.data, got, tt.want)
				}
			})
		}
	}
}
