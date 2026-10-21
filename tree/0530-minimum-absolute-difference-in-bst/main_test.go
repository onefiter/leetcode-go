package minbstdiff

import (
	"math"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode) int{
	"getMinimumDifference":            getMinimumDifference,
	"getMinimumDifferenceIterative":   getMinimumDifferenceIterative,
	"getMinimumDifferenceEnumeration": getMinimumDifferenceEnumeration,
}

// buildTree 按 LeetCode 层序数组构造树，null 表示空节点。
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	q := []*TreeNode{root}
	for i := 1; i < len(vals); {
		node := q[0]
		q = q[1:]
		if vals[i] != null {
			node.Left = &TreeNode{Val: vals[i]}
			q = append(q, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			node.Right = &TreeNode{Val: vals[i]}
			q = append(q, node.Right)
		}
		i++
	}
	return root
}

func TestGetMinimumDifference(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"示例1", []int{4, 2, 6, 1, 3}, 1},
		{"示例2", []int{1, 0, 48, null, null, 12, 49}, 1},
		{"两个节点", []int{1, null, 5}, 4},
		{"左斜链", []int{100, 50, null, 20, null, 1}, 19},
		{"右斜链", []int{1, null, 10, null, 30, null, 31}, 1},
		{"完全二叉树", []int{40, 20, 60, 10, 30, 50, 70}, 10},
		{"跨层相邻", []int{236, 104, 701, null, 227, null, 911}, 9},
		{"含负数", []int{0, -10, 10, -20, -5}, 5},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.vals)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.vals, got, tt.want)
				}
			})
		}
	}
}
