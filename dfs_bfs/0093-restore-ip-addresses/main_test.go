package restoreip

import (
	"slices"
	"testing"
)

var solutions = map[string]func(string) []string{
	"restoreIpAddresses":            restoreIpAddresses,
	"restoreIpAddressesEnumeration": restoreIpAddressesEnumeration,
}

func TestRestoreIpAddresses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"示例1", "25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"示例2-全零", "0000", []string{"0.0.0.0"}},
		{"示例3", "101023", []string{"1.0.10.23", "1.0.102.3", "10.1.0.23", "10.10.2.3", "101.0.2.3"}},
		{"太短", "123", []string{}},
		{"太长", "1234567890123", []string{}},
		{"超过255", "256256256256", []string{}},
		{"前导零", "010010", []string{"0.10.0.10", "0.100.1.0"}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				got := slices.Clone(fn(tt.s))
				slices.Sort(got)
				want := slices.Clone(tt.want)
				slices.Sort(want)
				if !slices.Equal(got, want) {
					t.Errorf("%s(%q) = %v, want %v", name, tt.s, got, want)
				}
			})
		}
	}
}
