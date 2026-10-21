package evaldivision

import (
	"math"
	"testing"
)

const eps = 1e-5

var solutions = map[string]func([][]string, []float64, [][]string) []float64{
	"calcEquation":      calcEquation,
	"calcEquationFloyd": calcEquationFloyd,
}

func TestCalcEquation(t *testing.T) {
	tests := []struct {
		name      string
		equations [][]string
		values    []float64
		queries   [][]string
		want      []float64
	}{
		{"示例1",
			[][]string{{"a", "b"}, {"b", "c"}}, []float64{2.0, 3.0},
			[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			[]float64{6.0, 0.5, -1.0, 1.0, -1.0}},
		{"示例2",
			[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}}, []float64{1.5, 2.5, 5.0},
			[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			[]float64{3.75, 0.4, 5.0, 0.2}},
		{"示例3",
			[][]string{{"a", "b"}}, []float64{0.5},
			[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			[]float64{0.5, 2.0, -1.0, -1.0}},
		{"不同分量不可求",
			[][]string{{"a", "b"}, {"c", "d"}}, []float64{2.0, 4.0},
			[][]string{{"a", "d"}, {"d", "c"}, {"b", "a"}},
			[]float64{-1.0, 0.25, 0.5}},
		{"合并两棵树后跨树查询",
			[][]string{{"a", "b"}, {"c", "d"}, {"b", "d"}}, []float64{2.0, 3.0, 5.0},
			[][]string{{"a", "c"}, {"c", "a"}, {"a", "d"}},
			[]float64{10.0 / 3.0, 0.3, 10.0}},
		{"长链",
			[][]string{{"a", "b"}, {"b", "c"}, {"c", "d"}, {"d", "e"}}, []float64{2.0, 2.0, 2.0, 2.0},
			[][]string{{"a", "e"}, {"e", "a"}, {"b", "d"}},
			[]float64{16.0, 0.0625, 4.0}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := fn(tt.equations, tt.values, tt.queries)
				if len(got) != len(tt.want) {
					t.Fatalf("%s 返回长度 %d, want %d", name, len(got), len(tt.want))
				}
				for i := range got {
					if math.Abs(got[i]-tt.want[i]) > eps {
						t.Errorf("%s 查询 %v = %v, want %v", name, tt.queries[i], got[i], tt.want[i])
					}
				}
			})
		}
	}
}
