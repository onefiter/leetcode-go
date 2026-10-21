package maximum69

import "testing"

var solutions = map[string]func(int) int{
	"maximum69Number":      maximum69Number,
	"maximum69NumberBrute": maximum69NumberBrute,
}

func TestMaximum69Number(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"示例1", 9669, 9969},
		{"示例2", 9996, 9999},
		{"示例3", 9999, 9999},
		{"单个6", 6, 9},
		{"单个9", 9, 9},
		{"全是6", 6666, 9666},
		{"最高位就是6", 6999, 9999},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.num)
				if got != tt.want {
					t.Errorf("%s(%d) = %d, want %d", name, tt.num, got, tt.want)
				}
			})
		}
	}
}
