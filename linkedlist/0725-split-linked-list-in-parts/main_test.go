package splitlist

import (
	"leetcode-go/template"
	"slices"
	"testing"
)

var solutions = map[string]func(*ListNode, int) []*ListNode{
	"splitListToParts":            splitListToParts,
	"splitListToPartsEnumeration": splitListToPartsEnumeration,
}

func TestSplitListToParts(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		k    int
		want [][]int // nil 部分用空切片表示
	}{
		{"示例1-k大于长度", []int{1, 2, 3}, 5, [][]int{{1}, {2}, {3}, {}, {}}},
		{"示例2", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3, [][]int{{1, 2, 3, 4}, {5, 6, 7}, {8, 9, 10}}},
		{"空链表", []int{}, 3, [][]int{{}, {}, {}}},
		{"k为1", []int{1, 2, 3}, 1, [][]int{{1, 2, 3}}},
		{"单节点k为1", []int{1}, 1, [][]int{{1}}},
		{"恰好整除", []int{1, 2, 3, 4, 5, 6}, 3, [][]int{{1, 2}, {3, 4}, {5, 6}}},
		{"k等于长度", []int{1, 2, 3}, 3, [][]int{{1}, {2}, {3}}},
		{"余数为1", []int{1, 2, 3, 4}, 3, [][]int{{1, 2}, {3}, {4}}},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				// 解法会原地断开链表，每次重新构造输入。
				parts := fn(template.BuildList(tt.vals), tt.k)
				if len(parts) != len(tt.want) {
					t.Fatalf("%s(%v, %d) 返回 %d 段, want %d", name, tt.vals, tt.k, len(parts), len(tt.want))
				}
				for i, p := range parts {
					if got := template.ListToSlice(p); !slices.Equal(got, tt.want[i]) {
						t.Errorf("%s(%v, %d) 第 %d 段 = %v, want %v", name, tt.vals, tt.k, i, got, tt.want[i])
					}
					if len(tt.want[i]) == 0 && p != nil {
						t.Errorf("%s 第 %d 段应为 nil", name, i)
					}
				}
			})
		}
	}
}
