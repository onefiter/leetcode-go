package goodnodes

import (
	"math"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode) int{
	"goodNodes":            goodNodes,
	"goodNodesBFS":         goodNodesBFS,
	"goodNodesEnumeration": goodNodesEnumeration,
}

// buildTree 按 LeetCode 层序数组构造二叉树，null 表示空节点。
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	for i := 1; i < len(vals); {
		cur := queue[0]
		queue = queue[1:]
		if vals[i] != null {
			cur.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, cur.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			cur.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, cur.Right)
		}
		i++
	}
	return root
}

func TestGoodNodes(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{"示例1", []int{3, 1, 4, 3, null, 1, 5}, 4},
		{"示例2", []int{3, 3, null, 4, 2}, 3},
		{"示例3-单节点", []int{1}, 1},
		{"空树", []int{}, 0},
		{"左斜链递增", []int{1, 2, null, 3, null, 4}, 4},
		{"右斜链递减", []int{4, null, 3, null, 2}, 1},
		{"完全二叉树", []int{2, 1, 3, 0, 5, 2, 4}, 4},
		{"负数", []int{-1, -2, -1, null, null, -3, 0}, 3},
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
