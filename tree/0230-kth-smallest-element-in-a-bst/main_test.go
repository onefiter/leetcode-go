package kthsmallestbst

import (
	"math"
	"testing"
)

const null = math.MinInt64

var solutions = map[string]func(*TreeNode, int) int{
	"kthSmallest":          kthSmallest,
	"kthSmallestRecursive": kthSmallestRecursive,
}

// buildTree 按 LeetCode 层序数组构造二叉树，null 表示空节点。
func buildTree(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == null {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != null {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != null {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		k    int
		want int
	}{
		{"示例1", []int{3, 1, 4, null, 2}, 1, 1},
		{"示例2", []int{5, 3, 6, 2, 4, null, null, 1}, 3, 3},
		{"单节点", []int{7}, 1, 7},
		{"左斜链-最大", []int{3, 2, null, 1}, 3, 3},
		{"右斜链-中间", []int{1, null, 2, null, 3}, 2, 2},
		{"完全二叉树-第1", []int{4, 2, 6, 1, 3, 5, 7}, 1, 1},
		{"完全二叉树-第5", []int{4, 2, 6, 1, 3, 5, 7}, 5, 5},
		{"完全二叉树-第7", []int{4, 2, 6, 1, 3, 5, 7}, 7, 7},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.vals), tt.k); got != tt.want {
					t.Errorf("%s(%v, %d) = %d, want %d", name, tt.vals, tt.k, got, tt.want)
				}
			})
		}
	}
}
