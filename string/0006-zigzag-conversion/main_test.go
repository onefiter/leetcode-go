package zigzag

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		numRows int
		want    string
	}{
		{"示例1-三行", "PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"示例2-四行", "PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"示例3-单行", "A", 1, "A"},
		{"单行多字符", "ABCDE", 1, "ABCDE"},
		{"空串", "", 3, ""},
		{"行数大于长度", "AB", 5, "AB"},
		{"行数等于长度", "ABC", 3, "ABC"},
		{"两行", "ABCD", 2, "ACBD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.s, tt.numRows); got != tt.want {
				t.Errorf("convert(%q, %d) = %q, want %q", tt.s, tt.numRows, got, tt.want)
			}
		})
	}
}
