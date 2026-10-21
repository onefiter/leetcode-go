package rob3

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
	"rob":     rob,
	"robMemo": robMemo,
}

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		root []int
		want int
	}{
		{"示例1", []int{3, 2, 3, null, 3, null, 1}, 7},
		{"示例2", []int{3, 4, 5, 1, 3, null, 1}, 9},
		{"空树", []int{}, 0},
		{"单节点", []int{5}, 5},
		{"隔层更优", []int{4, 1, null, 2, null, 3}, 7},
		{"偷两个孩子", []int{1, 5, 6}, 11},
		{"满二叉树", []int{1, 2, 3, 4, 5, 6, 7}, 23},
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
