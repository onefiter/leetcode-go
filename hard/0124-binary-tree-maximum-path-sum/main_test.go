package maxpathsum

import (
	"math"
	"testing"
)

const null = math.MinInt64

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

var solutions = map[string]func(*TreeNode) int{
	"maxPathSum":          maxPathSum,
	"maxPathSumIterative": maxPathSumIterative,
}

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want int
	}{
		{"示例1", []int{1, 2, 3}, 6},
		{"示例2", []int{-10, 9, 20, null, null, 15, 7}, 42},
		{"单个负数节点", []int{-3}, -3},
		{"全为负数取最大", []int{-2, -1, -3}, -1},
		{"舍弃负子树", []int{2, -1, 3}, 5},
		{"不经过根", []int{-10, 1, -20, 5, 6}, 12},
		{"左斜链", []int{1, 2, null, 3, null, 4}, 10},
	}
	for name, fn := range solutions {
		for _, tt := range tests {
			t.Run(name+"/"+tt.name, func(t *testing.T) {
				if got := fn(buildTree(tt.root)); got != tt.want {
					t.Errorf("%s(%v) = %d, want %d", name, tt.root, got, tt.want)
				}
			})
		}
	}
}
