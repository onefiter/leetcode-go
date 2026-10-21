package compareversion

import "testing"

var solutions = map[string]func(string, string) int{
	"compareVersion":     compareVersion,
	"compareVersionScan": compareVersionScan,
}

func TestCompareVersion(t *testing.T) {
	tests := []struct {
		name string
		v1   string
		v2   string
		want int
	}{
		{"前导零相等", "1.01", "1.001", 0},
		{"补零相等", "1.0", "1.0.0", 0},
		{"小于", "0.1", "1.1", -1},
		{"大于", "1.0.1", "1", 1},
		{"示例-7比10小", "7.5.2.4", "7.5.3", -1},
		{"相同", "1.2.3", "1.2.3", 0},
		{"单段大于", "2", "1", 1},
		{"单段小于", "1", "2", -1},
		{"多零与单零", "1.0.0.0", "1", 0},
		{"尾部非零大于", "1.1", "1.0", 1},
		{"大数字", "1000.1", "999.9", 1},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(tt.v1, tt.v2); got != tt.want {
					t.Errorf("%s(%q, %q) = %d, want %d", name, tt.v1, tt.v2, got, tt.want)
				}
			})
		}
	}
}
