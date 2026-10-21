package pathsum

import (
	"math"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode, int) bool{
	"hasPathSum":          hasPathSum,
	"hasPathSumIterative": hasPathSumIterative,
	"hasPathSumBFS":       hasPathSumBFS,
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

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		name   string
		vals   []int
		target int
		want   bool
	}{
		{"示例1", []int{5, 4, 8, 11, null, 13, 4, 7, 2, null, null, null, 1}, 22, true},
		{"示例2", []int{1, 2, 3}, 5, false},
		{"示例3-空树", []int{}, 0, false},
		{"单节点-命中", []int{1}, 1, true},
		{"单节点-不命中", []int{1}, 2, false},
		{"非叶子不算", []int{1, 2}, 1, false},
		{"左斜链", []int{1, 2, null, 3, null, 4}, 10, true},
		{"右斜链", []int{1, null, 2, null, 3}, 6, true},
		{"负数", []int{-2, null, -3}, -5, true},
		{"完全二叉树", []int{1, 2, 3, 4, 5, 6, 7}, 10, true},
		{"完全二叉树-不存在", []int{1, 2, 3, 4, 5, 6, 7}, 9, false},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.vals), tt.target); got != tt.want {
					t.Errorf("%s(%v, %d) = %v, want %v", name, tt.vals, tt.target, got, tt.want)
				}
			})
		}
	}
}
