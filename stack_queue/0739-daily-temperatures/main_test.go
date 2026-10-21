package dailytemperatures

import (
	"math/rand"
	"slices"
	"testing"
)

var solutions = map[string]func([]int) []int{
	"dailyTemperatures":            dailyTemperatures,
	"dailyTemperaturesJump":        dailyTemperaturesJump,
	"dailyTemperaturesEnumeration": dailyTemperaturesEnumeration,
}

func TestDailyTemperatures(t *testing.T) {
	tests := []struct {
		name  string
		temps []int
		want  []int
	}{
		{"示例1", []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{"示例2", []int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{"示例3", []int{30, 60, 90}, []int{1, 1, 0}},
		{"空数组", []int{}, []int{}},
		{"单元素", []int{50}, []int{0}},
		{"全相同", []int{70, 70, 70}, []int{0, 0, 0}},
		{"严格递减", []int{90, 80, 70}, []int{0, 0, 0}},
		{"最后升温", []int{50, 40, 40, 60}, []int{3, 2, 1, 0}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.temps); !slices.Equal(got, tt.want) {
					t.Errorf("%s(%v) = %v, want %v", name, tt.temps, got, tt.want)
				}
			})
		}
	}
}

// TestAgainstEnumeration 随机数组对拍。
func TestAgainstEnumeration(t *testing.T) {
	r := rand.New(rand.NewSource(739))
	for k := 0; k < 500; k++ {
		temps := make([]int, r.Intn(30))
		for i := range temps {
			temps[i] = 30 + r.Intn(10)
		}
		want := dailyTemperaturesEnumeration(temps)
		for name, fn := range solutions {
			if got := fn(temps); !slices.Equal(got, want) {
				t.Fatalf("%s(%v) = %v, want %v", name, temps, got, want)
			}
		}
	}
}
