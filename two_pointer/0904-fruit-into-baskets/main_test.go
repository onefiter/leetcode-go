package fruitbaskets

import "testing"

var solutions = map[string]func([]int) int{
	"totalFruit":            totalFruit,
	"totalFruitEnumeration": totalFruitEnumeration,
}

func TestTotalFruit(t *testing.T) {
	tests := []struct {
		name   string
		fruits []int
		want   int
	}{
		{"示例1", []int{1, 2, 1}, 3},
		{"示例2", []int{0, 1, 2, 2}, 3},
		{"示例3", []int{1, 2, 3, 2, 2}, 4},
		{"示例4", []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}, 5},
		{"空数组", []int{}, 0},
		{"单元素", []int{7}, 1},
		{"全相同", []int{4, 4, 4, 4}, 4},
		{"只有两种", []int{1, 2, 2, 1, 1, 2}, 6},
		{"每个都不同", []int{1, 2, 3, 4, 5}, 2},
		{"含种类0", []int{0, 0, 1, 1, 2}, 4},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.fruits); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.fruits, got, tt.want)
				}
			})
		}
	}
}
