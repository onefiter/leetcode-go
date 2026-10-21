package maximumunits

import (
	"slices"
	"testing"
)

var solutions = map[string]func([][]int, int) int{
	"maximumUnits":       maximumUnits,
	"maximumUnitsBucket": maximumUnitsBucket,
}

// deepClone 深拷贝二维切片，避免解法原地排序影响其他用例
func deepClone(src [][]int) [][]int {
	out := make([][]int, len(src))
	for i, row := range src {
		out[i] = slices.Clone(row)
	}
	return out
}

func TestMaximumUnits(t *testing.T) {
	tests := []struct {
		name      string
		boxTypes  [][]int
		truckSize int
		want      int
	}{
		{"示例1", [][]int{{1, 3}, {2, 2}, {3, 1}}, 4, 8},
		{"示例2", [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10, 91},
		{"卡车装不满", [][]int{{1, 5}, {2, 3}}, 10, 11},
		{"只装一类的一部分", [][]int{{10, 4}}, 3, 12},
		{"单元数相同", [][]int{{2, 6}, {3, 6}}, 4, 24},
		{"卡车容量为1", [][]int{{3, 1}, {1, 100}, {2, 50}}, 1, 100},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(deepClone(tt.boxTypes), tt.truckSize)
				if got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.boxTypes, tt.truckSize, got, tt.want)
				}
			})
		}
	}
}
